package auth

import (
    "errors"
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name    string
        headers http.Header
        want    string
        wantErr error
    }{
        {
            name:    "valid header",
            headers: http.Header{"Authorization": {"ApiKey few34t43tg"}},
            want:    "few34t43tg",
            wantErr: nil,
        },
        {
            name:    "missing header",
            headers: http.Header{},
            want:    "",
            wantErr: ErrNoAuthHeaderIncluded,
        },
        {
            name:    "malformed header",
            headers: http.Header{"Authorization": {"Bearer few34t43tg"}},
            want:    "",
            wantErr: errors.New("malformed authorization header"),
        },
        {
            name:    "incomplete header",
            headers: http.Header{"Authorization": {"ApiKey"}},
            want:    "",
            wantErr: errors.New("malformed authorization header"),
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := GetAPIKey(tt.headers)
            if !errors.Is(err, tt.wantErr) {
                t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
            }
            if got != tt.want {
                t.Fatalf("expected: %v, got: %v", tt.want, got)
            }
        })
    }
}
