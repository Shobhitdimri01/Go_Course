package main

import "fmt"

/*
Interface Composition
Go allows composite interfaces, which combine multiple interfaces into
one. A type must implement all the methods from all interfaces
in the composite to satisfy it.
*/

type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct{}

func (File) Read(p []byte) (int, error)  { return len(p), nil }
func (File) Write(p []byte) (int, error) { return len(p), nil }

func main() {

	// dog1 := &Dog{
	// 	name: "ABC",
	// }
	// AnimalCalls(dog1)
	showType("hello")

	// var rw ReadWriter = File{}
	// rw.Read([]byte("data"))
	// rw.Write([]byte("data"))
	// Describe("Hello") // String: Hello
	// Describe(42)      // Not a string

	// var p Printer
	// doc := &Document{Content: "Pointer Example"}
	// p = doc // Works because Print() has a pointer receiver
	// p.Print()

	//standard library
	// person := Person{Name: "John"}
	// fmt.Println(person) // Output: Person: John
	// fmt.Println("hello")
	// databaseInitilization()
}

func showType(i interface{}) {
	var p int
	var ok bool
	p, ok = i.(int)
	if !ok {
		fmt.Println(ok)
	} else {
		fmt.Println(p)
		fmt.Printf("%v is of type %T\n", i, i)
	}
}

func Describe(i interface{}) {
	v, ok := i.(string)
	if ok {
		fmt.Println("String:", v)
	} else {
		fmt.Println("Not a string")
	}
}

type Printer interface {
	Print()
}

type Document struct {
	Content string
}

func (d *Document) Print() {
	fmt.Println(d.Content)
}

type Person struct {
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("This is not your typical println Person: %s", p.Name)
}

type Animal interface {
	Speak()
	Walk()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func AnimalCalls(a Animal) {
	a.Speak()
}

func (d *Dog) Speak() {
	fmt.Println("Woof! ->", d.name)
}
func (d *Dog) Walk() {
	fmt.Println("Woof! ->", d.name)
}

func (c *Cat) Speak() {
	fmt.Println("meow ->", c.name)
}
