//autogenerated:yes
//nolint:revive,lll
package uuid_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
)

type UniqueID struct {
	msg.Package `ros:"uuid_msgs"`
	Uuid        [16]uint8
}
