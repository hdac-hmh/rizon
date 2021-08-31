package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// constant used to indicate that some field should not be updated
const (
	TypeMsgIssueDenom	 = "issue_denom"
	TypeMsgTransferNFT	 = "transfer_nft"
	TypeMsgEditNFT		 = "edit_nft"
	TypeMsgMintNFT		 = "mint_nft"
	TypeMsgBurnNFT		 = "burn_nft"
	TypeMsgTransferDenom = "transfer_denom"
)

var (
	_ sdk.Msg = &MsgIssueDenom{}
)

// NewMsgIssueDenom is a constructor function for MsgSetName
func NewMsgIssueDenom(denomID, denomName, schema, sender, symbol string, mintRestricted, updateRestricted bool) *MsgIssueDenom {
	return &MsgIssueDenom{
		Sender:			  sender,
		Id:				  denomID,
		Name:			  denomName,
		Schema:			  schema,
		Symbol:			  symbol,
		MintRestricted:	  mintRestricted,
		UpdateRestricted: updateRestricted,
	}
}

// Route Implements Msg
func (msg MsgIssueDenom) Route() string { return RouterKey }

// Type Implements Msg
func (msg MsgIssueDenom) Type() string { return TypeMsgIssueDenom }

// ValidateBasic Implements Msg
func (msg MsgIssueDenom) ValidateBasic() error {
	if err := ValidateDenomID(msg.Id); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg
func (msg MsgIssueDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg
func (msg MsgIssueDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

//// NewMsgTransferNFT is a constructor function for MsgSetName
//func NewMsgTransferNFT (
//	tokenID, denomID, tokenName, tokenURI, tokenData, sender, recipient string,
//	) *MsgTransferNFT {
//	return &MsgTransferNFT {
//		Id:			tokenID,
//		DenomId:	denomID,
//		Name:		tokenName,
//		URI:		tokenURI,
//		Data:		tokenData,
//		Sender:		sender,
//		Recipient:	recipient,
//	}
//}
//
//func (msg MsgTransferNFT) Route() string {}
//func (msg MsgTransferNFT) Type() string {}
//func (msg MsgTransferNFT) ValidateBasic() error {}
//func (msg MsgTransferNFT) GetSignBytes() []byte {}
//func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {}
//
//// NewMsgEditNFT is a constructor function for MsgSetName
//func NewMsgEditNFT(
//	tokenID, denomID, tokenName, tokenURI, tokenData, sender string,
//	) *MsgEditNFT {
//	return &MsgEditNFT{
//		Id:			tokenID,
//		DenomId:	denomID,
//		Name:		tokenName,
//		URI:		tokenURI,
//		Data:		tokenData,
//		Sender:		sender,
//	}
//}
//
//func (msg NewMsgEditNFT) Route() string {}
//func (msg NewMsgEditNFT) Type() string {}
//func (msg NewMsgEditNFT) ValidateBasic() error {}
//func (msg NewMsgEditNFT) GetSignBytes() []byte {}
//func (msg NewMsgEditNFT) GetSigners() []sdk.AccAddress {}