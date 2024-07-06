package service

import (
	"context"
	"main/db"
	api "main/proto/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserDetails(t *testing.T) {
	userService := NewUserService(db.NewUserRepository())

	tests := map[string]struct {
		req *api.UserRequest
		exp *api.UserResponse
		err error
	}{
		"Valid_user_ID": {
			req: &api.UserRequest{Id: "1"},
			exp: &api.UserResponse{
				User: &api.UserInfo{
					Id:      "1",
					Name:    "Akhil Raj",
					City:    "Thiruvananthapuram",
					Phone:   9876543210,
					Height:  5.800000190734863,
					Married: api.MaritalStatus_YES,
				},
			},
			err: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := userService.GetUserDetails(context.Background(), tc.req)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.exp, res)
		})
	}
}
func TestUserService_FindUsers(t *testing.T) {
	userService := NewUserService(db.NewUserRepository())

	tests := map[string]struct {
		req *api.FindUserRequest
		exp *api.UserListResponse
		err error
	}{
		"Valid_height_filter": {
			req: &api.FindUserRequest{
				PageNumber: 0,
				PageSize:   3,
				Filters: &api.Filter{
					Height: 5.9,
				},
			},
			exp: &api.UserListResponse{
				Users: []*api.UserInfo{
					{
						Id:      "3",
						Name:    "Vishnu Kumar",
						City:    "Kozhikode",
						Phone:   int64(7654321098),
						Height:  5.9,
						Married: api.MaritalStatus_YES,
					},
				},
			},
			err: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := userService.FindUsers(context.Background(), tc.req)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.exp, res)
		})
	}
}

//
