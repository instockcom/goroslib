//autogenerated:yes
//nolint:revive,lll
package std_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
)

type Int8MultiArray struct {
	msg.Package `ros:"std_msgs"`
	Layout      MultiArrayLayout
	Data        []int8
}
