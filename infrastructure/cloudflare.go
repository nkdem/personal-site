package main

import (
	"github.com/nkdem/personal-site/infrastructure/config"
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiconfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// setupCloudflareDNS creates DNS records for the domain via Cloudflare.
func setupCloudflareDNS(ctx *pulumi.Context, server *hcloud.Server) error {
	cfg := config.MustLoad()
	pulumiCfg := pulumiconfig.New(ctx, "")

	cfAPIToken := pulumiCfg.RequireSecret("cloudflareAPIToken")
	zoneID := pulumiCfg.Require("cloudflareZoneId")

	provider, err := cloudflare.NewProvider(ctx, "cloudflare-provider", &cloudflare.ProviderArgs{
		ApiToken: cfAPIToken,
	})
	if err != nil {
		return err
	}

	// A record: nkdem.net → server IPv4 (proxied through Cloudflare)
	_, err = cloudflare.NewRecord(ctx, "root-a-record", &cloudflare.RecordArgs{
		ZoneId:  pulumi.String(zoneID),
		Name:    pulumi.String(cfg.Domain),
		Type:    pulumi.String("A"),
		Content: server.Ipv4Address,
		Proxied: pulumi.Bool(true),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	// CNAME: www → nkdem.net (proxied through Cloudflare)
	_, err = cloudflare.NewRecord(ctx, "www-cname-record", &cloudflare.RecordArgs{
		ZoneId:  pulumi.String(zoneID),
		Name:    pulumi.String("www"),
		Type:    pulumi.String("CNAME"),
		Content: pulumi.String(cfg.Domain),
		Proxied: pulumi.Bool(true),
	}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	return nil
}
