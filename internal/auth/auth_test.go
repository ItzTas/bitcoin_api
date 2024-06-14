package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
)

func TestGetBearerToken(t *testing.T) {
	const authorizationHeader = "Authorization"
	type args struct {
		header http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "No error",
			args: args{
				header: http.Header{
					authorizationHeader: []string{"Bearer first"},
				},
			},
			want:    "first",
			wantErr: false,
		},
		{
			name: "No bearer",
			args: args{
				header: http.Header{
					authorizationHeader: []string{"badly formated"},
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid length",
			args: args{
				header: http.Header{
					authorizationHeader: []string{"Bearer token badly"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBearerToken(tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBearerToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIDByToken(t *testing.T) {
	const secret = "testSecret"
	id1 := uuid.New()
	token1, err := NewJWT(database.User{
		ID: id1,
	}, secret, time.Hour)
	if err != nil {
		t.Error(err)
	}
	type args struct {
		token     string
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test one correct",
			args: args{
				token:     token1,
				secretKey: secret,
			},
			want:    id1.String(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIDByToken(tt.args.token, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIDByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIDByToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
