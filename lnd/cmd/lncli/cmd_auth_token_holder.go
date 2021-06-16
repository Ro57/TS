package main

import (
	"context"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/replicator"
	"github.com/urfave/cli"
)

var authTokenHolderCommand = cli.Command{
	Name:     "authtokenholder",
	Category: "Tokens",
	Usage:    "Authorize user session by jwt",
	Description: `Authorize session by jwt 
		
	You need to pass the username and password to log in as a token holder.
A session will be created for you, which will allow you to carry out further operations with the wallet.`,

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "login",
			Usage: "(required) Unique login for which the token was generated ",
		},
		cli.StringFlag{
			Name:  "password",
			Usage: "(required) User password used for sign in",
		},
	},
	Action: authTokenHolder,
}

func authTokenHolder(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	login, err := requiredString(ctx, "login")
	if err != nil {
		return er.E(err)
	}

	password, err := requiredString(ctx, "password")
	if err != nil {
		return er.E(err)
	}

	// Request offers
	req := &replicator.AuthRequest{
		Login:    login,
		Password: password,
	}

	resp, err := client.AuthTokenHolder(context.TODO(), req)
	if err != nil {
		return er.E(err)
	}
	printRespJSON(resp)

	return nil
}
