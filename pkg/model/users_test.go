package model

import (
	"encoding/json"
	"testing"
)

func TestContact_TableName(t *testing.T) {
	u := Email{}
	dataByes, _ := json.Marshal(u)
	t.Log(string(dataByes))
}
