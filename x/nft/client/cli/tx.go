package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"io/ioutil"

	"github.com/rizon-world/rizon/x/nft/types"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:						types.ModuleName,
		Short:						"NFT transactions subcommands",
		DisableFlagParsing: 		true,
		SuggestionsMinimumDistance: 2,
		RunE:						client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdIssueDenom(),
	)

	return txCmd
}

// GetCmdIssueDenom is the CLI command for an IssueDenom transaction
func GetCmdIssueDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "issue [denom-id]",
		Long: "Issue a new denom.",
		Example: fmt.Sprintf(
			"$ %s tx nft issue <denom-id> "+
				"--from=<key-name> "+
				"--name=<denom-name> "+
				"--symbol=<denom-symbol> "+
				"--mint-restricted=<mint-restricted> "+
				"--update-restricted=<update-restricted> "+
				"--schema=<schema-content or path to schema.json> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denomName, err := cmd.Flags().GetString(FlagDenomName)
			if err != nil {
				return err
			}
			schema, err := cmd.Flags().GetString(FlagSchema)
			if err != nil {
				return err
			}
			symbol, err := cmd.Flags().GetString(FlagSymbol)
			if err != nil {
				return err
			}
			mintRestricted, err := cmd.Flags().GetBool(FlagMintRestricted)
			if err != nil {
				return err
			}
			updateRestricted, err := cmd.Flags().GetBool(FlagUpdateRestricted)
			if err != nil {
				return err
			}
			optionsContent, err := ioutil.ReadFile(schema)
			if err == nil {
				schema = string(optionsContent)
			}

			msg := types.NewMsgIssueDenom(
				args[0],
				denomName,
				schema,
				clientCtx.GetFromAddress().String(),
				symbol,
				mintRestricted,
				updateRestricted,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsIssueDenom)
	_ = cmd.MarkFlagRequired(FlagMintRestricted)
	_ = cmd.MarkFlagRequired(FlagUpdateRestricted)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

//// GetCmdMintNFT is the CLI command for a MintNFT transaction
//func GetCmdMintNFT() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:  "mint [denom-id] [nft-id]",
//		Long: "Mint an NFT and set teh owner to the recipient.",
//		Example: fmt.Sprintf(
//			"$ %s tx nft mint <denom-id> <nft-id> "+
//				"--uri=<uri> "+
//				"--recipient=<recipient> "+
//				"--from=<key-name> "+
//				"--chain-id=<chain-id> "+
//				"--fees=<fee>",
//				version.AppName,
//			),
//		Args: cobra.ExactArgs(2),
//		RunE: func(cmd *cobra.Command, args []string) error {
//
//		},
//	}
//	cmd.Flags().AddFlagSet(FsMintNFT)
//	flags.AddTxFlagsToCmd(cmd)
//
//	return cmd
//}