package peer

import (
	"github.com/bugfan/wireguard-auth/srv/dhcp"
	"github.com/bugfan/wireguard-auth/srv/key"
	"github.com/sirupsen/logrus"
)

func NewClient(name string, g *General) *Client {
	return &Client{
		Name: name,
		Interface: Interface{
			ListenPort: g.ListenPort,
			MTU:        g.MTU,
			DNS:        g.DNS,
		},
		Peer: Peer{
			PublicKey:           g.PublicKey,
			PSK:                 g.PSK,
			AllowedIPs:          g.AllowedIPs,
			PersistentKeepalive: g.PersistentKeepalive,
		},
	}
}

type KeyPair struct {
	PrivateKey string
	PublicKey  string
}

type Interface struct {
	ListenPort string
	Address    string
	MTU        string
	DNS        string
}

type Peer struct {
	PublicKey           string
	PSK                 string
	Endpoint            string //variable
	AllowedIPs          string
	PersistentKeepalive string
}

type Client struct {
	Name string
	KeyPair
	Interface
	Peer
}

func (c *Client) Init() {
	// gen keypair
	key, err := key.GeneratePrivateKey()
	if err != nil {
		logrus.Error("Error encountered while generating private key")
	} else {
		c.KeyPair.PrivateKey = key.String()
		c.KeyPair.PublicKey = key.PublicKey().String()
	}
	// gen cidr
	c.Interface.Address = dhcp.GetCIDR()
}
