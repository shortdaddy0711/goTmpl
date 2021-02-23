package main

import (
	"os"
	"text/template"
)

// User -
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	user1 := User{Name: "namsoo", Email: "namsoo@gmail.com", Age: 41}
	user2 := User{Name: "olivia", Email: "olivia@gmail.com", Age: 8}
	tmpl, err := template.New("users").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge: {{.Age}}\n")

	if err != nil {
		panic(err)
	}
	
	tmpl.Execute(os.Stdout, user1)
	tmpl.Execute(os.Stdout, user2)

}
