package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
)

// ProtoV5ProviderServerFactory returns a muxed terraform-plugin-go protocol v5 provider factory function.
// This factory function is suitable for use with the terraform-plugin-go Serve function.
func ProtoV5ProviderServerFactory(ctx context.Context) (func() tfprotov5.ProviderServer, error) {
	providers := []func() tfprotov5.ProviderServer{
		Provider().GRPCProvider,
	}

	muxServer, err := tf5muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		return nil, err
	}

	return muxServer.ProviderServer, nil
}

// ProtoV5ProviderFactory returns a muxed terraform-plugin-go protocol v5 provider factory function.
// This factory function is suitable for use with the terraform-plugin-sdk/v2 acceptance testing framework.
func ProtoV5ProviderFactory(ctx context.Context) func() (tfprotov5.ProviderServer, error) {
	return func() (tfprotov5.ProviderServer, error) {
		providerServerFactory, err := ProtoV5ProviderServerFactory(ctx)

		if err != nil {
			return nil, err
		}

		return providerServerFactory(), nil
	}
}
