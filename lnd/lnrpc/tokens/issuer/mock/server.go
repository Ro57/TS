package mocks

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net"

	"github.com/lightninglabs/protobuf-hex-display/json"
	"github.com/pkg/errors"
	"github.com/pkt-cash/pktd/lnd/lnrpc/tokens/issuer"
	"google.golang.org/grpc"
)

type Server struct {
	// Nest unimplemented server implementation in order to satisfy server interface
	issuer.UnimplementedIssuerServer
}

func RunServerServing(host string, stopSig <-chan struct{}) {
	var (
		child = &Server{}
		root  = grpc.NewServer()
	)
	issuer.RegisterIssuerServer(root, child)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	go func() {
		err := root.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		<-stopSig
		root.Stop()
	}()
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
