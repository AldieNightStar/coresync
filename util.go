package coresync

import (
	"encoding/json"
	"strings"
)

func toJson(o any) string {
	sb := strings.Builder{}
	err := json.NewEncoder(&sb).Encode(o)
	if err != nil {
		return ""
	}
	return sb.String()
}
