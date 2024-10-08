//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type ESCInfo struct {
	msg.Package    `ros:"mavros_msgs"`
	Header         std_msgs.Header
	Counter        uint16
	Count          uint8
	ConnectionType uint8
	Info           uint8
	EscInfo        []ESCInfoItem
}
