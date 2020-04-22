package uuid

import (
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	uid := uuid.NewV4()
	return uid.String()
}