package types

// NFT module event types
var (
	EventTypeIssueDenom	   = "issue_denom"
	EventTypeMintNFT       = "mint_nft"

	AttributeValueCategory = ModuleName

	AttributeKeyCreator   = "creator"
	AttributeKeyRecipient = "recipient"
	AttributeKeyTokenID   = "token_id"
	AttributeKeyTokenURI  = "token_uri"
	AttributeKeyDenomID	  = "denom_id"
	AttributeKeyDenomName = "denom_name"
)