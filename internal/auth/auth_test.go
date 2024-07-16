package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		want       string
		wantErrMsg string
	}{
		{
			name:       "valid header",
			headers:    http.Header{"Authorization": {"ApiKey few34t43tg"}},
			want:       "few34t43tg",
			wantErrMsg: "",
		},
		{
			name:       "missing header",
			headers:    http.Header{},
			want:       "",
			wantErrMsg: "no authorization header included",
		},
		{
			name:       "malformed header",
			headers:    http.Header{"Authorization": {"Bearer few34t43tg"}},
			want:       "",
			wantErrMsg: "malformed authorization header",
		},
		{
			name:       "incomplete header",
			headers:    http.Header{"Authorization": {"ApiKey"}},
			want:       "",
			wantErrMsg: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if err != nil {
				if err.Error() != tt.wantErrMsg {
					t.Fatalf("expected error: %v, got: %v", tt.wantErrMsg, err.Error())
				}
			} else {
				if tt.wantErrMsg != "" {
					t.Fatalf("expected error: %v, got no error", tt.wantErrMsg)
				}
			}
			if got != tt.want {
				t.Fatalf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}
