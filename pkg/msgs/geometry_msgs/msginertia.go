//autogenerated:yes
//nolint:revive,lll
package geometry_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
)

type Inertia struct {
	msg.Package `ros:"geometry_msgs"`
	M           float64
	Com         Vector3
	Ixx         float64
	Ixy         float64
	Ixz         float64
	Iyy         float64
	Iyz         float64
	Izz         float64
}
