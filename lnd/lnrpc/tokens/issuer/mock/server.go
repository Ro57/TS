package mocks

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/lightninglabs/protobuf-hex-display/json"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/issuer"
	"google.golang.org/grpc"
)

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	issuer.UnimplementedIssuerServer
}

func RegisterServer(root *grpc.Server) {
	child := &Server{}

	issuer.RegisterIssuerServer(root, child)
}

// Override method of unimplemented server
func (s *Server) SignTokenPurchase(ctx context.Context, req *issuer.SignTokenPurchaseRequest) (*issuer.SignTokenPurchaseResponse, error) {
	bytes, err := json.Marshal(req)
	if err != nil {
		return nil, errors.WithMessage(err, "marshalling request")
	}

	hash := sha256.Sum256(bytes)

	resp := &issuer.SignTokenPurchaseResponse{
		IssuerSignature: fmt.Sprintf("%x", hash),
	}

	return resp, nil
}
