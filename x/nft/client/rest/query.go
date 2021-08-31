package rest

import (
	"encoding/binary"
	"fmt"

	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/rizon-world/rizon/x/nft/types"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router, queryRoute string) {
	// Get the total supply of a collection or owner
	r.HandleFunc(fmt.Sprintf("/%s/collctions/{%s}/supply", types.ModuleName, RestParamDenomID), querySupply(cliCtx, queryRoute)).Methods("GET")
}

func querySupply(cliCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		denomID := mux.Vars(r)[RestParamDenomID]
		err := types.ValidateDenomID(denomID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		// TODO:
		// need to refactoring
		var owner sdk.AccAddress
		ownerStr := r.FormValue(RestParamOwner)
		if len(ownerStr) > 0 {
			owner, err = sdk.AccAddressFromBech32(ownerStr)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}
		params := types.NewQuerySupplyParams(denomID, owner)
		bz, err := cliCtx.LegacyAmino.MarshalJSON(params)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s", queryRoute, types.QuerySupply), bz,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		out := binary.LittleEndian.Uint64(res)
		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, out)
	}
}
