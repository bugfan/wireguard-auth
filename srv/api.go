package srv

import (
	"encoding/hex"
	"net/http"

	"github.com/bugfan/logrus"
	"github.com/bugfan/srv"
	"github.com/bugfan/wireguard-auth/srv/peer"
	"github.com/bugfan/wireguard-auth/utils"
)

type Auth struct {
}

func (*Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*
		todo
		api auth middleware
	*/
}

type Config struct {
}

func (s *Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
	pubkey, err := hex.DecodeString(r.URL.Query().Get("publickey"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	data, err := peer.GetPeer(string(pubkey))
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
