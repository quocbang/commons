package metadata

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/metadata"

	dmRS "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/rs"
	commonsMD "gitlab.kenda.com.tw/kenda/commons/v2/util/metadata"
)

// metadata key
const (
	Station           = "station"
	Operator          = commonsMD.UserID
	Site              = "site"
	TimeZone          = "timezone"
	Token             = commonsMD.Token
	ProgramType       = commonsMD.ProgramType
	ProgramID         = commonsMD.ProgramID
	RemoteHost        = "remoteHost"
	PID               = commonsMD.PID
	ResourceFactoryID = "factoryID" // for MES v2 ResourceBind API, this is short-term usage.
)

// Metadata for mes api.
type Metadata struct {
	Station           string
	Operator          string
	Site              string
	TimeZone          string // for safety reasons, we recommend to use SetTimeZone method instead of assigning value directly.
	Token             string
	ProgramType       string
	ProgramID         string
	RemoteHost        string
	PID               string
	ResourceFactoryID dmRS.FactoryID
}

// ToOutgoingContext returns outgoing context for mes gRPC api.
func (m *Metadata) ToOutgoingContext(ctx context.Context) context.Context {
	md := metadata.Pairs(
		Station, m.Station,
		Operator, m.Operator,
		Site, m.Site,
		TimeZone, m.TimeZone,
		Token, m.Token,
		ProgramType, m.ProgramType,
		ProgramID, m.ProgramID,
		RemoteHost, m.RemoteHost,
		PID, m.PID,
		ResourceFactoryID, m.ResourceFactoryID.String(),
	)
	return metadata.NewOutgoingContext(ctx, md)
}

// SetTimeZone sets specified time location name to metadata.
func (m *Metadata) SetTimeZone(name string) error {
	if len(name) == 0 {
		return errors.New("missing name")
	}
	if _, err := time.LoadLocation(name); err != nil {
		return err
	}
	m.TimeZone = name
	return nil
}
