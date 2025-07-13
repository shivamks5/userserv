package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/shivamks5/userserv/endpoint"
	"github.com/shivamks5/userserv/errs"
	"github.com/shivamks5/userserv/model"
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
	updateUserHandler := httptransport.NewServer(
		eps.UpdateUserEndpoint,
		decodeUpdateRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(encodeError),
	)
	patchUserHandler := httptransport.NewServer(
		eps.PatchUserEndpoint,
		decodePatchRequest,
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
	r.Methods("PUT").Path("/users/{id}").Handler(updateUserHandler)
	r.Methods("PATCH").Path("/users/{id}").Handler(patchUserHandler)
	r.Methods("DELETE").Path("/users/{id}").Handler(deleteUserHandler)
	r.Methods("GET").Path("/users").Handler(listUsersHandler)
	return r
}

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["id"]
	return id, nil
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errs.ErrBadRequest
	}
	return req, nil
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var user model.User
	id := mux.Vars(r)["id"]
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, nil
}

func decodePatchRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var user map[string]interface{}
	id := mux.Vars(r)["id"]
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	user["id"] = id
	return user, nil
}

func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id := mux.Vars(r)["id"]
	return id, nil
}

func decodeListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var err error
	var minmax model.MinMax
	query := r.URL.Query()
	minAge := query.Get("min")
	maxAge := query.Get("max")
	if minAge != "" {
		minmax.Mini, err = strconv.Atoi(minAge)
		if err != nil {
			return nil, errs.ErrBadRequest
		}
	}
	if maxAge != "" {
		minmax.Maxi, err = strconv.Atoi(maxAge)
		if err != nil {
			return nil, errs.ErrBadRequest
		}
	}
	return minmax, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var code int
	switch err {
	case errs.ErrBadRequest, errs.ErrInvalidField:
		code = http.StatusBadRequest
	case errs.ErrNotFound:
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
