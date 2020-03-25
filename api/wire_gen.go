// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package api

import (
	"github.com/opensourceai/go-api-service/api/v1"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/service"
	"github.com/opensourceai/go-api-service/pkg/gredis"
)

import (
	_ "github.com/opensourceai/go-api-service/docs"
)

// Injectors from wire.go:

func InitApi() (*Api, func(), error) {
	db, cleanup, err := mysql.NewDao()
	if err != nil {
		return nil, nil, err
	}
	boardDao, err := mysql.NewBoardDao(db)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	redisDao, cleanup2, err := gredis.NewRedis()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	boardService, err := service.NewBoardService(boardDao, redisDao)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	boardApi, err := v1.NewBoardApi(boardService)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	postDao, err := mysql.NewPostDao(db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	postService, err := service.NewPostService(postDao)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	postApi, err := v1.NewPostApi(postService)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	commentDao, err := mysql.NewCommentDao(db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userDao, err := mysql.NewUserDao(db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	commentService, err := service.NewCommentService(commentDao, userDao, postDao)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	commentApi, err := v1.NewCommentService(commentService)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceUserService, err := service.NewUserService(userDao)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	authApi, err := NewAuthApi(serviceUserService)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userApi, err := v1.NewUserApi(serviceUserService)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	api := &Api{
		BoardApi:   boardApi,
		PostApi:    postApi,
		CommentAPi: commentApi,
		AuthApi:    authApi,
		UserApi:    userApi,
	}
	return api, func() {
		cleanup2()
		cleanup()
	}, nil
}
