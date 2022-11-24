package metadata

import (
	"context"
	"strconv"
	"testing"

	"google.golang.org/grpc/metadata"

	"gitlab.kenda.com.tw/kenda/commons/v2/util/math"
)

func TestFromGrpcContext(t *testing.T) {
	myMaxRecords := 100

	ctx := fakeContext(strconv.Itoa(myMaxRecords))
	md, err := FromGrpcContext(ctx)
	if err != nil {
		t.Fatal("occur error:", err)
	}
	if md.MaxRecords() != myMaxRecords {
		t.Fatal("wrong MaxRecords, got:", md.MaxRecords(), "expected:", myMaxRecords)
	}

	ctx = fakeContextWithoutMaxRecords()
	md, err = FromGrpcContext(ctx)
	if err != nil {
		t.Fatal("occur error:", err)
	}
	maxInt := math.MaxInt
	if md.MaxRecords() != maxInt {
		t.Fatal("wrong MaxRecords, got:", md.MaxRecords(), "expected:", maxInt)
	}
}

func fakeContextWithoutMaxRecords() context.Context {
	return metadata.NewIncomingContext(
		context.Background(),
		metadata.Pairs(
			Token, "Token",
			UserID, "UserID",
			ProgramType, "ProgramType",
			ProgramID, "ProgramID",
			PID, "PID",
		),
	)
}

func fakeContext(max string) context.Context {
	return metadata.NewIncomingContext(
		context.Background(),
		metadata.Pairs(
			Token, "Token",
			UserID, "UserID",
			ProgramType, "ProgramType",
			ProgramID, "ProgramID",
			PID, "PID",
			MaxRecords, max,
		),
	)
}
