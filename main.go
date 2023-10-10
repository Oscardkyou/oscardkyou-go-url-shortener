package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/Oscardkyou/oscardkyou-go-url-shortener/pkg/urlencoder"
)

type User struct {
	Name  string
	Age   int
	Money int
}

func (u User) GetAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) SetNewName(newName string) {
	u.Name = newName
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/HomePage.html")
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		r.Body.Close()

		longURL := string(body)
		us := urlencoder.New()
		shortUrl := us.Encode(longURL)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(shortUrl))
		return
	}
	tmpl.Execute(w, nil)
}

func contact_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is good!!!")
}

func handleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts", contact_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
