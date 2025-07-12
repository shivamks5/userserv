package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/shivamks5/userserv/endpoint"
	"github.com/shivamks5/userserv/model"
	"github.com/shivamks5/userserv/service"
)

func MakeHTTPHandler(eps endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	getUserHandler := httptransport.NewServer(
		eps.GetUserEndpoint,
		decodeGetRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
	createUserHandler := httptransport.NewServer(
		eps.CreateUserEndpoint,
		decodeCreateRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
	deleteUserHandler := httptransport.NewServer(
		eps.DeleteUserEndpoint,
		decodeDeleteRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
	listUsersHandler := httptransport.NewServer(
		eps.ListUsersEndpoint,
		decodeListRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
	r.Methods("GET").Path("/users/{id}").Handler(getUserHandler)
	r.Methods("POST").Path("/users").Handler(createUserHandler)
	r.Methods("DELETE").Path("/users/{id}").Handler(deleteUserHandler)
	r.Methods("GET").Path("/users").Handler(listUsersHandler)
	return r
}

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["id"]
	if id == "" {
		return nil, service.ErrBadRequest
	}
	return model.IDNumber{ID: id}, nil
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, service.ErrBadRequest
	}
	return req, nil
}

func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["id"]
	if id == "" {
		return nil, service.ErrBadRequest
	}
	return model.IDNumber{ID: id}, nil
}

func decodeListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var code int
	switch err {
	case service.ErrBadRequest, service.ErrInvalidField:
		code = http.StatusBadRequest
	case service.ErrNotFound:
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
