package utils

import (
	"service/common/define/constant"
	"service/presentation/grpc/pb"

	"github.com/grpc-boot/gateway"
)

// ConvertGatewayInfo 转换gateway信息
func ConvertGatewayInfo(info *gateway.Info) *pb.GwReply {
	data := &pb.GwReply{
		Code: constant.Success,
		Msg:  "ok",
		Data: &pb.GwInfo{
			Qps:        info.Qps,
			Total:      info.Total,
			MethodList: make([]*pb.MethodInfo, 0, len(info.MethodList)),
		},
	}

	for _, mi := range info.MethodList {
		mInfo := &pb.MethodInfo{
			Name:        mi.Name,
			Path:        mi.Path,
			SecondLimit: int32(mi.SecondLimit),
			Qps:         mi.Qps,
			Total:       mi.Total,
			Avg:         mi.Avg,
			Min:         mi.Min,
			Max:         mi.Max,
			Line_90:     mi.L90,
			Line_95:     mi.L95,
			CodeMap:     make(map[int32]uint64, len(mi.CodeMap)),
		}

		for code, count := range mi.CodeMap {
			mInfo.CodeMap[int32(code)] = count
		}

		data.Data.MethodList = append(data.Data.MethodList, mInfo)
	}

	return data
}
