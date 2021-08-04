package contracts

import "github.com/facilittei/checkout-listener/core/entities"

// Payment contract
type Payment interface {
	Process(message entities.Message) error
}
