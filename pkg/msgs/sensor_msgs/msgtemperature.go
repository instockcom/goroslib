//autogenerated:yes
//nolint:revive,lll
package sensor_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type Temperature struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Temperature float64
	Variance    float64
}
