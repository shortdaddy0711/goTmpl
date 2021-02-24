package main

import (
	"os"
	"text/template"
)

// User -
type User struct {
	Name  string
	Email string
	Age   uint
}

func main() {
	user1 := User{Name: "namsoo", Email: "namsoo@gmail.com", Age: 41}
	user2 := User{Name: "olivia", Email: "olivia@gmail.com", Age: 8}
	tmpl, err := template.New("users").ParseFiles("templates/templ1.tmpl")

	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user1)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user2)

}
