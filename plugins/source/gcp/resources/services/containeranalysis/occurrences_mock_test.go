// Code generated by codegen; DO NOT EDIT.

package containeranalysis

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"

	"cloud.google.com/go/containeranalysis/apiv1beta1"

	pb "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"

	"google.golang.org/api/option"
)

func createOccurrences() (*client.Services, error) {
	fakeServer := &fakeOccurrencesServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterGrafeasV1Beta1Server(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := containeranalysis.NewGrafeasV1Beta1Client(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		ContaineranalysisGrafeasV1Beta1Client: svc,
	}, nil
}

type fakeOccurrencesServer struct {
	pb.UnimplementedGrafeasV1Beta1Server
}

func (f *fakeOccurrencesServer) ListOccurrences(context.Context, *pb.ListOccurrencesRequest) (*pb.ListOccurrencesResponse, error) {
	resp := pb.ListOccurrencesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestOccurrences(t *testing.T) {
	client.MockTestHelper(t, Occurrences(), createOccurrences, client.TestOptions{})
}
