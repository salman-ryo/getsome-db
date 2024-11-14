package db

import (
	"errors"
	"sync"
)

type Session struct {
	Database *Database // Active database instance for this session
}

type SessionManager struct {
	sessions map[string]*Session //Map of Session Ids to a session
	mu       sync.Mutex
}

func NewsessionManager() *SessionManager {
	return &SessionManager{
		sessions: make(map[string]*Session),
	}
}

// CreateSession method for Session Manager initializes a new session with a given database name.
func (sm *SessionManager) CreateSession(dbName string) (*Session, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Check if session already exists
	if session, exists := sm.sessions[dbName]; exists {
		return session, nil
	}

	// Create a new database instance and session
	dbInstance, err := NewDatabase("/storage" + dbName + ".json")
	if err != nil {
		return nil, err
	}

	mySession := &Session{Database: dbInstance}
	sm.sessions[dbName] = mySession
	return mySession, nil
}

// GetSession method retrieves an existing session by database name.
func (sm *SessionManager) GetSession(dbName string) (*Session, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	session, exists := sm.sessions[dbName]
	if !exists {
		return nil, errors.New("session not found")
	}

	return session, nil
}

// CloseSession method removes a session, optionally saving its data.

func (sm *SessionManager) CloseSession(dbName string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	_, exists := sm.sessions[dbName]
	if !exists {
		return errors.New("session not found")
	}
	delete(sm.sessions, dbName)
	return nil
}
