package main

import (
	"fmt"

	"github.com/nkdem/personal-site/infrastructure/config"
	"github.com/pulumi/pulumi-hcloud/sdk/go/hcloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pulumiconfig "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// HetznerInfrastructure holds shared Hetzner resources.
type HetznerInfrastructure struct {
	Provider *hcloud.Provider
	SSHKey   *hcloud.SshKey
	Firewall *hcloud.Firewall
}

// setupHetznerInfrastructure creates shared Hetzner infrastructure.
func setupHetznerInfrastructure(ctx *pulumi.Context) (*HetznerInfrastructure, error) {
	cfg := config.MustLoad()
	pulumiCfg := pulumiconfig.New(ctx, "")
	hetznerToken := pulumiCfg.RequireSecret("hetznerToken")

	provider, err := hcloud.NewProvider(ctx, "hetzner-provider", &hcloud.ProviderArgs{
		Token: hetznerToken,
	})
	if err != nil {
		return nil, err
	}

	sshPublicKey := pulumiCfg.RequireSecret("hetznerSSHPublicKey")
	sshKey, err := hcloud.NewSshKey(ctx, "personal-site-ssh-key", &hcloud.SshKeyArgs{
		Name:      pulumi.String(cfg.Hetzner.SSHKeyName),
		PublicKey: sshPublicKey,
		Labels: pulumi.StringMap{
			"managed-by": pulumi.String("pulumi"),
			"project":    pulumi.String(cfg.Hetzner.Project),
		},
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}

	firewall, err := createHetznerFirewall(ctx, provider)
	if err != nil {
		return nil, err
	}

	return &HetznerInfrastructure{
		Provider: provider,
		SSHKey:   sshKey,
		Firewall: firewall,
	}, nil
}

// createHetznerFirewall creates a cloud firewall with security rules.
func createHetznerFirewall(ctx *pulumi.Context, provider *hcloud.Provider) (*hcloud.Firewall, error) {
	cfg := config.MustLoad()
	pulumiCfg := pulumiconfig.New(ctx, "")

	// Empty = SSH blocked from public internet (Tailscale-only access)
	allowedSSHIPs := pulumiCfg.Get("allowedSSHIPs")

	anyIPv4 := pulumi.String("0.0.0.0/0")
	anyIPv6 := pulumi.String("::/0")

	rules := hcloud.FirewallRuleArray{
		// Tailscale VPN
		&hcloud.FirewallRuleArgs{
			Direction: pulumi.String("in"),
			Protocol:  pulumi.String("udp"),
			Port:      pulumi.String(fmt.Sprintf("%d", cfg.Tailscale.VPNPort)),
			SourceIps: pulumi.StringArray{anyIPv4, anyIPv6},
		},
		// HTTP
		&hcloud.FirewallRuleArgs{
			Direction: pulumi.String("in"),
			Protocol:  pulumi.String("tcp"),
			Port:      pulumi.String(fmt.Sprintf("%d", cfg.Ports.HTTP)),
			SourceIps: pulumi.StringArray{anyIPv4, anyIPv6},
		},
		// HTTPS
		&hcloud.FirewallRuleArgs{
			Direction: pulumi.String("in"),
			Protocol:  pulumi.String("tcp"),
			Port:      pulumi.String(fmt.Sprintf("%d", cfg.Ports.HTTPS)),
			SourceIps: pulumi.StringArray{anyIPv4, anyIPv6},
		},
	}

	if allowedSSHIPs != "" {
		rules = append(rules, &hcloud.FirewallRuleArgs{
			Direction: pulumi.String("in"),
			Protocol:  pulumi.String("tcp"),
			Port:      pulumi.String("22"),
			SourceIps: pulumi.StringArray{pulumi.String(allowedSSHIPs)},
		})
	}

	return hcloud.NewFirewall(ctx, "personal-site-firewall", &hcloud.FirewallArgs{
		Name:  pulumi.String(cfg.Hetzner.FirewallName),
		Rules: rules,
		Labels: pulumi.StringMap{
			"managed-by": pulumi.String("pulumi"),
			"project":    pulumi.String(cfg.Hetzner.Project),
		},
	}, pulumi.Provider(provider))
}
