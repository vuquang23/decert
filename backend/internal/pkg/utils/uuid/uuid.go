package uuid

import (
	"github.com/google/uuid"
)

func MsgWithUUID(msg string) string {
	uuidString := uuid.New()
	ret := "RequestID: " + uuidString.String() + " " + msg
	return ret
}
