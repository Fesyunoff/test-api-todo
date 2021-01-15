//+build !swipe

// Code generated by Swipe v2.0.0-beta6. DO NOT EDIT.

package transport

import (
	"context"
	"time"

	"github.com/fesyunoff/api/pkg/controller"
	"github.com/fesyunoff/api/pkg/controller/dto"
	"github.com/go-kit/kit/log"
)

type ServiceLoggingMiddleware struct {
	next   controller.Service
	logger log.Logger
}

func (s *ServiceLoggingMiddleware) Add(ctx context.Context, task dto.Task) (string, error) {
	var (
		result string
		err    error
	)
	defer func(now time.Time) {
		logErr := err
		if le, ok := err.(interface{ LogError() error }); ok {
			logErr = le.LogError()
		}
		s.logger.Log("method", "Add", "took", time.Since(now), "task", task, "result", result, "err", logErr)
	}(time.Now())
	result, err = s.next.Add(ctx, task)
	return result, err
}

func (s *ServiceLoggingMiddleware) CreateUser(ctx context.Context, user dto.User) (string, error) {
	var (
		result string
		err    error
	)
	defer func(now time.Time) {
		logErr := err
		if le, ok := err.(interface{ LogError() error }); ok {
			logErr = le.LogError()
		}
		s.logger.Log("method", "CreateUser", "took", time.Since(now), "user", user, "result", result, "err", logErr)
	}(time.Now())
	result, err = s.next.CreateUser(ctx, user)
	return result, err
}

func (s *ServiceLoggingMiddleware) Delete(ctx context.Context, task dto.Task) (string, error) {
	var (
		result string
		err    error
	)
	defer func(now time.Time) {
		logErr := err
		if le, ok := err.(interface{ LogError() error }); ok {
			logErr = le.LogError()
		}
		s.logger.Log("method", "Delete", "took", time.Since(now), "task", task, "result", result, "err", logErr)
	}(time.Now())
	result, err = s.next.Delete(ctx, task)
	return result, err
}

func (s *ServiceLoggingMiddleware) DeleteUser(ctx context.Context, user dto.User) (string, error) {
	var (
		result string
		err    error
	)
	defer func(now time.Time) {
		logErr := err
		if le, ok := err.(interface{ LogError() error }); ok {
			logErr = le.LogError()
		}
		s.logger.Log("method", "DeleteUser", "took", time.Since(now), "user", user, "result", result, "err", logErr)
	}(time.Now())
	result, err = s.next.DeleteUser(ctx, user)
	return result, err
}

func (s *ServiceLoggingMiddleware) Get(ctx context.Context, task dto.Task) ([]*dto.Task, string, error) {
	var (
		out []*dto.Task
		msg string
		err error
	)
	defer func(now time.Time) {
		logErr := err
		if le, ok := err.(interface{ LogError() error }); ok {
			logErr = le.LogError()
		}
		s.logger.Log("method", "Get", "took", time.Since(now), "task", task, "out", len(out), "msg", msg, "err", logErr)
	}(time.Now())
	out, msg, err = s.next.Get(ctx, task)
	return out, msg, err
}

func (s *ServiceLoggingMiddleware) GetUsers(ctx context.Context, user dto.User) ([]*dto.User, string, error) {
	var (
		out []*dto.User
		msg string
		err error
	)
	defer func(now time.Time) {
		logErr := err
		if le, ok := err.(interface{ LogError() error }); ok {
			logErr = le.LogError()
		}
		s.logger.Log("method", "GetUsers", "took", time.Since(now), "user", user, "out", len(out), "msg", msg, "err", logErr)
	}(time.Now())
	out, msg, err = s.next.GetUsers(ctx, user)
	return out, msg, err
}

func (s *ServiceLoggingMiddleware) Update(ctx context.Context, task dto.Task) (string, error) {
	var (
		result string
		err    error
	)
	defer func(now time.Time) {
		logErr := err
		if le, ok := err.(interface{ LogError() error }); ok {
			logErr = le.LogError()
		}
		s.logger.Log("method", "Update", "took", time.Since(now), "task", task, "result", result, "err", logErr)
	}(time.Now())
	result, err = s.next.Update(ctx, task)
	return result, err
}

func NewLoggingServiceMiddleware(s controller.Service, logger log.Logger) controller.Service {
	return &ServiceLoggingMiddleware{next: s, logger: logger}
}