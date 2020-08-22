package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`
		got, err := Decode(token)
		want := `{"header":{"alg":"HS256","typ":"JWT"},"payload":{"name":"John Doe","sub":"1234567890","iat":1516239022},"signature":"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`
		assert.NoError(t, err)
		assert.JSONEq(t, want, got)
	})

	t.Run("failed", func(t *testing.T) {
		tests := []struct{ token string }{
			{token: `xxxxx.yyyyy`},
			{token: `xxxxx.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.zzzzz`},
			{token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.yyyyy.zzzzz`},
		}
		for _, test := range tests {
			_, err := Decode(test.token)
			assert.Error(t, err)
		}
	})
}
