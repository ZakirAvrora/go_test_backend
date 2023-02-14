package app

import (
	"ZakirAvrora/go_test_backend/service1/internal/model"
	"ZakirAvrora/go_test_backend/service1/internal/repository"
	"ZakirAvrora/go_test_backend/service1/internal/store"
	"ZakirAvrora/go_test_backend/service1/utility/client"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

const url = "http://app2:8000/generate-salt"

type app struct {
	repo repository.Repository
}

func New(repo repository.Repository) *app {
	return &app{repo: repo}
}

func (a *app) PostUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user model.PostReqUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "body message is not appropriate", http.StatusBadRequest)
		return
	}

	if !validEmail(user.Email) {
		http.Error(w, "email is not appropriate", http.StatusBadRequest)
		return
	}

	c := client.New(url)
	body, err := c.PostReq()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var salt model.Salt
	if err := json.Unmarshal(body, &salt); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	hashPswd := GetMD5Hash(salt.Salt + user.Password)

	dbUser := &model.DbUser{
		Email:    user.Email,
		Salt:     salt.Salt,
		Password: hashPswd,
	}

	if err := a.repo.CreateUser(*dbUser); err != nil {

		if errors.Is(err, store.ErrUserExists) {
			http.Error(w, "User with same email already exists", http.StatusBadRequest)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	jsonResponse(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func (a *app) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := chi.URLParam(r, "email")

	user, err := a.repo.GetUser(strings.TrimSpace(email))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "No such user", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server error", http.StatusMethodNotAllowed)
		}
		return
	}

	err = jsonResponse(w, http.StatusOK, *user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func jsonResponse(w http.ResponseWriter, code int, body interface{}) error {
	response, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error in sending json response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	return nil
}

func validEmail(email string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
