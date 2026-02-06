package config

import (
	_ "embed"
	"fmt"
	"sync"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configYAML []byte

// Config holds all infrastructure configuration.
// Secrets are NOT included — they remain in Pulumi.
type Config struct {
	Domain    string          `yaml:"domain"`
	Tailscale TailscaleConfig `yaml:"tailscale"`
	Server    ServerConfig    `yaml:"server"`
	Hetzner   HetznerConfig   `yaml:"hetzner"`
	Ports     PortsConfig     `yaml:"ports"`
}

// TailscaleConfig holds Tailscale network settings.
type TailscaleConfig struct {
	VPNPort int    `yaml:"vpn_port"`
	Tag     string `yaml:"tag"`
}

// ServerConfig holds configuration for the web server.
type ServerConfig struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Role    string `yaml:"role"`
	Backups bool   `yaml:"backups"`
	WorkDir string `yaml:"work_dir"`
}

// HetznerConfig holds Hetzner Cloud settings.
type HetznerConfig struct {
	DefaultLocation string `yaml:"default_location"`
	Image           string `yaml:"image"`
	Project         string `yaml:"project"`
	FirewallName    string `yaml:"firewall_name"`
	SSHKeyName      string `yaml:"ssh_key_name"`
}

// PortsConfig holds service port numbers.
type PortsConfig struct {
	HTTP  int `yaml:"http"`
	HTTPS int `yaml:"https"`
}

var (
	cfg     *Config
	cfgOnce sync.Once
	cfgErr  error
)

// Load parses the embedded YAML configuration.
// Safe for concurrent use — returns the same instance on subsequent calls.
func Load() (*Config, error) {
	cfgOnce.Do(func() {
		cfg = &Config{}
		cfgErr = yaml.Unmarshal(configYAML, cfg)
		if cfgErr != nil {
			cfgErr = fmt.Errorf("parse config.yaml: %w", cfgErr)
		}
	})
	return cfg, cfgErr
}

// MustLoad loads configuration and panics on error.
func MustLoad() *Config {
	c, err := Load()
	if err != nil {
		panic(err)
	}
	return c
}
