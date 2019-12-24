package Windows

import (
	"github.com/satori/go.uuid"
	"strings"
	"syscall"
)

var userdata = make(map[string]interface{})

func putData(key string, value interface{}) interface{} {
	ov := userdata[key]
	userdata[key] = value
	return ov
}

func getData(key string) interface{} {
	return userdata[key]
}

func sto16(str string) *uint16 {
	ptr, _ := syscall.UTF16PtrFromString(str)
	return ptr
}

func newUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func ifNull(a, b interface{}) interface{} {
	if a == nil {
		return b
	}
	return a
}
