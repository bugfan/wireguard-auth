package main

import (
	"fmt"
	"log"

	_ "github.com/bugfan/wireguard-auth/models"
	"github.com/bugfan/wireguard-auth/srv/dhcp"
	"github.com/bugfan/wireguard-auth/srv/key"
)

func main() {

	key, err := key.GeneratePrivateKey()
	if err != nil {
		log.Fatal("Error encountered while generating private key")
	}
	privateKey := key.String()
	publicKey := key.PublicKey().String()
	fmt.Println("333:", privateKey, publicKey)

	ip := dhcp.GetCIDR()
	fmt.Println(22, ip)
}
