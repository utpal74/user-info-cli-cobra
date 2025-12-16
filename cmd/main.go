package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/user-info-cli-tool/internal/controller/user"
	httphandler "github.com/user-info-cli-tool/internal/handler/http"
	"github.com/user-info-cli-tool/internal/repository/memory"
)

func main() {
	memory := memory.New()
	ctrl := user.New(memory)
	h := httphandler.New(ctrl)
	r := mux.NewRouter()
	r.HandleFunc("/users", h.GetUser).Methods("GET")
	r.HandleFunc("/users", h.CreateUser).Methods("POST")
	log.Print("user service starting at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
