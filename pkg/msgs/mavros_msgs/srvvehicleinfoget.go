//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
)

const (
	VehicleInfoGetReq_GET_MY_SYSID  uint8 = 0
	VehicleInfoGetReq_GET_MY_COMPID uint8 = 0
)

type VehicleInfoGetReq struct {
	msg.Package     `ros:"mavros_msgs"`
	msg.Definitions `ros:"uint8 GET_MY_SYSID=0,uint8 GET_MY_COMPID=0"`
	Sysid           uint8
	Compid          uint8
	GetAll          bool
}

type VehicleInfoGetRes struct {
	msg.Package `ros:"mavros_msgs"`
	Success     bool
	Vehicles    []VehicleInfo
}

type VehicleInfoGet struct {
	msg.Package `ros:"mavros_msgs"`
	VehicleInfoGetReq
	VehicleInfoGetRes
}
