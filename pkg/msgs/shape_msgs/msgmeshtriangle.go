//autogenerated:yes
//nolint:revive,lll
package shape_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
)

type MeshTriangle struct {
	msg.Package   `ros:"shape_msgs"`
	VertexIndices [3]uint32
}
