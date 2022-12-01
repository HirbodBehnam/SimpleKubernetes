package master

import (
	log "github.com/sirupsen/logrus"
	"net"
)

// Server of master
type Server struct {
	// Users which can access this master
	Users map[string]string
	// List of all slaves
}

// RunServer will run a server for client connections and another one
// for slave connections.
func (s *Server) RunServer(clientListen, slaveListen string) {
	go s.runClientServer(clientListen)
	go s.runSlaveServer(slaveListen)
	select {}
}

// Runs the server which clients connect to and batch the jobs in it
func (s *Server) runClientServer(listenAddress string) {
	// Listen for client
	l, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.WithError(err).Fatalln("cannot listen for client's server")
	}
	// Wait for connections
	for {
		client, err := l.Accept()
		if err != nil {
			log.WithError(err).Error("cannot accept clients connection")
			continue
		}
		// Handle this client
		s.handleClient(client)
	}
}

// Runs the server which slaves will connect to it
func (s *Server) runSlaveServer(listenAddress string) {

}
