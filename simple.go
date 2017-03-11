package main

import (
	"fmt"

	mux "gitolite.aws.sgdev.org/gorilla/mux.git"
)

// Router is a router from gorilla/mux
type Router struct{ mux.Router }

func main() {
	fmt.Println("hello world")
}
