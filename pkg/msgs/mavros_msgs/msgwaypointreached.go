//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type WaypointReached struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	WpSeq       uint16
}
