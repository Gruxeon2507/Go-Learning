package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s\n", address)
}
func number() int {
	return 2
}

type test1 struct {
	number int
}

func (n test1) testNum() int {
	n.number = 2
	return n.number
}
func main() {
	temp := test1{
		number: 1,
	}
	for i = temp.testNum(); i < 5; i++ {
		fmt.Println(i)
	}
	a := func() {
		fmt.Println("huhu")
	}

	a()
}
