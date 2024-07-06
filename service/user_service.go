package service

import (
	"context"
	"fmt"
	"main/db"
	api "main/proto/api"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService implements the UserService gRPC interface
type UserService struct {
	api.UnimplementedUserServiceServer
	userRepo *db.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepo *db.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetUserDetails retrieves user details by ID
func (s *UserService) GetUserDetails(ctx context.Context, req *api.UserRequest) (*api.UserResponse, error) {
	if len(req.Id) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid user ID")
	}

	// Retrieves user by ID using repository method
	user, ok := s.userRepo.GetUserByID(req.Id)
	if !ok {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	// Converts user details to proto format and return
	userProto := convertToProto(user)
	return &api.UserResponse{User: userProto}, nil
}

// GetUserList retrieves a list of users based on IDs
func (s *UserService) GetUserList(ctx context.Context, req *api.UserListRequest) (*api.UserListResponse, error) {
	var users []*api.UserInfo
	var wg sync.WaitGroup
	userChan := make(chan *api.UserInfo, len(req.Ids))

	// Retrieves user info asynchronously
	for _, id := range req.Ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if user, ok := s.userRepo.GetUserByID(id); ok {
				userProto := convertToProto(user)
				userChan <- userProto
			}
		}(id)
	}

	go func() {
		wg.Wait()
		close(userChan)
	}()

	// Collects user info from channels
	for userProto := range userChan {
		users = append(users, userProto)
	}

	// Paginat the user list
	return paginateUsers(users, req.PageNumber, req.PageSize), nil
}

// FindUsers retrieves users based on specific criteria
func (s *UserService) FindUsers(ctx context.Context, req *api.FindUserRequest) (*api.UserListResponse, error) {
	filteredUsersChan := make(chan *api.UserInfo)
	var wg sync.WaitGroup

	// Worker function to filter users based on criteria
	filterWorker := func(users []db.Account) {
		defer wg.Done()
		for _, user := range users {
			if (req.Filters.Id == "" || req.Filters.Id == user.ID) &&
				(req.Filters.Name == "" || req.Filters.Name == user.FullName) &&
				(req.Filters.City == "" || req.Filters.City == user.Location) &&
				(req.Filters.Phone == 0 || req.Filters.Phone == user.Contact) &&
				(req.Filters.Height == 0 || req.Filters.Height == user.Height) &&
				(req.Filters.Married == api.MaritalStatus_UNKNOWN || req.Filters.Married == convertMaritalStatusToProto(user.IsMarried)) {
				filteredUsersChan <- convertToProto(&user)
			}
		}
	}

	// Number of workers and chunk size for splitting user records
	numWorkers := 4
	chunkSize := (len(db.UserRecords) + numWorkers - 1) / numWorkers

	userSlice := make([]db.Account, 0, len(db.UserRecords))
	userSlice = append(userSlice, db.UserRecords...)

	for i := 0; i < len(userSlice); i += chunkSize {
		end := i + chunkSize
		if end > len(userSlice) {
			end = len(userSlice)
		}
		wg.Add(1)
		go filterWorker(userSlice[i:end])
	}

	go func() {
		wg.Wait()
		close(filteredUsersChan)
	}()

	var userProtos []*api.UserInfo
	for userProto := range filteredUsersChan {
		userProtos = append(userProtos, userProto)
	}

	return paginateUsers(userProtos, req.PageNumber, req.PageSize), nil
}

// paginateUsers paginates a list of users based on page number and page size
func paginateUsers(users []*api.UserInfo, pageNumber, pageSize int32) *api.UserListResponse {
	start := pageNumber * pageSize
	end := start + pageSize

	fmt.Println("Start:", start)
	fmt.Println("End:", end)

	if int(start) >= len(users) {
		return &api.UserListResponse{Users: []*api.UserInfo{}}
	}

	if int(end) > len(users) {
		end = int32(len(users))
	}

	return &api.UserListResponse{Users: users[start:end]}
}

// convertToProto converts user details to proto format
func convertToProto(user *db.Account) *api.UserInfo {
	maritalStatus := api.MaritalStatus_UNKNOWN
	if user.IsMarried == "TRUE" {
		maritalStatus = api.MaritalStatus_YES
	} else if user.IsMarried == "FALSE" {
		maritalStatus = api.MaritalStatus_NO
	}

	return &api.UserInfo{
		Id:      user.ID,
		Name:    user.FullName,
		City:    user.Location,
		Phone:   user.Contact,
		Height:  user.Height,
		Married: maritalStatus,
	}
}

// convertMaritalStatusToProto converts string marital status to proto enum
func convertMaritalStatusToProto(maritalStatus string) api.MaritalStatus {
	switch maritalStatus {
	case "TRUE":
		return api.MaritalStatus_YES
	case "FALSE":
		return api.MaritalStatus_NO
	default:
		return api.MaritalStatus_UNKNOWN
	}
}
