package shared

import "github.com/hashicorp/go-plugin"

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BOUTIQUE_PLUGIN",
	MagicCookieValue: "boutique",
}
