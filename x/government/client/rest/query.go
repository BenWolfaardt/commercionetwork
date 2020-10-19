package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	"github.com/commercionetwork/commercionetwork/x/government/types"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/government/address",
		getGovernmentAddr(cliCtx)).
		Methods("GET")

	r.HandleFunc(
		"/government/tumbler",
		getTumblerAddr(cliCtx)).
		Methods("GET")
}

func getGovernmentAddr(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryGovernmentAddress)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getTumblerAddr(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTumblerAddress)
		res, _, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
