package wyre

import (
	"strconv"
	"time"
)

func normalizePath(path string) string {
	if path[0] == '/' {
		return path
	}
	return "/" + path
}

func getTimestampString() string {
	now := time.Now().UnixNano() / 10000
	return strconv.FormatInt(now, 10)
}
