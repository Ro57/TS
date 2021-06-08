package mocks

import (
	"context"
	"fmt"
	"time"

	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/replicator"
	"google.golang.org/grpc"
)

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	replicator.UnimplementedReplicatorServer
}

func RegisterServer(root *grpc.Server) {
	child := &Server{}

	replicator.RegisterReplicatorServer(root, child)
}

// Override method of unimplemented server
func (s *Server) GetTokenOffers(ctx context.Context, filter *replicator.TokenOffersFilter) (*replicator.TokenOffers, error) {
	const (
		eachIssuerTokensNum = 3
		offersNum           = 1000 * eachIssuerTokensNum
	)

	offers := make([]*replicator.TokenOffer, 0, offersNum)

	// Fill mocked offers such, that each issuer has several tokens present
	for i := offersNum / eachIssuerTokensNum; i > 0; i-- {
		offer := &replicator.TokenOffer{
			ValidUntilMillis: time.Now().UnixNano() + int64(i)*1000,
			IssuerInfo: &replicator.IssuerInfo{
				Id:             fmt.Sprintf("issuer_%d", i),
				IdentityPubkey: "issuer_node_pub_key",
				Host:           "issuer_ip",
			},
			Token: fmt.Sprintf("token_%d", i),
			Price: uint64(1 + i*2),
		}
		offers = append(offers, offer)

		offer = &replicator.TokenOffer{
			ValidUntilMillis: time.Now().UnixNano() + int64(i)*1000,
			IssuerInfo: &replicator.IssuerInfo{
				Id:             fmt.Sprintf("issuer_%d", i),
				IdentityPubkey: "issuer_node_pub_key",
				Host:           "issuer_ip",
			},
			Token: fmt.Sprintf("token_%d", i+1),
			Price: uint64(1 + i*4),
		}
		offers = append(offers, offer)

		offer = &replicator.TokenOffer{
			ValidUntilMillis: time.Now().UnixNano() + int64(i)*1000,
			IssuerInfo: &replicator.IssuerInfo{
				Id:             fmt.Sprintf("issuer_%d", i),
				IdentityPubkey: "issuer_node_pub_key",
				Host:           "issuer_ip",
			},
			Token: fmt.Sprintf("token_%d", i+2),
			Price: uint64(1 + i*8),
		}
		offers = append(offers, offer)
	}

	resp := &replicator.TokenOffers{
		Offers: offers,
		Total:  offersNum,
	}

	// Apply filter by issuer id
	if filter.IssuerId != "" {
		issuerOffers := make([]*replicator.TokenOffer, 0, eachIssuerTokensNum)

		for _, offer := range resp.Offers {
			if len(issuerOffers) == eachIssuerTokensNum {
				break
			}

			if offer.IssuerInfo.Id == filter.IssuerId {
				issuerOffers = append(issuerOffers, offer)
			}
		}

		resp.Offers = issuerOffers
		resp.Total = uint64(len(resp.Offers))
	}

	// Apply pagination
	if filter.Params.Offset > 0 {
		if int(filter.Params.Offset) <= len(resp.Offers)-1 {
			resp.Offers = resp.Offers[filter.Params.Offset:]
		} else {
			resp.Offers = nil
		}
	}
	if filter.Params.Limit > 0 {
		if int(filter.Params.Limit) <= len(resp.Offers)-1 {
			resp.Offers = resp.Offers[:filter.Params.Limit]
		}
	}

	return resp, nil
}

// Override method of unimplemented server
func (s *Server) GetTokenBalances(ctx context.Context, filter *replicator.TokenBalancesFilter) (*replicator.TokenBalances, error) {
	const (
		tokensNum = 100
	)
	balances := make([]*replicator.TokenBalance, 0, tokensNum)

	// Fill mocked token balances. At this time balances are not owned by a specific holder
	for i := tokensNum; i > 0; i-- {
		balance := &replicator.TokenBalance{
			Token:     fmt.Sprintf("token_%d", i),
			Available: uint64(i*2 + 1),
			Frozen:    uint64(i*3 + 1),
		}
		balances = append(balances, balance)
	}

	resp := &replicator.TokenBalances{
		Balances: balances,
		Total:    tokensNum,
	}

	// Apply pagination
	if filter.Params.Offset > 0 {
		if int(filter.Params.Offset) <= len(resp.Balances)-1 {
			resp.Balances = resp.Balances[filter.Params.Offset:]
		} else {
			resp.Balances = nil
		}
	}

	if filter.Params.Limit > 0 {
		if int(filter.Params.Limit) <= len(resp.Balances)-1 {
			resp.Balances = resp.Balances[:filter.Params.Limit]
		}
	}

	return resp, nil
}
