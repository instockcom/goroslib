//autogenerated:yes
//nolint:revive,lll
package vision_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
)

type ObjectHypothesis struct {
	msg.Package `ros:"vision_msgs"`
	Id          int64
	Score       float64
}
