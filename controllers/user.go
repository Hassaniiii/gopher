package controllers

import (
	"github.com/Hassaniiii/gopher/models"

	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

type userController struct {
	userIDPattern	*regexp.Regexp
}

/// Constructor
func newUserController() *userController {
	return &userController {
		userIDPattern	: regexp.MustCompile(`^/users/(\d+)`),
	}
}

/// HTTP Handlers
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		uc.ServeUsers(w, r)
	} else {
		match := uc.userIDPattern.MatchString(r.URL.Path)
		if !match {
			w.WriteHeader(http.StatusBadRequest); return
		}
		uc.ServeUser(w, r)
	}
}

/// GET /users
func (uc *userController) getAllUsers(w http.ResponseWriter) {
	users := models.GetUsers()
	encodeResponseToJSON(users, w)
}

/// POST /users
func (uc *userController)addNewUser(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest); return
	}

	user, err = models.AddUser(user)
	if err != nil {
		w.WriteHeader(http.StatusConflict); return
	}
	encodeResponseToJSON(user, w)
}

/// GET /users/<id>
func (uc *userController)getUserByID(id int, w http.ResponseWriter) {
	user, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound); return
	}
	encodeResponseToJSON(user, w)
}

/// DELETE /users/<id>
func (uc *userController)removeUserByID(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound); return
	}
	w.WriteHeader(http.StatusOK)
}

/// PUT /users/<id>
func (uc *userController)updateUser(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest); return
	}
	user, err = models.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound); return
	}
	encodeResponseToJSON(user, w)
}

/// Helper functions
func (uc userController)ServeUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uc.getAllUsers(w)
	case http.MethodPost:
		uc.addNewUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (uc userController)ServeUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(uc.userIDPattern.FindStringSubmatch(r.URL.Path)[1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest); return
	}

	switch r.Method {
	case http.MethodGet:
		uc.getUserByID(id, w)
	case http.MethodPut:
		uc.updateUser(w, r)
	case http.MethodDelete:
		uc.removeUserByID(id, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (uc *userController)parseRequest(r *http.Request) (models.User, error) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}