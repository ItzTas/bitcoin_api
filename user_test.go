package main

import (
	"reflect"
	"testing"

	"github.com/ItzTas/bitcoinAPI/internal/database"
	"github.com/google/uuid"
)

func Test_databaseUserToUser(t *testing.T) {
	type args struct {
		user database.User
	}

	id := uuid.New()
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "Test1",
			args: args{
				user: database.User{
					ID:       id,
					UserName: "testName",
					Email:    "exemplo@email",
					Password: "123",
					Currency: "20",
				},
			},
			want: User{
				ID:       id,
				UserName: "testName",
				Email:    "exemplo@email",
				Password: "123",
				Currency: "20",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := databaseUserToUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("databaseUserToUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
