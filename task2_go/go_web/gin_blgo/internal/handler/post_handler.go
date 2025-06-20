package handler

import (
	"fmt"
	"gin_blog/internal/model"
	"gin_blog/internal/pkg/dao"
	"gin_blog/internal/pkg/logger"
	"gin_blog/internal/types"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostCreateHandler(c *gin.Context) {
	var post types.PostCreateReq
	if err := c.ShouldBind(&post); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}

	userID, _ := c.Get("userID")
	uId, ok := userID.(uint)
	if !ok {
		types.ErrorRes(c, 0, "用户信息获取错误")
		return
	}
	err := dao.DB.Create(
		&model.Post{
			Title:   post.Title,
			Content: post.Content,
			UserID:  uId,
		}).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("创建文章失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "创建文章成功.", nil)
}

func PostListHandler(c *gin.Context) {
	var post types.PostListReq
	if err := c.ShouldBind(&post); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}
	if post.PageSize == 0 {
		post.PageSize = 10
	}
	if post.PageNum == 0 {
		post.PageNum = 1
	}
	fmt.Println(post)

	var posts []model.Post
	if post.Title != "" {

	}
	err := dao.DB.
		Find(&posts).
		Limit(post.PageSize).
		Offset((post.PageNum - 1) * post.PageSize).
		Scan(&posts).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("文章查询失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "创建文章成功.", posts)
}

func PostDetailHandler(c *gin.Context) {
	var post types.PostListReq
	if err := c.ShouldBind(&post); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}

	var posts []model.Post
	err := dao.DB.
		Find(&posts).
		Limit(post.PageSize).
		Offset((post.PageNum - 1) * post.PageSize).
		Scan(&posts).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("文章查询失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "创建文章成功.", posts)
}
