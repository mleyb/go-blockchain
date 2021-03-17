package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mleyb/go-blockchain/chain"
)

func GetBlockchainHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(chain.Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
