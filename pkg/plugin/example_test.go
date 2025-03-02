package plugin

import (
	"fmt"

	plug "github.com/hashicorp/go-plugin"
)

// ExamplePlugin shows how to create and use a plugin
func Example() {
	// Define your plugin implementation and create plugin map
	pluginMap := map[string]plug.Plugin{
		"example": &BasePlugin{},
	}

	// Create new plugin client
	client := NewPluginClient("./path/to/plugin", pluginMap)

	// Connect to new plugin instance
	if err := client.Connect(nil); err != nil {
		panic(err)
	}
	defer client.Disconnect()

	// Get plugin instance
	_, err := client.GetPlugin("example")
	if err != nil {
		panic(err)
	}

	// Get reattach config for future connections
	config, err := client.ReattachConfig()
	if err != nil {
		panic(err)
	}

	// Later, reconnect using the reattach config
	newClient := NewPluginClient("./path/to/plugin", pluginMap)
	if err := newClient.Connect(config); err != nil {
		panic(err)
	}
	defer newClient.Disconnect()

	fmt.Println("Plugin connected successfully")
}
