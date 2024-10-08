//autogenerated:yes
//nolint:revive,lll
package vision_msgs

import (
	"github.com/instockcom/goroslib/pkg/msg"
	"github.com/instockcom/goroslib/pkg/msgs/sensor_msgs"
	"github.com/instockcom/goroslib/pkg/msgs/std_msgs"
)

type Classification2D struct {
	msg.Package `ros:"vision_msgs"`
	Header      std_msgs.Header
	Results     []ObjectHypothesis
	SourceImg   sensor_msgs.Image
}
