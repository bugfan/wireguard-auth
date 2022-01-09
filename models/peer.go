package models

import (
	"encoding/json"
	"time"
)

type Peer struct {
	ID   int64
	Name string
	// interface
	Address    string
	PrivateKey string
	PublicKey  string
	ListenPort string
	DNS        string
	// peer
	ServerPublicKey string `xorm:"index"`
	Endpoint        string
	AllowedIPs      string `xorm:"allow_ips"`
	// other
	Additive json.RawMessage
	Tag      string
	Created  time.Time `xorm:"CREATED"`
	Updated  time.Time `xorm:"UPDATED"`
	Deleted  time.Time `xorm:"deleted"`
}
