package main

import (
	server "Gallery"
	"Gallery/internal/app/db"
	"Gallery/internal/app/handler"
	"flag"
	"fmt"
	_ "gorm.io/gorm/logger"
	"log"
)

func main() {
	srv := new(server.Server)
	log.Println("Started program")
	defer log.Println("Ended Program")
	session := db.Connect()
	myHandler := handler.Handler{DB: session}
	migr := flag.Bool("migrate", false, fmt.Sprint("Migrating process"))
	start := flag.Bool("start", false, fmt.Sprint("Starting server"))
	flag.Parse()
	if *migr {
		log.Println("Migrating progress")
		db.Migrate(session)
	}
	if *start {
		log.Println("Openning server")
		if err := srv.Run(":4000", myHandler.InitRoutes()); err != nil {
			log.Fatalln(err)
		}
	}
	//if *start {
	//	log.Println("Openning server")
	//	r := chi.NewRouter()
	//	r.Use(middleware.Heartbeat("/"))
	//	if err := http.ListenAndServe(":4000", r); err != nil {
	//		log.Fatalln(err)
	//	}
	//}

}
