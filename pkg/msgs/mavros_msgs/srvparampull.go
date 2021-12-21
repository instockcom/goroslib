//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type ParamPullReq struct {
	msg.Package `ros:"mavros_msgs"`
	ForcePull   bool
}

type ParamPullRes struct {
	msg.Package   `ros:"mavros_msgs"`
	Success       bool
	ParamReceived uint32
}

type ParamPull struct {
	msg.Package `ros:"mavros_msgs"`
	ParamPullReq
	ParamPullRes
}
