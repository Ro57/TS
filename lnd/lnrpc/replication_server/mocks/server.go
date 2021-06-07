package mocks

import (
	"context"
	"fmt"

	"github.com/pkt-cash/pktd/lnd/lnrpc/replication_server"
	"google.golang.org/grpc"
)

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	replication_server.UnimplementedReplicationServerServer
}

func RegisterServer(root *grpc.Server) {
	child := &Server{}

	replication_server.RegisterReplicationServerServer(root, child)
}

// Override method of unimplemented server
func (s *Server) GetTokenOffers(ctx context.Context, req *replication_server.GetTokenOffersRequest) (*replication_server.GetTokenOffersResponse, error) {
	const (
		eachIssuerTokensNum = 3
		offersNum           = 1000 * eachIssuerTokensNum
	)

	offers := make([]*replication_server.TokenOffer, 0, offersNum)

	// Fill mocked offers such, that each issuer has several tokens present
	for i := offersNum / eachIssuerTokensNum; i > 0; i-- {
		offer := &replication_server.TokenOffer{
			IssuerId: fmt.Sprintf("issuer #%d", i),
			Token:    fmt.Sprintf("token_%d", i),
			Price:    uint64(1 + i*2),
		}
		offers = append(offers, offer)

		offer = &replication_server.TokenOffer{
			IssuerId: fmt.Sprintf("issuer #%d", i),
			Token:    fmt.Sprintf("token_%d", i+1),
			Price:    uint64(1 + i*4),
		}
		offers = append(offers, offer)

		offer = &replication_server.TokenOffer{
			IssuerId: fmt.Sprintf("issuer #%d", i),
			Token:    fmt.Sprintf("token_%d", i+2),
			Price:    uint64(1 + i*8),
		}
		offers = append(offers, offer)
	}

	resp := &replication_server.GetTokenOffersResponse{
		Offers: offers,
		Total:  offersNum,
	}

	// Apply filter by issuer id
	if req.IssuerId != "" {
		issuerOffers := make([]*replication_server.TokenOffer, 0, eachIssuerTokensNum)

		for _, offer := range resp.Offers {
			if len(issuerOffers) == eachIssuerTokensNum {
				break
			}

			if offer.IssuerId == req.IssuerId {
				issuerOffers = append(issuerOffers, offer)
			}
		}

		resp.Offers = issuerOffers
		resp.Total = uint64(len(resp.Offers))
	}

	// Apply pagination
	if req.Params.Offset > 0 {
		if int(req.Params.Offset) <= len(resp.Offers)-1 {
			resp.Offers = resp.Offers[req.Params.Offset:]
		} else {
			resp.Offers = nil
		}
	}
	if req.Params.Limit > 0 {
		if int(req.Params.Limit) <= len(resp.Offers)-1 {
			resp.Offers = resp.Offers[:req.Params.Limit]
		}
	}

	return resp, nil
}

// Override method of unimplemented server
func (s *Server) GetTokenBalances(ctx context.Context, req *replication_server.GetTokenBalancesRequest) (*replication_server.GetTokenBalancesResponse, error) {
	const (
		tokensNum = 100
	)
	balances := make([]*replication_server.TokenBalance, 0, tokensNum)

	// Fill mocked token balances. At this time balances are not owned by a specific holder
	for i := tokensNum; i > 0; i-- {
		balance := &replication_server.TokenBalance{
			Token:     fmt.Sprintf("token_%d", i),
			Available: uint64(i*2 + 1),
			Frozen:    uint64(i*3 + 1),
		}
		balances = append(balances, balance)
	}

	resp := &replication_server.GetTokenBalancesResponse{
		Balances: balances,
		Total:    tokensNum,
	}

	// Apply pagination
	if req.Params.Offset > 0 {
		if int(req.Params.Offset) <= len(resp.Balances)-1 {
			resp.Balances = resp.Balances[req.Params.Offset:]
		} else {
			resp.Balances = nil
		}
	}

	if req.Params.Limit > 0 {
		if int(req.Params.Limit) <= len(resp.Balances)-1 {
			resp.Balances = resp.Balances[:req.Params.Limit]
		}
	}

	return resp, nil
}
