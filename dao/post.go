package dao

import (
	"github.com/opensourceai/go-api-service/models"
	"github.com/opensourceai/go-api-service/pkg/page"
)

type PostDao interface {
	// 新增帖子
	Add(post *models.Post) (err error)
	// 删除帖子
	DeleteByIds(ids ...uint) (err error)
	// 通过用户ID查找用户的帖子
	FindAllByUserId(page page.Page, userId uint) (err error, postList []models.Post)
	// 返回该帖子的信息
	FindById(id uint) (err error, post *models.Post)
	// 更新帖子
	Update(post *models.Post) (err error)
}