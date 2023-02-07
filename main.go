package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	route "simulator/application/route"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

func main() {
	routeObj := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	err := routeObj.LoadPositions()
	if err != nil {
		panic(err)
	}
	routes, err := routeObj.ExportJsonPosition()
	if err != nil {
		panic(err)
	}
	fmt.Println(routes[0])
}
