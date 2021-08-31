package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagDenomName		 = "name"
	FlagDenomID			 = "denom-id"
	FlagSchema			 = "schema"
	FlagSymbol			 = "symbol"
	FlagMintRestricted	 = "mint-restricted"
	FlagUpdateRestricted = "update-restricted"
)

var (
	FsIssueDenom		= flag.NewFlagSet("", flag.ContinueOnError)

)

func init() {
	FsIssueDenom.String(FlagSchema, "", "Denom data structure definition")
	FsIssueDenom.String(FlagDenomName, "", "The name of the denom")
	FsIssueDenom.String(FlagSymbol, "", "The symbol of the denom")
	FsIssueDenom.Bool(FlagMintRestricted, false, "mint restricted of nft under denom")
	FsIssueDenom.Bool(FlagUpdateRestricted, false, "update restricted of nft under denom")
}
