package transfer

import "github.com/jbiers/boring-api/pkg/user"

type TransferId uint

type Transfer struct {
	Id       TransferId  `json:"id"`
	Amount   int         `json:"amount"`
	Sender   user.UserID `json:"sender"`
	Receiver user.UserID `json:"receiver"`
}
