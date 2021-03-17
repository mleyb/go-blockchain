package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/mleyb/go-blockchain/api"
	"github.com/mleyb/go-blockchain/chain"
)

func WriteBlockHandler(w http.ResponseWriter, r *http.Request) {
	var m api.Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}

	defer r.Body.Close()

	newBlock, err := chain.GenerateBlock(chain.Blockchain[len(chain.Blockchain)-1], m.Data)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	if chain.IsBlockValid(newBlock, chain.Blockchain[len(chain.Blockchain)-1]) {
		newBlockchain := append(chain.Blockchain, newBlock)
		chain.ReplaceChain(newBlockchain)
		spew.Dump(chain.Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}
