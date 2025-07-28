package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintUserProfile_AllFields(t *testing.T) {
	user := map[string]interface{}{
		"id":         "u1",
		"username":   "tester",
		"email":      "test@example.com",
		"name":       "Test User",
		"created_at": "2024-01-01",
		"updated_at": "2024-02-01",
	}

	var buf bytes.Buffer
	printUserProfile(&buf, user)

	out := buf.String()
	assert.Contains(t, out, "User Profile:")
	assert.Contains(t, out, "ID: u1")
	assert.Contains(t, out, "Username: tester")
	assert.Contains(t, out, "Email: test@example.com")
	assert.Contains(t, out, "Display Name: Test User")
	assert.Contains(t, out, "Created: 2024-01-01")
	assert.Contains(t, out, "Updated: 2024-02-01")
}

func TestPrintUserProfile_MissingFields(t *testing.T) {
	user := map[string]interface{}{
		"id":       "u2",
		"username": "tester2",
	}

	var buf bytes.Buffer
	printUserProfile(&buf, user)

	out := buf.String()
	assert.Contains(t, out, "ID: u2")
	assert.Contains(t, out, "Username: tester2")
	assert.NotContains(t, out, "Email:")
	assert.NotContains(t, out, "Display Name:")
}
