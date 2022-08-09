package dto

type CreateUserDto struct {
	Name string `json:"name"`
}

type UpdateUserDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
