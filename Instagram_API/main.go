package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/shhhivam/Instagram_API/controllers"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func main() {
	password := "secret"
	hash, _ := HashPassword(password)
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.GET("/posts/:id", uc.GetPost)
	r.POST("/posts", uc.CreatePost)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
