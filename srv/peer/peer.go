package peer

import (
	"github.com/bugfan/wireguard-auth/env"
	"github.com/bugfan/wireguard-auth/models"
	"github.com/bugfan/wireguard-auth/srv/key"
)

type Factory interface {
	Make(string) *Client
}

var (
	F                   Factory
	WireguarsPublicKey  string
	WireguarsPrivateKey string
)

func init() {
	initWireguardConfig()
	F = NewGeneral(
		env.Get("wg_listen_port"),
		env.Get("wg_mtu"),
		env.Get("wg_dns"),
		WireguarsPublicKey,
		env.Get("wg_psk"),
		env.Get("wg_allow_ips"),
		env.Get("wg_keepalive"),
	)
}

func initWireguardConfig() {
	WireguarsPublicKey = env.Get("wg_public_key")
	WireguarsPrivateKey = env.Get("wg_private_key")
	if (WireguarsPublicKey == "" || WireguarsPrivateKey == "") && models.GetValue("wg_public_key") == "" {
		key, _ := key.GeneratePrivateKey()
		WireguarsPrivateKey = key.String()
		WireguarsPublicKey = key.PublicKey().String()
		models.Insert(&models.Setting{
			Key:   "wg_public_key",
			Value: WireguarsPublicKey,
		}, &models.Setting{
			Key:   "wg_private_key",
			Value: WireguarsPrivateKey,
		})
	}
	WireguarsPublicKey = models.GetValue("wg_public_key")
	WireguarsPrivateKey = models.GetValue("wg_private_key")
	models.InitSetting(map[string]string{
		"wg_address":     env.Get("wg_address"),
		"wg_listen_port": env.Get("wg_listen_port"),
	})
}

func NewGeneral(listenPort, mtu, dns, pk, psk, ips, keepalive string) Factory {
	return &General{
		ListenPort:          listenPort,
		MTU:                 mtu,
		DNS:                 dns,
		PublicKey:           pk,
		AllowedIPs:          ips,
		PersistentKeepalive: keepalive,
	}
}

type General struct {
	ListenPort          string
	MTU                 string
	DNS                 string
	PublicKey           string
	PSK                 string
	AllowedIPs          string
	PersistentKeepalive string
}

func (s *General) Make(name string) *Client {
	cli := NewClient(name, s)
	cli.Init()
	return cli
}
