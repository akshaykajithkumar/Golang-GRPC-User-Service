package db

import (
	"sync"
)

// UserRecords represents a slice of pre-defined user data
var UserRecords = []Account{
	{ID: "1", FullName: "Akhil Raj", Location: "Thiruvananthapuram", Contact: 9876543210, Height: 5.8, IsMarried: "TRUE"},
	{ID: "2", FullName: "Lakshmi Nair", Location: "Kochi", Contact: 8765432109, Height: 5.5, IsMarried: "FALSE"},
	{ID: "3", FullName: "Vishnu Kumar", Location: "Kozhikode", Contact: 7654321098, Height: 5.9, IsMarried: "TRUE"},
	{ID: "4", FullName: "Anjali Menon", Location: "Thrissur", Contact: 6543210987, Height: 5.6, IsMarried: "FALSE"},
	{ID: "5", FullName: "Ravi Shankar", Location: "Alappuzha", Contact: 5432109876, Height: 5.7, IsMarried: "TRUE"},
	{ID: "6", FullName: "Divya Suresh", Location: "Palakkad", Contact: 4321098765, Height: 5.4, IsMarried: "FALSE"},
	{ID: "7", FullName: "Manoj Varma", Location: "Kollam", Contact: 3210987654, Height: 5.8, IsMarried: "TRUE"},
	{ID: "8", FullName: "Meera Mohan", Location: "Kannur", Contact: 2109876543, Height: 5.5, IsMarried: "TRUE"},
}

// UserRepository is responsible for managing user data operations
type UserRepository struct {
	userMap sync.Map // Using sync.map for concurrent-safe access to user data
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository() *UserRepository {
	repo := &UserRepository{}
	for _, user := range UserRecords {
		// Creating a new instance of UserAccount for each entry
		userCopy := user
		repo.userMap.Store(user.ID, &userCopy)
	}
	return repo
}

// GetUserByID retrieves a user by ID
func (repo *UserRepository) GetUserByID(id string) (*Account, bool) {
	// Load user by ID from the userMap
	user, ok := repo.userMap.Load(id)
	if !ok {
		return nil, false
	}

	// Cast the loaded user to *Account pointer
	userPtr, ok := user.(*Account)
	if !ok {
		return nil, false
	}

	return userPtr, true
}

// GetUsersByIDs retrieves users by a slice of IDs
func (repo *UserRepository) GetUsersByIDs(ids []string) []*Account {
	var users []*Account
	var wg sync.WaitGroup
	userChan := make(chan *Account, len(ids))

	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if user, ok := repo.GetUserByID(id); ok {
				userChan <- user
			}
		}(id)
	}

	go func() {
		wg.Wait()
		close(userChan)
	}()

	for user := range userChan {
		users = append(users, user)
	}

	return users
}
