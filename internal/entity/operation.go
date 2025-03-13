package entity

import (
	"time"

	"github.com/google/uuid"
)

type Operation struct {
	AccountId     uuid.UUID `db:"account_id"`
	Amount        int       `db:"amount"`
	OperationType string    `db:"operation_type"`
	CreatedAt     time.Time `db:"created_at"`
}

const (
	OperationTypeDeposit      = "deposit"
	OperationTypeWithdraw     = "withdraw"
	OperationTypeTransferFrom = "transfer_from"
	OperationTypeTransferTo   = "transfer_to"
)
