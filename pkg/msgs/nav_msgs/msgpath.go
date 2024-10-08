//autogenerated:yes
//nolint:revive,lll
package nav_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/geometry_msgs"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type Path struct {
	msg.Package `ros:"nav_msgs"`
	Header      std_msgs.Header
	Poses       []geometry_msgs.PoseStamped
}
