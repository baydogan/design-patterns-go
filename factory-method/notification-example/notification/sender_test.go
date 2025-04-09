package notification_test

import (
	"github.com/baydogan/design-patterns/factorymethod-example2/notification"
	"testing"
)

func TestMockSender_NotifyCalled(t *testing.T) {
	mock := &notification.MockSender{}

	mock.Notify("user42", "test message")

	if !mock.Called {
		t.Error("Cannot call Notify")
	}

	if mock.CalledWith.UserID != "user42" {
		t.Errorf("UserID cannot match: awaiting user42, recieved %s", mock.CalledWith.UserID)

	}
}
