package handler

import (
	"errors"
	"fmt"
	"gin_blog/internal/model"
	"gin_blog/internal/pkg/dao"
	"gin_blog/internal/pkg/logger"
	"gin_blog/internal/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func PostCreateHandler(c *gin.Context) {
	var postReq types.PostCreateReq
	if err := c.ShouldBind(&postReq); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}

	userID, _ := c.Get("userID")
	uID, ok := userID.(uint)
	if !ok {
		types.ErrorRes(c, 0, "用户信息获取错误")
		return
	}
	err := dao.DB.Create(
		&model.Post{
			Title:   postReq.Title,
			Content: postReq.Content,
			UserID:  uID,
		}).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("创建文章失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "创建文章成功.", nil)
}

func PostListHandler(c *gin.Context) {
	var postReq types.PostListReq
	if err := c.ShouldBind(&postReq); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}
	if postReq.PageSize == 0 {
		postReq.PageSize = 10
	}
	if postReq.PageNum == 0 {
		postReq.PageNum = 1
	}

	var totalCount int64
	var posts []model.Post

	baseQuery := dao.DB.Model(&model.Post{}).Omit("content")
	if postReq.Title != "" {
		baseQuery = baseQuery.Where("title LIKE ?", "%"+postReq.Title+"%")
	}

	baseQuery.Count(&totalCount)
	err := baseQuery.
		Order("created_at DESC"). // 按创建时间倒序
		Limit(postReq.PageSize).
		Offset((postReq.PageNum - 1) * postReq.PageSize).
		Find(&posts).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("文章查询失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "文章查询成功.", gin.H{"total_count": totalCount, "posts": posts})
}

func PostDetailHandler(c *gin.Context) {
	var postReq types.PostBaseReq
	if err := c.ShouldBind(&postReq); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}

	var postDetail model.Post
	err := dao.DB.
		Preload("User").             // 加载作者信息
		Preload("Comments").         // 加载所有评论
		Preload("Comments.User").    // 加载每条评论的作者
		Preload("Comments.Post").    // 加载每条评论所属帖子
		Where("id = ?", postReq.ID). // 明确查询条件
		Take(&postDetail).           // 获取单条数据
		Error

	if postDetail.ID == 0 {
		types.ErrorRes(c, 0, "文章不存在")
		return
	}
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("文章查询失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "文章查询成功.", postDetail)
}

func PostUpdateHandler(c *gin.Context) {
	var postReq types.PostUpdateReq
	if err := c.ShouldBind(&postReq); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}

	var post model.Post
	err := dao.DB.
		Select("id").
		Take(&post, postReq.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 文章不存在，返回自定义提示
			types.ErrorRes(c, 0, "文章不存在")
			return
		}

		types.ErrorRes(c, 0, fmt.Sprintf("文章查询失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	post.Title = postReq.Title
	post.Content = postReq.Content

	err = dao.DB.Model(&model.Post{}).
		Where("id = ?", post.ID).
		Updates(post).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("文章更新失败: %s", err.Error()))
		logger.Log.Error("文章更新失败: ", zap.Any("err", err))
		return
	}
	types.SuccessRes(c, "文章更新成功.", nil)
}

func PostDeleteHandler(c *gin.Context) {
	var postReq types.PostDeleteReq
	if err := c.ShouldBind(&postReq); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}

	userID, _ := c.Get("userID")
	uID, ok := userID.(uint)
	if !ok {
		types.ErrorRes(c, 0, "用户信息获取错误")
		return
	}

	var post model.Post
	err := dao.DB.
		Select("id", "user_id").
		Take(&post, postReq.ID).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("文章查询失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}
	if post.ID == 0 {
		types.ErrorRes(c, 0, "文章不存在.")
		return
	}
	if post.UserID != uID {
		types.ErrorRes(c, 0, "您不属于文章作者.")
		return
	}
	err = dao.DB.Delete(&model.Post{}, post.ID).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("文章删除失败: %s", err.Error()))
		return
	}
	types.SuccessRes(c, "文章删除成功.", nil)
}
