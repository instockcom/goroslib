//autogenerated:yes
//nolint:revive,lll
package rosgraph_msgs

import (
	"time"

	"github.com/instockcom/goroslib/pkg/msg"
)

type Clock struct {
	msg.Package `ros:"rosgraph_msgs"`
	Clock       time.Time
}
