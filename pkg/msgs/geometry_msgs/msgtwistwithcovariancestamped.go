//autogenerated:yes
//nolint:revive,lll
package geometry_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type TwistWithCovarianceStamped struct {
	msg.Package `ros:"geometry_msgs"`
	Header      std_msgs.Header
	Twist       TwistWithCovariance
}
