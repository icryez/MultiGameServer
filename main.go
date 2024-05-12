package main

import (
	"fmt"
	"log"
	"net"

	playermodule "github.com/MultiGameServer/PlayerModule"
)

type Message struct {
	from    string
	payload []byte
}
type Server struct {
	listenAddress string
	listner       net.Listener
	quitChan      chan struct{}
	msgChan       chan Message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddress: listenAddr,
		quitChan:      make(chan struct{}),
		msgChan:       make(chan Message, 80),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddress)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.listner = ln
	go s.acceptLoop()
	<-s.quitChan
	close(s.msgChan)
	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listner.Accept()
		if err != nil {
			fmt.Println("Accept Loop Error:", err)
			continue
		}
		fmt.Println("New Connection from:", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Read Error - disconnected from %s : %s", conn.RemoteAddr(), err)
	}
	sessionId := playermodule.GetSessionIdFromBuf(buf[:n])
	playermodule.AllSessions.AddToSession(sessionId, conn.RemoteAddr())
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read Loop Error - disconnected from %s : %s", conn.RemoteAddr(), err)
			break
		}
		matchSession := playermodule.AllSessions.GetSession(sessionId)
		if matchSession.NetAddr[0].String() == conn.RemoteAddr().String(){
			conn.Write([]byte(matchSession.PlayerPos[1]))
		} else {
			conn.Write([]byte(matchSession.PlayerPos[0]))
		}
		s.msgChan <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}
	}
}

func main() {
	server := NewServer(":3000")

	go func() {
		for msg := range server.msgChan {
			fmt.Printf("Recieved message from connection (%s):%s\n", msg.from, string(msg.payload))
		}
	}()

	log.Fatal(server.Start())
}
