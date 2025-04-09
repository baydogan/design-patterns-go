package notification

import "fmt"

type MockSender struct {
	Called     bool
	CalledWith struct {
		UserID  string
		Message string
	}
}

func (m *MockSender) Notify(userID string, message string) {
	m.Called = true
	m.CalledWith.UserID = userID
	m.CalledWith.Message = message
	fmt.Println("[Mock] Notify called")
}
