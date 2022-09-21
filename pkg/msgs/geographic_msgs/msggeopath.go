//autogenerated:yes
//nolint:revive,lll
package geographic_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type GeoPath struct {
	msg.Package `ros:"geographic_msgs"`
	Header      std_msgs.Header
	Poses       []GeoPoseStamped
}
