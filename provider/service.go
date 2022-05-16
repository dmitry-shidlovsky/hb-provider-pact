package provider

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/dmitry-shidlovsky/TestPact/model"
)

var userRepository = &UserRepository{
	Users: map[string]*model.User{
		"sally": {
			FirstName: "Test First Name",
			LastName:  "Test Last Name",
			Username:  "dsh",
			Type:      "admin",
			ID:        1,
		},
	},
}

func WithCorrelationID(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		w.Header().Set("X-Api-Correlation-Id", uuid.String())
		h.ServeHTTP(w, r)
	}
}

func IsAuthenticated(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	a := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(a[len(a)-1])

	user, err := userRepository.ByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		resBody, _ := json.Marshal(user)
		w.Write(resBody)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	resBody, _ := json.Marshal(userRepository.GetUsers())
	w.Write(resBody)
}

func commonMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return WithCorrelationID(IsAuthenticated(f))
}

func GetHTTPHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/", commonMiddleware(GetUser))
	mux.HandleFunc("/users/", commonMiddleware(GetUsers))
	
	return mux
}
