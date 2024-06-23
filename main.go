package main

import (
	"log/slog"
	"net"
)

const defaultListenAddr = ":3000"
type Config struct {
	ListenAddr string
}
type Server struct {
	Config
	peers map[*Peer]bool
	ln    net.Listener
}

func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListenAddr
	}
	return &Server{
		Config: cfg,
		peers:  make(map[*Peer]bool),
	}
}
// this function starts the server
// then the loop is accepted
// the loop then lets the connection to be handled
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln

	return s.acceptLoop()
}

func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accept error", "err", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {

}

func main() {

}
