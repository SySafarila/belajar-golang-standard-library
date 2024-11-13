package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")
	var port *int = flag.Int("port", 0, "Put your database port")

	flag.Parse()

	fmt.Println("host", *host)
	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("port", *port)
}
