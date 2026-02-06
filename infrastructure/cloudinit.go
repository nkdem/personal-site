package main

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/nkdem/personal-site/infrastructure/config"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

//go:embed cloudinit.yaml.tmpl
var cloudInitTemplate string

// cloudInitConfig holds parameters for generating a cloud-init script.
type cloudInitConfig struct {
	ServerName       string
	TailscaleTag     string
	WorkDir          string
	TailscaleAuthKey string
}

// renderCloudInit generates a cloud-init script from the template,
// substituting the Tailscale auth key at runtime.
func renderCloudInit(authKey pulumi.StringInput) pulumi.StringOutput {
	cfg := config.MustLoad()

	return authKey.ToStringOutput().ApplyT(func(key string) (string, error) {
		tmpl, err := template.New("cloudinit").Parse(cloudInitTemplate)
		if err != nil {
			return "", err
		}

		data := cloudInitConfig{
			ServerName:       cfg.Server.Name,
			TailscaleTag:     cfg.Tailscale.Tag,
			WorkDir:          cfg.Server.WorkDir,
			TailscaleAuthKey: key,
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return "", err
		}
		return buf.String(), nil
	}).(pulumi.StringOutput)
}
