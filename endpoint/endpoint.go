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
	UpdateUserEndpoint endpoint.Endpoint
	PatchUserEndpoint  endpoint.Endpoint
	DeleteUserEndpoint endpoint.Endpoint
	ListUsersEndpoint  endpoint.Endpoint
}

func NewEndpoints(svc service.UserService) Endpoints {
	return Endpoints{
		GetUserEndpoint:    makeGetUserEndpoint(svc),
		CreateUserEndpoint: makeCreateUserEndpoint(svc),
		UpdateUserEndpoint: makeUpdateUserEndpoint(svc),
		PatchUserEndpoint:  makePatchUserEndpoint(svc),
		DeleteUserEndpoint: makeDeleteUserEndpoint(svc),
		ListUsersEndpoint:  makeListUsersEndpoint(svc),
	}
}

func makeGetUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		user, err := svc.GetUser(id)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.User)
		user, err := svc.CreateUser(req)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeUpdateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.User)
		user, err := svc.UpdateUser(req)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makePatchUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(map[string]interface{})
		user, err := svc.PatchUser(req)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeDeleteUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		err := svc.DeleteUser(id)
		if err != nil {
			return nil, err
		}
		return map[string]string{"message": "user deleted successfully"}, nil
	}
}

func makeListUsersEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.MinMax)
		users := svc.ListUsers(req.Mini, req.Maxi)
		return users, nil
	}
}
