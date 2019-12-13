package main

import (
	"log"
	"sync"
)

type sessionStore struct {
	sessions map[string]*session
	sync.Mutex
}

func newSessionStore() *sessionStore {
	return &sessionStore{
		sessions: map[string]*session{},
	}
}

func (ss *sessionStore) register(sess *session) {
	ss.Lock()
	defer ss.Unlock()

	ss.sessions[sess.uid] = sess

	log.Printf("Registered session %s", sess.uid)
}

func (ss *sessionStore) unregister(sess *session) {
	ss.Lock()
	defer ss.Unlock()

	delete(ss.sessions, sess.uid)

	log.Printf("Unegistered session %s", sess.uid)
}

func (ss sessionStore) getSessions() []session {
	ss.Lock()
	defer ss.Unlock()

	sessions := []session{}
	for _, s := range ss.sessions {
		sessions = append(sessions, *s)
	}

	return sessions
}
