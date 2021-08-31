package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

// RegisterHandlers registers the NFT REST routes
func RegisterHandlers(cliCtx client.Context, r *mux.Router, queryRoute string) {
	registerQueryRoutes(cliCtx, r, queryRoute)
	registerTxRoutes(cliCtx, r, queryRoute)
}

const (
	RestParamDenomID = "denom-id"
	RestParamOwner   = "owner"
)

type issueDenomReq struct {
	BaseReq          rest.BaseReq `json:"base_req"`
	Owner            string       `json:"owner"`
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Schema           string       `json:"schema"`
	Symbol           string       `json:"symbol"`
	MintRestricted   bool         `json:"mint_restricted"`
	UpdateRestricted bool         `json:"update_restricted"`
}
