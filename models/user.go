package models

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}
