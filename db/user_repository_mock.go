// user_repository_mock.go

package db

// UserRepositoryMock is a mock implementation of UserRepository for testing purposes
type UserRepositoryMock struct {
	Users map[string]Account
}

// GetUserByID retrieves a user by ID from the mock UserRepository
func (u *UserRepositoryMock) GetUserByID(id string) (Account, bool) {
	user, ok := u.Users[id]
	return user, ok
}
