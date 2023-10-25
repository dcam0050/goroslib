//autogenerated:yes
//nolint:revive,lll
package vision_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/geometry_msgs"
)

type ObjectHypothesisWithPose struct {
	msg.Package `ros:"vision_msgs"`
	Id          int64
	Score       float64
	Pose        geometry_msgs.PoseWithCovariance
}