package discovery

import (
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/config"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/config/plugins/discovery/universal"
	"github.com/pkg/errors"
)

var _ config.Config = &DiscoveryConfig{}

type DiscoveryConfig struct {
	Universal *universal.UniversalDiscoveryConfig `yaml:"universal"`
}

func (d *DiscoveryConfig) Validate() error {
	return errors.Wrap(d.Universal.Validate(), "Discovery validation failed")
}

func DefaultDiscoveryConfig() *DiscoveryConfig {
	return &DiscoveryConfig{
		Universal: universal.DefaultUniversalDiscoveryConfig(),
	}
}
