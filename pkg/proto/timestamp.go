package proto

import (
	"time"

	"github.com/evgenivanovi/gophkeeper/api/pb/common"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/* __________________________________________________ */

// NewTimestampFromTime ...
func NewTimestampFromTime(value *time.Time) *timestamp.Timestamp {
	if value == nil {
		return nil
	}
	return timestamppb.New(*value)
}

/* __________________________________________________ */

func NewTimeFromTimestamp(value *timestamp.Timestamp) *time.Time {

	if value == nil {
		return nil
	}

	val := value.AsTime()
	return &val

}

func NewTimeFromOptionalTimestamp(value *common.OptionalTimestamp) *time.Time {
	switch kind := value.GetKind().(type) {
	case nil:
		return nil
	case *common.OptionalTimestamp_Null:
		return nil
	case *common.OptionalTimestamp_Data:
		return NewTimeFromTimestamp(kind.Data)
	default:
		return nil
	}
}

/* __________________________________________________ */

// NilOptionalTimestamp ...
func NilOptionalTimestamp() *common.OptionalTimestamp {
	return &common.OptionalTimestamp{Kind: nil}
}

// NewOptionalTimestampFromTime ...
func NewOptionalTimestampFromTime(value *time.Time) *common.OptionalTimestamp {
	if value == nil {
		return NilOptionalTimestamp()
	}
	return &common.OptionalTimestamp{
		Kind: &common.OptionalTimestamp_Data{
			Data: NewTimestampFromTime(value),
		},
	}
}

/* __________________________________________________ */
