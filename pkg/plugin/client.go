package plugin

import (
	"fmt"
	"os/exec"

	plug "github.com/hashicorp/go-plugin"
)

// PluginClient manages plugin lifecycle and connections
type PluginClient struct {
	config     *plug.ClientConfig
	client     *plug.Client
	pluginMap  map[string]plug.Plugin
	pluginPath string
}

// NewPluginClient creates a new plugin client instance
func NewPluginClient(path string, pluginMap map[string]plug.Plugin) *PluginClient {
	return &PluginClient{
		pluginPath: path,
		pluginMap:  pluginMap,
	}
}

// Connect starts a new plugin process or attaches to an existing one
func (pc *PluginClient) Connect(reattach *plug.ReattachConfig) error {
	pc.config = &plug.ClientConfig{
		HandshakeConfig: plug.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "BOUTIQUE_PLUGIN",
			MagicCookieValue: "boutique",
		},
		Plugins:  pc.pluginMap,
		Cmd:      exec.Command(pc.pluginPath),
		Reattach: reattach,
		AllowedProtocols: []plug.Protocol{
			plug.ProtocolGRPC,
		},
	}

	pc.client = plug.NewClient(pc.config)
	return nil
}

// Disconnect terminates the plugin connection
func (pc *PluginClient) Disconnect() error {
	if pc.client != nil {
		pc.client.Kill()
	}
	return nil
}

// GetPlugin retrieves a plugin instance by name
func (pc *PluginClient) GetPlugin(name string) (interface{}, error) {
	if pc.client == nil {
		return nil, fmt.Errorf("client not connected")
	}

	rpcClient, err := pc.client.Client()
	if err != nil {
		return nil, fmt.Errorf("failed to create rpc client: %w", err)
	}

	raw, err := rpcClient.Dispense(name)
	if err != nil {
		return nil, fmt.Errorf("failed to dispense plugin: %w", err)
	}

	return raw, nil
}

// ReattachConfig returns the configuration needed to reattach to this plugin
func (pc *PluginClient) ReattachConfig() (*plug.ReattachConfig, error) {
	if pc.client == nil {
		return nil, fmt.Errorf("client not connected")
	}
	return pc.client.ReattachConfig(), nil
}
