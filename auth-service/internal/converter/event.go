package converter

import (
	"strconv"

	"github.com/AwesomeXjs/registration-service-with-checking-mail/auth-service/internal/model"
)

// ToModelPayload converts topic, user ID, and user information into a Payload model.
func ToModelPayload(topic string, UserID int, userInfo *model.UserInfo) *model.Payload {
	return &model.Payload{
		Topic:   topic,
		Key:     strconv.Itoa(UserID),
		Message: userInfo.Email,
	}
}

// ToModelSendEvent converts an event name and payload into a SendEvent model.
func ToModelSendEvent(event string, payload []byte) *model.SendEvent {
	return &model.SendEvent{
		Event:   event,
		Payload: string(payload),
	}
}
