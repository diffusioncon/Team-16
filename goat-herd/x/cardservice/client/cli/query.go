package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	//"github.com/DecentralCardGame/Cardchain/x/cardservice"
)

// GetCmdResolveCard queries information about a card
func GetCmdCard(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "resolve [cardId]",
		Short: "resolve card",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve name - %s \n", string(name))
				return nil
			}

			fmt.Println(string(res))

			return nil

			//var out cardservice.QueryResResolve
			//cdc.MustUnmarshalJSON(res, &out)
			//return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdWhois queries information about a user
func GetCmdWhois(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "whois [name]",
		Short: "Query whois info of user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve whois - %s \n", string(name))
				return nil
			}

			fmt.Println(string(res))

			return nil

			/*
			var out cardservice.Whois
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
			*/
		},
	}
}

// GetCmdNames queries a list of 50 cards
func GetCmdCardList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "cards",
		Short: "cards",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/cards", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query names\n")
				return nil
			}

			fmt.Println(string(res))

			return nil
/*
			var out cardservice.QueryResCards
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
			*/
		},
	}
}

// GetCmdVotableCardList looks up the cards votable by a user
func GetCmdVotableCardList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "votable-cards [name]",
		Short: "Query cards votable by a user.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/votable-cards/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not get query names\n")
				return nil
			}

			fmt.Println(string(res))

			return nil
/*
			var out cardservice.QueryResCards
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
			*/
		},
	}
}
