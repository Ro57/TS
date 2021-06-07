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
	}

	// Apply pagination
	if req.Offset > 0 {
		resp.Offers = resp.Offers[req.Offset:]
	}
	if req.Limit > 0 {
		if offersNum := len(resp.Offers); offersNum > int(req.Limit) {
			resp.Total = uint64(offersNum) // keep unlimited offers total num
		} else {
			resp.Total = req.Limit
		}

		resp.Offers = resp.Offers[:req.Limit]
	}

	return resp, nil
}
