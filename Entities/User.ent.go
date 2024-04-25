package Entities

type User struct {
	GUID     string `json:"guid"`
	IsActive bool   `json:"isActive"`
	Balance  string `json:"balance"`
	Picture  string `json:"picture"`
	Age      int    `json:"age"`
	EyeColor string `json:"eyeColor"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
