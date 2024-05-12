package playermodule

import (
	"net"
	"strings"
	"sync"
)

type Sessions struct {
	sessions         map[string]MatchSession
	numberOfSessions int
	sync.Mutex
}

type MatchSession struct {
	NetAddr [2]net.Addr
	PlayerPos [2]string
}

var AllSessions Sessions

func GenAllSessions(){

		AllSessions = *new(Sessions)
		AllSessions.sessions = make(map[string]MatchSession)
}

func (s *Sessions) AddToSession(sessionId string, playerAddr net.Addr) {
	s.Lock()
	defer s.Unlock()
	s.Unlock()
	entry := s.GetSession(sessionId)
	s.Lock()
	if entry.NetAddr[0] == nil {
		entry.NetAddr[0] = playerAddr
	} else {
		entry.NetAddr[1] = playerAddr
	}
	s.sessions[sessionId] = entry
}

func (s *Sessions) GetSession(sessionId string) MatchSession {
	s.Lock()
	defer s.Unlock()
	if entry, ok := s.sessions[sessionId]; ok {
		return entry
	} else {
		return *new(MatchSession)
	}
}

func GetSessionIdFromBuf(buf []byte) string{
	str := string(buf)
	str = strings.TrimSpace(str[0:5])
	return str
}
