//autogenerated:yes
//nolint:revive,lll
package sensor_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type PointCloud2 struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Height      uint32
	Width       uint32
	Fields      []PointField
	IsBigendian bool
	PointStep   uint32
	RowStep     uint32
	Data        []uint8
	IsDense     bool
}
