package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"urlshortener/pkg"
)

func (u User) GetAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) SetNewName(newName string) {
	u.Name = newName
}

tmpl, _ := template.ParseFiles("templates/HomePage.html") // Используйте правильное имя файла


			return
		}

		longURL := string(body)
		us := pkg.New()
		shortUrl := us.Encode(longURL)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(shortUrl))
		return
	}

	//bob.SetNewName("Alexander")
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	//fmt.Fprintf(w, bob.GetAllInfo())
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
	us := pkg.New()

	fmt.Println("Long URL:", longUrl)
}
