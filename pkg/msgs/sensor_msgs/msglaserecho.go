//autogenerated:yes
//nolint:revive,lll
package sensor_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
)

type LaserEcho struct {
	msg.Package `ros:"sensor_msgs"`
	Echoes      []float32
}
