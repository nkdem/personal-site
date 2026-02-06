package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiconfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pulumiCfg := pulumiconfig.New(ctx, "")

		hetznerInfra, err := setupHetznerInfrastructure(ctx)
		if err != nil {
			return err
		}

		tailscaleAuthKey := pulumiCfg.RequireSecret("tailscaleAuthKey")
		cloudInit := renderCloudInit(tailscaleAuthKey)

		server, err := createWebServer(ctx, hetznerInfra, cloudInit)
		if err != nil {
			return err
		}

		if err := setupCloudflareDNS(ctx, server); err != nil {
			return err
		}

		return nil
	})
}
