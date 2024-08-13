package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/satioO/basics/v2/helpers"
	"github.com/satioO/basics/v2/models"
	"github.com/satioO/basics/v2/usecase"
)

type userHandler struct {
	us usecase.UserService
}

func RegisterUserHandler(router *http.ServeMux, us usecase.UserService) {
	handler := userHandler{us}

	router.HandleFunc("GET /users", helpers.Register(handler.findUsers))
	router.HandleFunc("GET /users/{userId}", helpers.Register(handler.findUserById))
	router.HandleFunc("POST /users", helpers.Register(handler.createUser))
	router.HandleFunc("PUT /users/{userId}", helpers.Register(handler.updateUser))
	router.HandleFunc("DELETE /users/{userId}", helpers.Register(handler.deleteUser))
}

func (h *userHandler) findUsers(w http.ResponseWriter, r *http.Request) error {

	users := h.us.FindUsers()
	return helpers.WriteToJSON(w, http.StatusOK, users)
}

func (h *userHandler) findUserById(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(r.PathValue("userId"))

	if err != nil {
		return helpers.InvalidRequestData(map[string]string{"error": "invalid user id"})
	}

	user, err := h.us.FindUserById(userId)
	if err != nil {
		return helpers.NewApiError(http.StatusInternalServerError, err)
	}

	return helpers.WriteToJSON(w, http.StatusOK, user)
}

func (h *userHandler) createUser(w http.ResponseWriter, r *http.Request) error {
	var body models.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return helpers.InvalidJson()
	}

	defer r.Body.Close()

	user, err := h.us.CreateUser(&body)
	if err != nil {
		return helpers.NewApiError(http.StatusInternalServerError, err)
	}

	return helpers.WriteToJSON(w, http.StatusCreated, user)
}

func (h *userHandler) updateUser(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(r.PathValue("userId"))

	if err != nil {
		return helpers.InvalidRequestData(map[string]string{"error": "invalid user id"})
	}

	user, err := h.us.UpdateUser(userId)
	if err != nil {
		return helpers.NewApiError(http.StatusInternalServerError, err)
	}

	return helpers.WriteToJSON(w, http.StatusOK, user)
}

func (h *userHandler) deleteUser(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(r.PathValue("userId"))

	if err != nil {
		return helpers.InvalidRequestData(map[string]string{"error": "invalid user id"})
	}

	err = h.us.DeleteUser(userId)

	if err != nil {
		return helpers.NewApiError(http.StatusInternalServerError, err)
	}

	return helpers.WriteToJSON(w, http.StatusOK, nil)
}
