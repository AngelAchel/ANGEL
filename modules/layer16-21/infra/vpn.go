package infra

import (
	"fmt"

	"github.com/angel-platform/angel/pkg/logger"
)

type VPNManager struct {
	config *InfraConfig
	log    *logger.Logger
}

func NewVPNManager(config *InfraConfig) *VPNManager {
	if config == nil {
		config = DefaultInfraConfig()
	}
	return &VPNManager{
		config: config,
		log:    logger.New("vpn-mgr", logger.LevelInfo),
	}
}

func (v *VPNManager) GenerateWireGuardConfig() (string, error) {
	if v.config.Provider == "" {
		return "", fmt.Errorf("VPN not enabled")
	}

	result := "[Interface]\nPrivateKey = {{.PrivateKey}}\nAddress = {{.Address}}\n"
	if v.config.SSHKeyPath != "" {
		result += fmt.Sprintf("key %s\n", v.config.SSHKeyPath)
	}
	if v.config.SubnetID != "" {
		result += fmt.Sprintf("ca %s\n", v.config.SubnetID)
	}

	v.log.Info("WireGuard config generated for %s:%s", v.config.Provider, v.config.Region)
	return result, nil
}

func (v *VPNManager) GenerateOpenVPNConfig() (string, error) {
	if v.config.Provider == "" {
		return "", fmt.Errorf("VPN not enabled")
	}

	result := "client\ndev tun\nproto udp\nremote %s 1194\n"
	result = fmt.Sprintf(result, v.config.Region)
	if v.config.SSHKeyPath != "" {
		result += fmt.Sprintf("key %s\n", v.config.SSHKeyPath)
	}
	if v.config.VPCID != "" {
		result += fmt.Sprintf("ca %s\n", v.config.VPCID)
	}

	v.log.Info("OpenVPN config generated for %s:%s", v.config.Provider, v.config.Region)
	return result, nil
}
