package cmd

import (
	"bytes"
	"encoding/json"
	"testing"
)

func FuzzPrintUserProfile(f *testing.F) {
	seeds := []string{
		`{"id":"1","username":"u"}`,
		`{}`,
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, input string) {
		var user map[string]interface{}
		_ = json.Unmarshal([]byte(input), &user)
		var buf bytes.Buffer
		printUserProfile(&buf, user)
	})
}
