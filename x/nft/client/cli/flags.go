package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagTokenName = "name"
	FlagTokenURI  = "uri"
	FlagTokenData = "data"
	FlagRecipient = "recipient"

	FlagDenomName		 = "name"
	FlagDenomID			 = "denom-id"
	FlagSchema			 = "schema"
	FlagSymbol			 = "symbol"
	FlagMintRestricted	 = "mint-restricted"
	FlagUpdateRestricted = "update-restricted"
)

var (
	FsIssueDenom	= flag.NewFlagSet("", flag.ContinueOnError)
	FsMintNFT       = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsIssueDenom.String(FlagSchema, "", "Denom data structure definition")
	FsIssueDenom.String(FlagDenomName, "", "The name of the denom")
	FsIssueDenom.String(FlagSymbol, "", "The symbol of the denom")
	FsIssueDenom.Bool(FlagMintRestricted, false, "mint restricted of nft under denom")
	FsIssueDenom.Bool(FlagUpdateRestricted, false, "update restricted of nft under denom")

	FsMintNFT.String(FlagTokenURI, "", "URI for supplemental off-chain tokenData (should return a JSON object)")
	FsMintNFT.String(FlagRecipient, "", "Receiver of the nft, if not filled, the default is the sender of the transaction")
	FsMintNFT.String(FlagTokenData, "", "The origin data of the nft")
	FsMintNFT.String(FlagTokenName, "", "The name of the nft")
}
