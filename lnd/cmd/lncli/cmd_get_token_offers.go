package main

import (
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/urfave/cli"
)

var getTokenOffersCommand = cli.Command{
	Name:     "gettokenoffers",
	Category: "Tokens",
	Usage:    "List information about available token offers per issuer.",
	Description: `List information about available token offers per issuer. 

	"Available offers" means such offers, that officially registered on the off-chain ecosystem and 
all the related deals would be tracked and protected by an overseer. 

	There is an opportunity to list available offers in a pagination-like manner. A such behaviour 
can be achieved by providing additional flags to the command.`,

	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "issuer-id",
			Usage: "(optional) If a value provided, returned offers would belong to the specified token issuer",
		},
		cli.UintFlag{
			Name:  "limit",
			Usage: "(optional) If a value provided, returned offers number would be limited to the specified value",
		},
		cli.UintFlag{
			Name:  "offset",
			Usage: "(optional) If a value provided, returned offers would be at the specified offset \"height\"",
		},
	},
	Action: getTokenOffers,
}

func getTokenOffers(ctx *cli.Context) er.R {
	client, cleanUp := getClient(ctx)
	defer cleanUp()

	var ( // Default request parameters - no pagination
		limit    uint
		offset   uint
		issuerID string
	)

	// Acquire passed values, that are not zero
	if v := ctx.Uint("limit"); v != 0 {
		limit = v
	}
	if v := ctx.Uint("offset"); v != 0 {
		offset = v
	}
	if v := ctx.String("issuer-id"); v != "" {
		issuerID = v
	}

	// TODO: make request; implement gRPC
	_ = limit
	_ = offset
	_ = issuerID

	// TODO: call LND; implement gRPC
	// client.GetTokenOffers()
	_ = client

	// TODO: print response
	printRespJSON(nil)

	return nil
}
