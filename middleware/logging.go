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

func (lmw *loggingMiddleware) CreateUser(user model.User) (resp model.User, err error) {
	defer func(start time.Time) {
		_ = lmw.logger.Log(
			"method", "CreateUser",
			"id", resp.ID,
			"name", user.Name,
			"error", err,
			"took", time.Since(start),
		)
	}(time.Now())
	resp, err = lmw.next.CreateUser(user)
	return
}

func (lmw *loggingMiddleware) UpdateUser(updatedUser model.User) (user model.User, err error) {
	defer func(start time.Time) {
		_ = lmw.logger.Log(
			"method", "UpdateUser",
			"id", user.ID,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())
	user, err = lmw.next.UpdateUser(updatedUser)
	return
}

func (lmw *loggingMiddleware) PatchUser(patchedUser map[string]interface{}) (user model.User, err error) {
	defer func(start time.Time) {
		_ = lmw.logger.Log(
			"method", "PatchUser",
			"id", user.ID,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())
	user, err = lmw.next.PatchUser(patchedUser)
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

func (lmw *loggingMiddleware) ListUsers(minAge, maxAge int) []model.User {
	defer func(start time.Time) {
		_ = lmw.logger.Log(
			"method", "ListUsers",
			"minAge", minAge,
			"maxAge", maxAge,
			"took", time.Since(start),
		)
	}(time.Now())
	return lmw.next.ListUsers(minAge, maxAge)
}
