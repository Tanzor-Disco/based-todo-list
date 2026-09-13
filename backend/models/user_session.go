package models

type SessionID int

type UserSession struct {
	ID            SessionID
	UserID        UserID
	SessionString string
}

func NewUserSession(userID UserID, sessionString string) UserSession {
	return UserSession{
		UserID:        userID,
		SessionString: sessionString,
	}
}
