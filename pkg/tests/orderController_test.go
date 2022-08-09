package tests

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type UserRepositoryMock struct {
// 	mock.Mock
// }

// func (r UserRepositoryMock) GetAllUsers() ([]User, error) {
// 	args := r.Called()
// 	users := []User{
// 		{"mock", "*****"},
// 	}
// 	return users, args.Error(1)
// }

// func TestService_GetUser(t *testing.T) {
// 	repository := UserRepositoryMock{}
// 	repository.On("GetAllUsers").Return([]User{}, nil)

// 	service := UserService{repository}
// 	users, _ := service.GetUser()
// 	for i := range users {
// 		assert.Equal(t, users[i].Password, "*****", "user password must be encrypted")
// 	}
// 	fmt.Println(users)
// }
