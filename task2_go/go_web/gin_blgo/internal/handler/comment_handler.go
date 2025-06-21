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

func CommentCreateHandler(c *gin.Context) {
	var commentReq types.CommetCreateReq
	if err := c.ShouldBind(&commentReq); err != nil {
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
		Select("id").
		Take(&post, commentReq.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 文章不存在，返回自定义提示
			types.ErrorRes(c, 0, "文章不存在")
			return
		}
	}

	err = dao.DB.Create(
		&model.Comment{
			UserID:  uID,
			PostID:  commentReq.ID,
			Content: commentReq.Content,
		}).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("创建评论失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "创建评论成功.", nil)
}

func CommentListHandler(c *gin.Context) {
	var commentReq types.CommetListReq
	if err := c.ShouldBind(&commentReq); err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("参数传递错误: %s", err.Error()))
		logger.Log.Error("参数传递错误: ", zap.Any("err", err))
		return
	}
	if commentReq.PageSize == 0 {
		commentReq.PageSize = 10
	}
	if commentReq.PageNum == 0 {
		commentReq.PageNum = 1
	}

	var post model.Post
	err := dao.DB.
		Select("id").
		Take(&post, commentReq.PostID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 文章不存在，返回自定义提示
			types.ErrorRes(c, 0, "文章不存在")
			return
		}
	}

	var comments []model.Comment
	err = dao.DB.
		Preload("User").
		Where("post_id = ?", commentReq.PostID).
		Order("created_at DESC").
		Limit(commentReq.PageSize).
		Offset((commentReq.PageNum - 1) * commentReq.PageSize).
		Find(&comments).Error
	if err != nil {
		types.ErrorRes(c, 0, fmt.Sprintf("评论查询失败: %s", err.Error()))
		logger.Log.Error("数据库错误: ", zap.Any("err", err))
		return
	}

	types.SuccessRes(c, "评论查询成功.", comments)
}
