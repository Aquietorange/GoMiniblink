package Utils

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func NewUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
