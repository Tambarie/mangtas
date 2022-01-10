package main

import (
	"Mangtas_/internal/router"
	"Mangtas_/internal/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main()  {
	env := os.Getenv("GIN_MODE")

	if env != "release"{
		if err := godotenv.Load();err != nil{
			log.Fatalf("could not load env vars : %v",err)
		}
	}

	s := &server.Server{Router: router.NewRouter()}
	s.Start()

}
