package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jessevdk/go-flags"
	"github.com/recoilme/graceful"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Vacancy struct {
	Id      int64  `json:"id"`
	Date    string `json:"date"`
	Vacancy string `json:"vacancy"`
	Code    string `json:"code"`
	Opened  bool   `json:"opened"`
}

type UserCreds struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id         int64  `json:"id,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	SecondName string `json:"second_name,omitempty"`
	Code       string `json:"code,omitempty"`
	Vacancy    string `json:"vacancy,omitempty"`
}

var UsersCreds = make(map[string]UserCreds)
var Users = make(map[int64]User)
var Vacancies = make([]Vacancy, 0)
var VacancyRequests = make([]VacancyRequest, 0)

var config struct {
	Port string `long:"port" env:"PORT" default:"8081"`
}

func main() {
	if _, err := flags.Parse(&config); err != nil {
		os.Exit(1)
	}
	rand.Seed(time.Now().Unix())

	err := ReadDatabase()
	if err != nil {
		log.Fatal(err)
	}

	_, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	graceful.Unignore(quit, func() error {
		log.Print("[WARN] interrupt signal")
		err := WriteDatabase()
		if err != nil {
			panic(err)
		}
		cancel()
		return nil
	}, graceful.Terminate...)

	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))
		r.Use(middleware.Logger)
		r.Use(middleware.StripSlashes)
		r.Use(middleware.Timeout(time.Second * 60))

		r.Post("/sing_in", Login)
		r.Post("/sing_up", Register)

		r.Post("/vacancy/get", GetVacancy)
		r.Post("/vacancy/add", AddVacancy)
		r.Post("/vacancy/delete", DeleteVacancy)

		r.Post("/request/get", GetRequests)
		r.Post("/request/add", CreateRequest)
		r.Post("/request/delete", DeleteRequest)
	})

	FileServer(r)

	s := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Println("Server starting")
	log.Println(s.ListenAndServe())
}

func ReadDatabase() error {
	data, err := ioutil.ReadFile("users.jsonl")
	if err != nil {
		return err
	}

	jsons := strings.Split(string(data), "\n")
	jsons = jsons[:len(jsons)-1]
	for _, u := range jsons {
		user := User{}
		err := json.Unmarshal([]byte(u), &user)
		if err != nil {
			return err
		}
		Users[user.Id] = user
	}

	data, err = ioutil.ReadFile("users_creds.jsonl")
	if err != nil {
		return err
	}

	jsons = strings.Split(string(data), "\n")
	jsons = jsons[:len(jsons)-1]
	for _, u := range jsons {
		user := UserCreds{}
		err := json.Unmarshal([]byte(u), &user)
		if err != nil {
			return err
		}
		UsersCreds[user.Username] = user
	}

	data, err = ioutil.ReadFile("vacancy.jsonl")
	if err != nil {
		return err
	}

	jsons = strings.Split(string(data), "\n")
	jsons = jsons[:len(jsons)-1]
	for _, v := range jsons {
		vacancy := Vacancy{}
		err := json.Unmarshal([]byte(v), &vacancy)
		if err != nil {
			return err
		}
		Vacancies = append(Vacancies, vacancy)
	}

	data, err = ioutil.ReadFile("vacancy_reqs.jsonl")
	if err != nil {
		return err
	}
	jsons = strings.Split(string(data), "\n")
	jsons = jsons[:len(jsons)-1]
	for _, v := range jsons {
		vRec := VacancyRequest{}
		err := json.Unmarshal([]byte(v), &vRec)
		if err != nil {
			return err
		}
		VacancyRequests = append(VacancyRequests, vRec)
	}

	return nil
}

func WriteDatabase() error {
	buf := bytes.Buffer{}
	for _, value := range Users {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		buf.WriteString(string(data) + "\n")
	}
	err := ioutil.WriteFile("users.jsonl", buf.Bytes(), 0644)
	if err != nil {
		return err
	}
	buf.Reset()

	for _, value := range UsersCreds {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		buf.WriteString(string(data) + "\n")
	}
	err = ioutil.WriteFile("users_creds.jsonl", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	buf.Reset()

	for _, value := range Vacancies {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		buf.WriteString(string(data) + "\n")
	}
	err = ioutil.WriteFile("vacancy.jsonl", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	buf.Reset()

	for _, value := range VacancyRequests {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		buf.WriteString(string(data) + "\n")
	}
	err = ioutil.WriteFile("vacancy_reqs.jsonl", buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func FileServer(router *chi.Mux) {
	root := "./dist"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}
