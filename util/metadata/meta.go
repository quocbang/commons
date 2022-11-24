package metadata

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

// metadata header
const (
	Token       = "Token"
	UserID      = "User-Id"
	ProgramType = "Prg-Type"
	ProgramID   = "Prg-No"
	PID         = "Pid"
	MaxRecords  = "Max-Records"
)

// MetaContext defines the functions needed for composing metadata
type MetaContext interface {
	UserToken() string
	UserID() string
	ProgramType() string
	ProgramID() string
	ProgramTrackingNumber() string
}

// ContextWith returns context with specific metadata
func ContextWith(ctx MetaContext) context.Context {
	md := metadata.Pairs(
		Token, ctx.UserToken(),
		UserID, ctx.UserID(),
		ProgramType, ctx.ProgramType(),
		ProgramID, ctx.ProgramID(),
		PID, ctx.ProgramTrackingNumber(),
	)

	return metadata.NewOutgoingContext(context.Background(), md)
}

// ContextWithMaxRecords returns context with the maximum amount of the records from DM
func ContextWithMaxRecords(ctx context.Context, max int) context.Context {
	return metadata.AppendToOutgoingContext(ctx, MaxRecords, strconv.Itoa(max))
}

// OutgoingMD defines an MD for outgoing
type OutgoingMD struct {
	userToken   string
	userID      string
	programType string
	programID   string
	pid         string
}

// UserToken 使用者token
func (md *OutgoingMD) UserToken() string {
	return md.userToken
}

// UserID 使用者帳號
func (md *OutgoingMD) UserID() string {
	return md.userID
}

// ProgramType 程式類別
func (md *OutgoingMD) ProgramType() string {
	return md.programType
}

// ProgramID 程式代號
func (md *OutgoingMD) ProgramID() string {
	return md.programID
}

// ProgramTrackingNumber process id
func (md *OutgoingMD) ProgramTrackingNumber() string {
	return md.pid
}

// NewOutgoingMD creates an MD for outgoing
func NewOutgoingMD(userToken, userID, programType, programID, pid string) *OutgoingMD {
	return &OutgoingMD{
		userToken:   userToken,
		userID:      userID,
		programType: programType,
		programID:   programID,
		pid:         pid,
	}
}
