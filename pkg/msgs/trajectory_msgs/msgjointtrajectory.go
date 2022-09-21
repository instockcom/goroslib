//autogenerated:yes
//nolint:revive,lll
package trajectory_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type JointTrajectory struct {
	msg.Package `ros:"trajectory_msgs"`
	Header      std_msgs.Header
	JointNames  []string
	Points      []JointTrajectoryPoint
}
