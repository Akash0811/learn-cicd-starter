package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestApiKeyHeaderValue(t *testing.T) {
	headers := http.Header{}
	authKey := "Random"
	headers.Add("Authorization", fmt.Sprintf("ApiKey %s", authKey))

	gotAuthKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no errors but got: %v\n", err)
	}

	if gotAuthKey != authKey {
		t.Fatalf("expected: %v, got: %v\n", gotAuthKey, authKey)
	}
}

func TestApiKeyHeaderSingleValue(t *testing.T) {
	headers := http.Header{}
	authKey := "ApiKey"
	headers.Add("Authorization", authKey)

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("expected errors %v but got none\n", err)
	}
}
