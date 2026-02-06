package main

import (
	"fmt"

	"github.com/nkdem/personal-site/infrastructure/config"
	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiconfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// createWebServer creates the Hetzner web server.
func createWebServer(ctx *pulumi.Context, infra *HetznerInfrastructure, cloudInit pulumi.StringInput) (*hcloud.Server, error) {
	cfg := config.MustLoad()
	pulumiCfg := pulumiconfig.New(ctx, "")

	location := pulumiCfg.Get("hetznerLocation")
	if location == "" {
		location = cfg.Hetzner.DefaultLocation
	}

	firewallID := infra.Firewall.ID().ApplyT(func(id string) int {
		var intID int
		_, _ = fmt.Sscanf(id, "%d", &intID)
		return intID
	}).(pulumi.IntOutput)

	server, err := hcloud.NewServer(ctx, cfg.Server.Name, &hcloud.ServerArgs{
		Name:        pulumi.String(cfg.Server.Name),
		ServerType:  pulumi.String(cfg.Server.Type),
		Image:       pulumi.String(cfg.Hetzner.Image),
		Location:    pulumi.String(location),
		SshKeys:     pulumi.StringArray{infra.SSHKey.Name},
		FirewallIds: pulumi.IntArray{firewallID},
		UserData:    cloudInit,
		Labels: pulumi.StringMap{
			"managed-by": pulumi.String("pulumi"),
			"project":    pulumi.String(cfg.Hetzner.Project),
			"role":       pulumi.String(cfg.Server.Role),
		},
	}, pulumi.Provider(infra.Provider))
	if err != nil {
		return nil, err
	}

	ctx.Export("webServerIP", server.Ipv4Address)
	ctx.Export("webServerIPv6", server.Ipv6Address)
	ctx.Export("webServerName", server.Name)

	return server, nil
}
