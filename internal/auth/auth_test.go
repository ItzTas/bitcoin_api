package auth

import (
	"net/http"
	"testing"
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
