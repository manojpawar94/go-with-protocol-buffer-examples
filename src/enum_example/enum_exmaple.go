package enum_example

import (
	"log"

	"github.com/manojpawar94/grpc/src/enum_example/enumpb"
)

func NewEnumMessage() *enumpb.EnumMessage {
	em := &enumpb.EnumMessage{Id: 1234, DayOfWeek: enumpb.Day_MONDAY}
	log.Printf("The enum message is %v\n", em)
	return em
}
