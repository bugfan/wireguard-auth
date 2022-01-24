package srv

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bugfan/de"
	"github.com/bugfan/srv"
	"github.com/bugfan/wireguard-auth/env"
	"github.com/bugfan/wireguard-auth/models"
	"github.com/bugfan/wireguard-auth/srv/peer"
	"github.com/bugfan/wireguard-auth/utils"
	"github.com/sirupsen/logrus"
)

func init() {
	de.SetKey(env.Get("des_key"))
}

func VerifyAuth(w http.ResponseWriter, r *http.Request) error {
	/*
		api auth middleware
	*/
	auth := r.Header.Get("Wgtoken")
	fmt.Println("auth is:", auth)
	_, err := de.DecodeWithBase64([]byte(auth))
	if auth == "" || err != nil {
		return errors.New("auth decode error")
	}
	return nil
}

type ServerConfig struct {
	ListenPort string `json:"wg_listen_port"`
	PrivateKey string `json:"wg_private_key"`
}
type Config struct {
}

func (s *Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conf := &ServerConfig{}
	serverIP := r.Header.Get("Wgserverip")
	logrus.Infof("[server-init] server ip %s\n", serverIP)
	conf.ListenPort = models.GetValue("wg_listen_port")
	conf.PrivateKey = models.GetValue("wg_private_key")
	data, _ := json.Marshal(conf)
	w.Write(data)
}

type Wireguard struct {
}

func (s *Wireguard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := VerifyAuth(w, r); err != nil {
		w.WriteHeader(403)
		return
	}
	switch r.Method {
	case http.MethodGet:
		s.verifyAuth(w, r)
		return
	case http.MethodPost:
		s.createPeer(w, r)
		return
	default:
	}
	http.NotFound(w, r)
}
func (*Wireguard) verifyAuth(w http.ResponseWriter, r *http.Request) {
	// get
	pubkey := r.Header.Get("Wgkey")
	fmt.Println("custom key is:", pubkey)
	serverIP := r.Header.Get("WgServerIp")
	logrus.Infof("[authentication] server ip %s,client publick-key is %s\n", serverIP, pubkey)
	data, err := peer.GetPeer(pubkey)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	bs, _ := utils.IJSON.Marshal(data)
	w.Write(bs)
}
func (*Wireguard) createPeer(w http.ResponseWriter, r *http.Request) {
	// post
	name := r.URL.Query().Get("name")
	client := peer.F.Make(name)
	data, err := peer.Store(client)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bs, _ := utils.IJSON.Marshal(data)
	w.WriteHeader(http.StatusCreated)
	w.Write(bs)
}

func NewServer(addr string) *Server {

	s := srv.New(addr)
	s.Handle("/wireguard", &Wireguard{}) // set wg handler
	s.Handle("/config", &Config{})
	return &Server{
		addr,
		s,
	}
}

type Server struct {
	addr string
	*srv.Server
}

func (s *Server) Run() {
	logrus.Fatal(s.Server.Run())
}
