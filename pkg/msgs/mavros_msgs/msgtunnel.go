//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
)

const (
	Tunnel_PAYLOAD_TYPE_UNKNOWN           uint16 = 0
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED0 uint16 = 200
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED1 uint16 = 201
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED2 uint16 = 202
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED3 uint16 = 203
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED4 uint16 = 204
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED5 uint16 = 205
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED6 uint16 = 206
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED7 uint16 = 207
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED8 uint16 = 208
	Tunnel_PAYLOAD_TYPE_STORM32_RESERVED9 uint16 = 209
)

type Tunnel struct {
	msg.Package     `ros:"mavros_msgs"`
	msg.Definitions `ros:"uint16 PAYLOAD_TYPE_UNKNOWN=0,uint16 PAYLOAD_TYPE_STORM32_RESERVED0=200,uint16 PAYLOAD_TYPE_STORM32_RESERVED1=201,uint16 PAYLOAD_TYPE_STORM32_RESERVED2=202,uint16 PAYLOAD_TYPE_STORM32_RESERVED3=203,uint16 PAYLOAD_TYPE_STORM32_RESERVED4=204,uint16 PAYLOAD_TYPE_STORM32_RESERVED5=205,uint16 PAYLOAD_TYPE_STORM32_RESERVED6=206,uint16 PAYLOAD_TYPE_STORM32_RESERVED7=207,uint16 PAYLOAD_TYPE_STORM32_RESERVED8=208,uint16 PAYLOAD_TYPE_STORM32_RESERVED9=209"`
	TargetSystem    uint8
	TargetComponent uint8
	PayloadType     uint16
	PayloadLength   uint8
	Payload         [128]uint8
}
