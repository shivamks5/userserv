package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/shivamks5/userserv/model"
	"github.com/shivamks5/userserv/service"
)

type Endpoints struct {
	GetUserEndpoint    endpoint.Endpoint
	CreateUserEndpoint endpoint.Endpoint
	DeleteUserEndpoint endpoint.Endpoint
	ListUsersEndpoint  endpoint.Endpoint
}

func NewEndpoints(svc service.UserService) Endpoints {
	return Endpoints{
		GetUserEndpoint:    makeGetUserEndpoint(svc),
		CreateUserEndpoint: makeCreateUserEndpoint(svc),
		DeleteUserEndpoint: makeDeleteUserEndpoint(svc),
		ListUsersEndpoint:  makeListUsersEndpoint(svc),
	}
}

func makeGetUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.IDNumber)
		user, err := svc.GetUser(req.ID)
		if err != nil {
			return nil, err
		}
		return model.Response{Data: user}, nil
	}
}

func makeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.User)
		id, err := svc.CreateUser(req)
		if err != nil {
			return nil, err
		}
		return model.Response{Data: map[string]string{"id": id}}, nil
	}
}

func makeDeleteUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.IDNumber)
		err := svc.DeleteUser(req.ID)
		if err != nil {
			return nil, err
		}
		return model.Response{Data: map[string]string{"message": "user deleted successfully"}}, nil
	}
}

func makeListUsersEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result := svc.ListUsers()
		return model.Response{Data: result}, nil
	}
}
