package middleware

import (
	"time"

	"github.com/shivamks5/userserv/metrics"
	"github.com/shivamks5/userserv/model"
	"github.com/shivamks5/userserv/service"
)

type metricsMiddleware struct {
	next   service.UserService
	metric *metrics.Metrics
}

func NewMetricsMiddleware(next service.UserService, metric *metrics.Metrics) service.UserService {
	return &metricsMiddleware{
		next:   next,
		metric: metric,
	}
}

func (mw *metricsMiddleware) GetUser(id string) (user model.User, err error) {
	defer func(start time.Time) {
		mw.track("GetUser", err, start)
	}(time.Now())
	user, err = mw.next.GetUser(id)
	return
}

func (mw *metricsMiddleware) CreateUser(user model.User) (resp model.User, err error) {
	defer func(start time.Time) {
		mw.track("CreateUser", err, start)
	}(time.Now())
	resp, err = mw.next.CreateUser(user)
	return
}

func (mw *metricsMiddleware) UpdateUser(updatedUser model.User) (user model.User, err error) {
	defer func(start time.Time) {
		mw.track("UpdateUser", err, start)
	}(time.Now())
	user, err = mw.next.UpdateUser(updatedUser)
	return
}

func (mw *metricsMiddleware) PatchUser(patchedUser map[string]interface{}) (user model.User, err error) {
	defer func(start time.Time) {
		mw.track("PatchUser", err, start)
	}(time.Now())
	user, err = mw.next.PatchUser(patchedUser)
	return
}

func (mw *metricsMiddleware) DeleteUser(id string) (err error) {
	defer func(start time.Time) {
		mw.track("DeleteUser", err, start)
	}(time.Now())
	err = mw.next.DeleteUser(id)
	return
}

func (mw *metricsMiddleware) ListUsers(name string, minAge, maxAge int) []model.User {
	defer func(start time.Time) {
		mw.metric.RequestCount.With("method", "ListUsers", "error", "false").Add(1)
		mw.metric.RequestLatency.With("method", "ListUsers", "error", "false").Observe(float64(time.Since(start).Microseconds()))
	}(time.Now())
	return mw.next.ListUsers(name, minAge, maxAge)
}

func (mw *metricsMiddleware) track(method string, err error, start time.Time) {
	var errLabel = "false"
	if err != nil {
		errLabel = "true"
		mw.metric.RequestErrors.With("method", method, "error", errLabel).Add(1)
	}
	mw.metric.RequestCount.With("method", method, "error", errLabel).Add(1)
	mw.metric.RequestLatency.With("method", method, "error", errLabel).Observe(float64(time.Since(start).Microseconds()))
}
