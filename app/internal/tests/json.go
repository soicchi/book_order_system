package tests

import (
	"encoding/json"
	"testing"
)

// convert response bytes to map
func UnmarshalResponse(t *testing.T, resBody []byte) map[string]interface{} {
	var res map[string]interface{}
	json.Unmarshal(resBody, &res)

	return res
}
