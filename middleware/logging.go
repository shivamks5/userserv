package middleware

import (
	"time"

	"github.com/go-kit/log"
	"github.com/shivamks5/userserv/model"
	"github.com/shivamks5/userserv/service"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.UserService
}

func NewLoggingMiddleware(logger log.Logger, next service.UserService) service.UserService {
	return &loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

func (lmw *loggingMiddleware) GetUser(id string) (user model.User, err error) {
	defer func(start time.Time) {
		_ = lmw.logger.Log(
			"method", "GetUser",
			"id", id,
			"error", err,
			"took", time.Since(start),
		)
	}(time.Now())
	user, err = lmw.next.GetUser(id)
	return
}

func (lmw *loggingMiddleware) CreateUser(user model.User) (id string, err error) {
	defer func(start time.Time) {
		_ = lmw.logger.Log(
			"method", "CreateUser",
			"id", id,
			"name", user.Name,
			"error", err,
			"took", time.Since(start),
		)
	}(time.Now())
	id, err = lmw.next.CreateUser(user)
	return
}

func (lmw *loggingMiddleware) DeleteUser(id string) (err error) {
	defer func(start time.Time) {
		_ = lmw.logger.Log(
			"method", "DeleteUser",
			"id", id,
			"error", err,
			"took", time.Since(start),
		)
	}(time.Now())
	err = lmw.next.DeleteUser(id)
	return
}

func (lmw *loggingMiddleware) ListUsers() []model.User {
	defer func(start time.Time) {
		_ = lmw.logger.Log(
			"method", "ListUsers",
			"took", time.Since(start),
		)
	}(time.Now())
	return lmw.next.ListUsers()
}
