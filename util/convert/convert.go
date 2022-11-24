package convert

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// StringToProtoTimestamp converts the string to google.protobuf.Timestamp.
func StringToProtoTimestamp(layout, value string) (*timestamp.Timestamp, error) {
	tt, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}
	return ptypes.TimestampProto(tt)
}

// ProtoTimestampToString converts the google.protobuf.Timestamp to string.
func ProtoTimestampToString(layout string, value *timestamp.Timestamp) (string, error) {
	tt, err := ptypes.Timestamp(value)
	if err != nil {
		return "", err
	}
	return tt.Format(layout), nil
}
