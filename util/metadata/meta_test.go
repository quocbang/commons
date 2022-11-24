package metadata

import (
	"context"
	"strconv"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestContextWithMaxRecords(t *testing.T) {
	myMaxRecords := 10

	emptyCtx := context.Background()
	emptyCtx = ContextWithMaxRecords(emptyCtx, myMaxRecords)
	maxRecordsInContext := getMaxRecordsFromContext(emptyCtx)
	if maxRecordsInContext != myMaxRecords {
		t.Fatal("got a wrong the maximum amound of records. expected:", myMaxRecords, "got:", maxRecordsInContext)
	}

	ctx := ContextWith(NewOutgoingMD("userToken", "userID", "programType", "programID", "pid"))
	ctx = ContextWithMaxRecords(ctx, myMaxRecords)
	maxRecordsInContext = getMaxRecordsFromContext(ctx)
	if maxRecordsInContext != myMaxRecords {
		t.Fatal("got a wrong the maximum amound of records. expected:", myMaxRecords, "got:", maxRecordsInContext)
	}
}

func getMaxRecordsFromContext(ctx context.Context) int {
	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		if myMaxRecord := md.Get(MaxRecords); len(myMaxRecord) > 0 {
			num, err := strconv.Atoi(myMaxRecord[0])
			if err != nil {
				panic(err)
			}
			return num
		}
	}

	return 0
}
