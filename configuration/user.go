package configuration

import (
   "github.com/cloud-ca/go-cloudca/api"
   "encoding/json"
)

type User struct {
   Id string `json:"id,omitempty"`
   Username string `json:"username,omitempty"`
   Users []User `json:"users"`
   Roles []Role `json:"roles"`
   Organization Organization `json:"tenant,omitempty"`
}


type UserService interface {
   Get(id string) (*User, error)
   List() ([]User, error)
   ListWithOptions(options map[string]string) ([]User, error)
}

type UserApi struct {
   configurationService ConfigurationService
}

func NewUserService(apiClient api.ApiClient) UserService {
   return &UserApi{
      configurationService: NewConfigurationService(apiClient, "users"),
   }
}

func parseUser(data []byte) *User {
   user := User{}
   json.Unmarshal(data, &user)
   return &user
}

func parseUserList(data []byte) []User {
   users := []User{}
   json.Unmarshal(data, &users)
   return users
}

//Get user with the specified id
func (userApi *UserApi) Get(id string) (*User, error) {
   data, err := userApi.configurationService.Get(id, map[string]string{})
   if err != nil {
      return nil, err
   }
   return parseUser(data), nil
}

//List all users
func (userApi *UserApi) List() ([]User, error) {
   return userApi.ListWithOptions(map[string]string{})
}

//List all instances for the current user. Can use options to do sorting and paging.
func (userApi *UserApi) ListWithOptions(options map[string]string) ([]User, error) {
   data, err := userApi.configurationService.List(options)
   if err != nil {
      return nil, err
   }
   return parseUserList(data), nil
}