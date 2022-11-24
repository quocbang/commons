package metadata

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"gitlab.kenda.com.tw/kenda/commons/v2/util/math"
)

const grpcGatewayForwardIP = "x-forwarded-for"

// Metadata is the metadata
type Metadata interface {
	Token() string
	UserID() string
	HostIP() string
	HostAddress() string
	ProgramType() string
	ProgramID() string
	ProgramTrackingNumber() string
	MaxRecords() int
	Timezone() string
}

// MD is the metadata for server.
type MD struct {
	token         string
	userID        string
	clientIP      string
	clientAddress string
	programType   string
	programID     string
	pid           string
	maxRecords    int
	timezone      string
}

// Token is the token.
func (md *MD) Token() string {
	return md.token
}

// UserID is user ID.
func (md *MD) UserID() string {
	return md.userID
}

// InsteadUserID is to be instead of the user id.
func (md *MD) InsteadUserID(id string) {
	md.userID = id
}

// HostIP is the IP from the client.
func (md *MD) HostIP() string {
	return md.clientIP
}

// HostAddress is the address from the client.
func (md *MD) HostAddress() string {
	return md.clientAddress
}

// ProgramType is program type.
func (md *MD) ProgramType() string {
	return md.programType
}

// ProgramID is program id.
func (md *MD) ProgramID() string {
	return md.programID
}

// ProgramTrackingNumber is program tracking number.
func (md *MD) ProgramTrackingNumber() string {
	return md.pid
}

// MaxRecords is the maximum number of query records.
func (md *MD) MaxRecords() int {
	return md.maxRecords
}

// InsteadDefaultMaxRecords is to be instead of the maximum number of query records.
func (md *MD) InsteadDefaultMaxRecords() {
	md.maxRecords = math.MaxInt
}

// Timezone is the timezone from the client.
func (md *MD) Timezone() string {
	return md.timezone
}

// FromGrpcContext returns metadata
func FromGrpcContext(ctx context.Context) (*MD, error) {
	md := &MD{}

	p, ok := peer.FromContext(ctx)
	if ok {
		md.clientIP = p.Addr.String()
	}

	tMd, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata: failed to get metadata from incoming context")
	}

	if tToken := tMd.Get(Token); len(tToken) > 0 {
		md.token = tToken[0]
	}
	if tUserID := tMd.Get(UserID); len(tUserID) > 0 {
		md.userID = tUserID[0]
	}
	if tPrgType := tMd.Get(ProgramType); len(tPrgType) > 0 {
		md.programType = tPrgType[0]
	}
	if tPrgNo := tMd.Get(ProgramID); len(tPrgNo) > 0 {
		md.programID = tPrgNo[0]
	}
	if tGatewayForWard := tMd.Get(grpcGatewayForwardIP); len(tGatewayForWard) > 0 {
		md.clientIP = tGatewayForWard[0]
		md.clientAddress = ""
	}
	if tPID := tMd.Get(PID); len(tPID) > 0 {
		md.pid = tPID[0]
	}
	if tMaxRecords := tMd.Get(MaxRecords); len(tMaxRecords) > 0 {
		amountMaxRecords, err := strconv.Atoi(tMaxRecords[0])
		if err != nil {
			return nil, fmt.Errorf("metadata: invalid %s, got '%s'", MaxRecords, tMaxRecords[0])
		}
		md.maxRecords = amountMaxRecords
	} else {
		md.InsteadDefaultMaxRecords()
	}

	return md, nil
}
