package transfer

import (
	"github.com/google/uuid"
)

type Transfer struct {
	Id       uuid.UUID `json:"id"`
	Amount   int64     `json:"amount"`
	Sender   uuid.UUID `json:"sender"`
	Receiver uuid.UUID `json:"receiver"`
}
