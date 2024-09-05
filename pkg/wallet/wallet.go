package wallet

import "github.com/google/uuid"

type Wallet struct {
	Id      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"userID"`
	Balance int       `json:"balance"`
}
