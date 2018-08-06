package defs

import (
	"testing"
	"daker.wang/Azen/go-streaming-source-code/api/defs"
	"encoding/json"
	"log"
)

func TestMers(t *testing.T) {
	uc := defs.UserCredential{
		Username:"azen",
		Pwd:"123",
	}

	mer, err := json.Marshal(uc); if err == nil {
		log.Println(string(mer))
	}

	uc2 := &defs.UserCredential{}

	json.Unmarshal(mer, uc2)

	log.Println(uc2)
}
