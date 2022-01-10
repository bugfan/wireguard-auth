package peer

import (
	"encoding/hex"
	"fmt"

	"github.com/bugfan/wireguard-auth/models"
)

func Store(c *Client) (*models.Peer, error) {
	p := &models.Peer{
		Name:            c.Name,
		Address:         c.Address,
		PrivateKey:      c.KeyPair.PrivateKey,
		PublicKey:       c.KeyPair.PublicKey,
		ListenPort:      c.Interface.ListenPort,
		DNS:             c.Interface.DNS,
		ServerPublicKey: c.Peer.PublicKey,
		Endpoint:        c.Peer.Endpoint,
		AllowedIPs:      c.Peer.AllowedIPs,
		Tag:             hex.EncodeToString([]byte(c.KeyPair.PublicKey)),
	}
	_, err := models.Insert(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetPeer(pubkey string) (*models.Peer, error) {
	p := new(models.Peer)
	has, err := models.GetEngine().Where("public_key=?", pubkey).Get(p)
	if err != nil || !has {
		return nil, err
	}
	return p, nil
}

func GetSetting(key string) string {
	set := new(models.Setting)
	has, err := models.GetEngine().Where("key=?", key).Get(set)
	if err != nil || !has {
		fmt.Println("not found key ,error is:", err)
		return ""
	}
	return set.Value
}
