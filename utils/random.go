package utils

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func Uuid() string {
	uuid := uuid.New()
	return uuid.String()
}

func GenerateId() string {
	id := time.Now().UTC().Format("20060102150405.000000000")
	return strings.ReplaceAll(id, ".", "")
}
func GenerateTime() int64 {
	time := time.Now().UnixNano() / 100000
	return time
}
