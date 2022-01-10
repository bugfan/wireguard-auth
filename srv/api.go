package srv

import (
	"encoding/json"
	"net/http"

	"github.com/bugfan/de"
	"github.com/bugfan/logrus"
	"github.com/bugfan/srv"
	"github.com/bugfan/wireguard-auth/env"
	"github.com/bugfan/wireguard-auth/models"
	"github.com/bugfan/wireguard-auth/srv/peer"
	"github.com/bugfan/wireguard-auth/utils"
)

func init() {
	de.SetKey(env.Get("des_key"))
}

type Auth struct {
}

func (*Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*
		api auth middleware
	*/
	// auth := r.Header.Get("WG_TOKEN")
	// _, err := de.DecodeWithBase64([]byte(auth))
	// if auth == "" || err != nil {
	// 	fmt.Println("auth decode error:", err)
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }
}

type ServerConfig struct {
	ListenPort string `json:"wg_listen_port"`
	PrivateKey string `json:"wg_private_key"`
}
type Config struct {
}

func (s *Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conf := &ServerConfig{}
	conf.ListenPort = models.GetValue("wg_listen_port")
	conf.PrivateKey = models.GetValue("wg_private_key")
	data, _ := json.Marshal(conf)
	w.Write(data)
}

type Wireguard struct {
}

func (s *Wireguard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	pubkey := r.Header.Get("WG_KEY")
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
	s.AddHeadHandler(&Auth{})            // set auth middleware
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
