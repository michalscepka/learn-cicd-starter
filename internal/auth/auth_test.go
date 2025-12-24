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
		wantErr bool
		err     error
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			want:    "my-secret-key",
			wantErr: false,
		},
		{
			name:    "no auth header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
			err:     ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed auth header",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-key"},
			},
			want:    "",
			wantErr: true,
			err:     errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("GetAPIKey() error = %v, want %v", err, tt.err)
				}
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
