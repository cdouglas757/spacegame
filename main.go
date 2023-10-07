package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cdouglas757/spacegame/db"
	"github.com/cdouglas757/spacegame/handlers"
	"github.com/cdouglas757/spacegame/spaceman"
)

func main() {
	sms := spaceman.NewSpaceManState()
	spacemenDb, err := db.NewSpaceMenStore()

	if err != nil {
		log.Fatal(err)
	}

	junkyardHandler := handlers.NewJunkyardHandler(sms, spacemenDb)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", junkyardHandler)

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      serveMux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	fmt.Printf("Listening on %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
