package main

import (
	"fmt"

	mux "gitolite.aws.sgdev.org/mux.git"
)

// Router is a router from gorilla/mux
type Router struct{ mux.Router }

func main() {
	var x = mux.Router{}
	fmt.Println("hello world")
}
