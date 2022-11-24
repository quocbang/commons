package metadata

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

// Meta is used to define the metadata of the client.
type Meta struct {
	Token       string
	UserID      string
	ProgramType string
	ProgramID   string
	PID         string
	MaxRecords  int
	Timezone    string
}

// ContextWith returns context with specific metadata.
func ContextWith(ctx context.Context, md *Meta) context.Context {
	pairs := make([]string, 0, 14)
	if md.Token != "" {
		pairs = append(pairs, Token, md.Token)
	}
	if md.UserID != "" {
		pairs = append(pairs, UserID, md.UserID)
	}
	if md.ProgramType != "" {
		pairs = append(pairs, ProgramType, md.ProgramType)
	}
	if md.ProgramID != "" {
		pairs = append(pairs, ProgramID, md.ProgramID)
	}
	if md.PID != "" {
		pairs = append(pairs, PID, md.PID)
	}
	if md.MaxRecords > 0 {
		pairs = append(pairs, MaxRecords, strconv.Itoa(md.MaxRecords))
	}
	if md.Timezone != "" {
		pairs = append(pairs, Timezone, md.Timezone)
	}
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(pairs...))
}
