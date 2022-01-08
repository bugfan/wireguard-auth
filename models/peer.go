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
	ListenPort int64
	DNS        string
	// peer
	PublicKey  string
	Endpoint   string
	AllowedIPs string `xorm:"allow_ips"`
	// other
	Additive json.RawMessage
	Created  time.Time `xorm:"CREATED"`
	Updated  time.Time `xorm:"UPDATED"`
	Deleted  time.Time `xorm:"deleted"`
}
