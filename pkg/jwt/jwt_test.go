package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tests := []struct {
			token string
			want  string
		}{
			{
				token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c`,
				want:  `{"header":{"alg":"HS256","typ":"JWT"},"payload":{"name":"John Doe","sub":"1234567890","iat":1516239022},"signature":"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}`,
			},
			{
				token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9l8J-OgyIsImlhdCI6MTUxNjIzOTAyMn0.LJEPu1YqnXWREF3ooCnu-jIKS-fesGumsPC4xPjZ68Y`,
				want:  `{"header":{"alg":"HS256","typ":"JWT"},"payload":{"name":"John Doe🎃","sub":"1234567890","iat":1516239022},"signature":"LJEPu1YqnXWREF3ooCnu-jIKS-fesGumsPC4xPjZ68Y"}`,
			},
		}
		for _, test := range tests {
			got, err := Decode(test.token)
			assert.NoError(t, err)
			assert.JSONEq(t, test.want, got)
		}
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
