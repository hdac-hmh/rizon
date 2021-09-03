package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/rizon-world/rizon/x/nft/types"
)

func registerTxRoutes(cliCtx client.Context, r *mux.Router, queryRoute string) {
	// Issue a denom
	r.HandleFunc("/nft/nfts/denoms/issue", issueDenomHandlerFn(cliCtx)).Methods("POST")
	// Mint an NFT
	r.HandleFunc("/nft/nfts/mint", mintNFTHandlerFn(cliCtx)).Methods("POST")
}

func issueDenomHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req issueDenomReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		// create the message
		msg := types.NewMsgIssueDenom(req.ID, req.Name, req.Schema, req.Owner, req.Symbol, req.MintRestricted, req.UpdateRestricted)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}

func mintNFTHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req mintNFTReq
		if !rest.ReadRESTReq(w, r, cliCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		if req.Recipient == "" {
			req.Recipient = req.Owner
		}
		// create the message
		msg := types.NewMsgMintNFT(
			req.ID,
			req.DenomID,
			req.Name,
			req.URI,
			req.Data,
			req.Owner,
			req.Recipient,
		)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		tx.WriteGeneratedTxResponse(cliCtx, w, req.BaseReq, msg)
	}
}