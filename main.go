package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"

	"github.com/mleyb/go-blockchain/chain"
	"github.com/mleyb/go-blockchain/router"
)

func run() error {
	mux := router.MakeMuxRouter()

	httpAddr := os.Getenv("PORT")

	log.Println("Listening on ", os.Getenv("PORT"))

	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()

		genesisBlock := chain.Block{
			Index:     0,
			Timestamp: t.String(),
			Data:      "Genesis block",
			Hash:      "",
			PrevHash:  "",
		}

		spew.Dump(genesisBlock)

		chain.Blockchain = append(chain.Blockchain, genesisBlock)
	}()

	log.Fatal(run())
}
