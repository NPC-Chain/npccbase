package rpc

import (
	"github.com/NPC-Chain/npccbase/client/account"
	"github.com/NPC-Chain/npccbase/client/context"
	"github.com/NPC-Chain/npccbase/types"
	"github.com/gorilla/mux"
	"net/http"
)

func registerQueryRoutes(ctx context.CLIContext, m *mux.Router) {
	m.HandleFunc("/accounts/{bech32Address}", queryAccountHandleFunc(ctx)).Methods("GET")
	m.HandleFunc("/accounts/pubkey/{bech32PubKey}", queryAccountPubkeyDecodeHandleFunc(ctx)).Methods("GET")
}

func queryAccountPubkeyDecodeHandleFunc(cliContext context.CLIContext) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		beh32PubKey := vars["bech32PubKey"]

		pk, err := types.GetAccPubKeyBech32(beh32PubKey)
		if err != nil {
			WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		PostProcessResponseBare(writer, cliContext, pk)
	}
}

func queryAccountHandleFunc(cliContext context.CLIContext) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		bech32addr := vars["bech32Address"]

		addr, err := types.AccAddressFromBech32(bech32addr)
		if err != nil {
			WriteErrorResponse(writer, http.StatusBadRequest, err.Error())
			return
		}

		br, _ := ParseRequestForm(request)
		ctx := br.Setup(cliContext)

		acc, err := account.GetAccount(ctx, addr)
		if err != nil {
			Write40XErrorResponse(writer, err)
			return
		}

		PostProcessResponseBare(writer, ctx, acc)
	}
}
