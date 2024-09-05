package user

type UserID uint

type User struct {
	Id      UserID `json:"id"`
	Name    string `json:"string"`
	Balance int    `json:"balance"`
}
