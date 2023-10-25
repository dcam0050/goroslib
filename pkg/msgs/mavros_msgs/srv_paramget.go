//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type ParamGetReq struct {
	msg.Package `ros:"mavros_msgs"`
	ParamId     string
}

type ParamGetRes struct {
	msg.Package `ros:"mavros_msgs"`
	Success     bool
	Value       ParamValue
}

type ParamGet struct {
	msg.Package `ros:"mavros_msgs"`
	ParamGetReq
	ParamGetRes
}