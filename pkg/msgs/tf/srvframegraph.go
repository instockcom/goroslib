//autogenerated:yes
//nolint:revive,lll
package tf

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type FrameGraphReq struct {
	msg.Package `ros:"tf"`
}

type FrameGraphRes struct {
	msg.Package `ros:"tf"`
	DotGraph    string
}

type FrameGraph struct {
	msg.Package `ros:"tf"`
	FrameGraphReq
	FrameGraphRes
}
