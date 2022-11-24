package metadata

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"gitlab.kenda.com.tw/kenda/commons/v2/util/math"
)

const grpcGatewayForwardIP = "x-forwarded-for"

//Metadata metadata.
type Metadata interface {
	MetaContext

	HostIP() string
	HostName() string

	MaxRecords() int
}

// MD metadata.
type MD struct {
	userToken  string
	userID     string
	hostIP     string
	hostName   string
	prgType    string
	prgNo      string
	pid        string
	maxRecords int
}

// UserToken 使用者token
func (md *MD) UserToken() string {
	return md.userToken
}

// UserID 使用者帳號
func (md *MD) UserID() string {
	return md.userID
}

// ProgramType 程式類別
func (md *MD) ProgramType() string {
	return md.prgType
}

// ProgramID 程式代號
func (md *MD) ProgramID() string {
	return md.prgNo
}

// ProgramTrackingNumber process id
func (md *MD) ProgramTrackingNumber() string {
	return md.pid
}

// HostIP 呼叫端ip
func (md *MD) HostIP() string {
	return md.hostIP
}

// HostName 呼叫端主機名稱
func (md *MD) HostName() string {
	return md.hostName
}

// MaxRecords 最大回傳資料量
func (md *MD) MaxRecords() int {
	return md.maxRecords
}

// String returns string
func (md *MD) String() string {
	return fmt.Sprintf(
		"userToken:%s userID:%s programType:%s programID:%s pid:%s ip:%s hostName:%s maxRecords:%d",
		md.UserToken(),
		md.UserID(),
		md.ProgramType(),
		md.ProgramID(),
		md.ProgramTrackingNumber(),
		md.HostIP(),
		md.HostName(),
		md.MaxRecords(),
	)
}

// FromGrpcContext returns metadata
func FromGrpcContext(ctx context.Context) (Metadata, error) {
	md := &MD{}

	p, ok := peer.FromContext(ctx)
	if ok {
		md.hostIP = p.Addr.String()
		if p.AuthInfo != nil {
			switch tlsInfo := p.AuthInfo.(type) {
			case credentials.TLSInfo:
				md.hostName = tlsInfo.State.ServerName
			}
		}
	}

	tMd, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata error:failed to get metadata from incoming context")
	}

	if tToken := tMd.Get(Token); len(tToken) > 0 {
		md.userToken = tToken[0]
	}
	if tUserID := tMd.Get(UserID); len(tUserID) > 0 {
		md.userID = tUserID[0]
	}
	if tPrgType := tMd.Get(ProgramType); len(tPrgType) > 0 {
		md.prgType = tPrgType[0]
	}
	if tPrgNo := tMd.Get(ProgramID); len(tPrgNo) > 0 {
		md.prgNo = tPrgNo[0]
	}
	if tGatewayForWard := tMd.Get(grpcGatewayForwardIP); len(tGatewayForWard) > 0 {
		md.hostIP = tGatewayForWard[0]
		md.hostName = ""
	}
	if tPID := tMd.Get(PID); len(tPID) > 0 {
		md.pid = tPID[0]
	}
	if tMaxRecords := tMd.Get(MaxRecords); len(tMaxRecords) > 0 {
		amountMaxRecords, err := strconv.Atoi(tMaxRecords[0])
		if err != nil {
			return nil, errors.New("metadata error:failed to convert string to int")
		}
		md.maxRecords = amountMaxRecords
	} else {
		md.maxRecords = math.MaxInt
	}

	return md, nil
}

// New returns Metadata
func New(userToken, userID, hostIP, hostName, prgType, prgNo, pid string, maxRecords int) Metadata {
	return &MD{
		userToken:  userToken,
		userID:     userID,
		hostIP:     hostIP,
		hostName:   hostName,
		prgType:    prgType,
		prgNo:      prgNo,
		pid:        pid,
		maxRecords: maxRecords,
	}
}
