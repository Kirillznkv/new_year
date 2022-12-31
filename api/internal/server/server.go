package server

import (
	"encoding/json"
	"fmt"
	"github.com/Kirillznkv/new_year/api/internal/model"
	"github.com/Kirillznkv/new_year/api/internal/store"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	router *mux.Router
	store  *store.Store
}

func NewServer(store *store.Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("Post")
	s.router.HandleFunc("/users", s.handleUsersGet()).Methods("Get")

	s.router.HandleFunc("/answers", s.handleAnswersCreate()).Methods("Post")
	s.router.HandleFunc("/answers", s.handleAnswersGet()).Methods("Get")

	s.router.HandleFunc("/text", s.handleTextCreate()).Methods("Post")
}

func (s *Server) handleTextCreate() http.HandlerFunc {
	type request struct {
		Text string `json:"text"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			log.Println(r.Body)
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.Users().FindById(user_id)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Text = fmt.Sprintf("%s\n%s", u.FirstName, req.Text)

		saveText(u)

		s.respond(w, r, http.StatusCreated, nil)
	}
}

func saveText(u *model.User) {
	savePath := fmt.Sprintf("./texts/%s", u.FirstName+u.SecondName)

	file, _ := os.Create(savePath)
	defer file.Close()

	data := u.Text

	file.Write([]byte(data))
}

func (s *Server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		FirstName  string `json:"first_name"`
		SecondName string `json:"second_name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			log.Println(r.Body)
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			FirstName:  req.FirstName,
			SecondName: req.SecondName,
		}
		if err := s.store.Users().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u.ID)
	}
}

func (s *Server) handleUsersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := s.store.Users().GetUsers()

		s.respond(w, r, http.StatusCreated, users)
	}
}

func (s *Server) handleAnswersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		file, _, err := r.FormFile("image")
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		lvl, err := strconv.Atoi(r.URL.Query().Get("lvl"))
		user_id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		image, err := ioutil.ReadAll(file)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		a := &model.Answer{
			Lvl:    lvl,
			UserId: user_id,
			Image:  image,
		}
		if err := s.store.Answers().Create(a); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, nil)
	}
}

func (s *Server) handleAnswersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := s.store.Answers().GetAnswers()

		s.respond(w, r, http.StatusCreated, users)
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
