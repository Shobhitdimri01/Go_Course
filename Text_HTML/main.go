package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

func main() {
	Texttemplate()
	fmt.Println()
	HTMLTemplate()
}

func Texttemplate() {
	tmpl := `Hello, {{.Name}}! Welcome to {{.Language}}. Date : {{.Date}}`
	data := map[string]string{
		"Name":     "Alice",
		"Language": "Go",
		"Date":     time.Now().String(),
	}

	t, err := template.New("example").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Name     string
	Age      int
	Language string
}

type Item struct {
	Name  string
	Price float64
}

type ShoppingList struct {
	Items []Item
}

func HTMLTemplate() {
	tmpl := `Name: {{.Name}}, Age: {{.Age}}, Favorite Language: {{.Language}}`
	user := User{Name: "Bob", Age: 25, Language: "Go"}

	t, err := template.New("user").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	t.Execute(os.Stdout, user)
}

func RangeList() {
	tmpl := `
    <h1>Shopping List</h1>
    {{range .Items}}
        <p>{{.Name}}: ${{.Price}}</p>
    {{end}}
    {{if eq (len .Items) 0}}
        <p>No items in the list!</p>
    {{end}}
    `
	list := ShoppingList{
		Items: []Item{
			{Name: "Apple", Price: 1.25},
			{Name: "Milk", Price: 2.50},
		},
	}

	t, err := template.New("shopping").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	t.Execute(os.Stdout, list)
}

func TempCondition() {
	tmpl := `
    {{if .}}
        You are logged in!
    {{else}}
        Please log in.
    {{end}}
    `
	isLoggedIn := true

	t := template.Must(template.New("example").Parse(tmpl))
	t.Execute(os.Stdout, isLoggedIn)
}

type Product struct {
	Name  string
	Price float64
}

func AllCombineConditions() {
	tmpl := `
    {{if gt (len .) 0}}
    Products List:
    {{range .}}
        Name: {{.Name}}, Price: ${{.Price}}
    {{end}}
    {{else}}
        No products available.
    {{end}}
    `
	products := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Mouse", Price: 25.00},
	}

	t := template.Must(template.New("example").Parse(tmpl))
	t.Execute(os.Stdout, products)
}
