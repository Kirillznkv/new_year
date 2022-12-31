package front

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handleHome()).Methods("Get")
	s.router.HandleFunc("/lvl/{lvl}", s.handleLvl()).Methods("Get")
	s.router.HandleFunc("/static/{lvl}/{name}", s.handleStatic()).Methods("Get")
	s.router.HandleFunc("/texts", s.handleTextGet()).Methods("Get")
}

func (s *Server) handleTextGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		texts := getTexts()

		t, err := template.ParseFiles("./internal/templates/text.html")
		if err != nil {
			s.error(w, r, 500, err)
		}

		t.Execute(w, texts)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
	}
}

func getTexts() []string {
	var res []string

	files, err := ioutil.ReadDir("./texts")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, path := range files {
		f, err := os.Open("./texts/" + path.Name())
		if err != nil {
			log.Fatal(err)
			return nil
		}
		defer f.Close()

		wr := bytes.Buffer{}
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			wr.WriteString(sc.Text())
		}

		res = append(res, wr.String())
	}

	return res
}

func (s *Server) handleStatic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		lvl := vars["lvl"]
		name := vars["name"]

		imagePath := fmt.Sprintf("./images/lvl_%s/%s", lvl, name)

		http.ServeFile(w, r, imagePath)
	}
}

func (s *Server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./internal/templates/home.html")
		if err != nil {
			log.Println(err)
		}

		t.Execute(w, nil)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
	}
}

func (s *Server) handleLvl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		lvl := vars["lvl"]
		imagesPath := getImagesName(lvl)

		t, err := template.ParseFiles("./internal/templates/answers.html")
		if err != nil {
			s.error(w, r, 500, err)
		}

		t.Execute(w, imagesPath)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
	}
}

func getImagesName(lvl string) []string {
	var res []string

	files, err := ioutil.ReadDir("./images/lvl_" + lvl)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	for _, f := range files {
		path := fmt.Sprintf("http://188.94.158.55/static/%s/%s", lvl, f.Name())
		res = append(res, path)
	}

	return res
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
