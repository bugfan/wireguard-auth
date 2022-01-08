package models

import (
	"os"

	"github.com/bugfan/wireguard-auth/env"
	"github.com/sirupsen/logrus"
)

func init() {
	Register(&Setting{})
	Register(&Peer{})

	_, err := SetEngine(&Config{
		User:     env.Get("db_user"),
		Password: env.Get("db_pwd"),
		Host:     env.Get("db_host"),
		Name:     env.Get("db_name"),
		ShowSQL:  env.GetBool("db_show_sql"),
	}, env.Get("db_scheme"))
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}
