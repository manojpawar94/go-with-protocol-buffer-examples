package enum_example

import (
	"testing"

	"github.com/manojpawar94/grpc/src/enum_example/enumpb"
)

func TestNewEnumMessage(t *testing.T) {
	em := NewEnumMessage()

	if em.GetDayOfWeek() != enumpb.Day_MONDAY {
		t.Failed()
	}
}
