package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mleyb/go-blockchain/handlers"
)

func MakeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", handlers.GetBlockchainHandler).Methods("GET")
	muxRouter.HandleFunc("/", handlers.WriteBlockHandler).Methods("POST")

	return muxRouter
}
