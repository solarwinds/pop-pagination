package pagination

import (
	"testing"
	"encoding/base64"
)

func Test_createTokenString(t *testing.T) {
	t.Run("ensure createTokenString returns base64url encoded string", func(t *testing.T) {
		testStr := createTokenString()
		_, err := base64.RawURLEncoding.DecodeString(testStr)
		if err != nil {
			t.Error("Could not decode token string", err)
		}
	})
}
