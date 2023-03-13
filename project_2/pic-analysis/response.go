// Get from https://github.com/GoogleCloudPlatform/serverless-photosharing-workshop/blob/master/functions/image-analysis/go/response.go
package pic_analysis

import (
	"encoding/json"
	"fmt"
	"sort"

	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
)

type visionResponse struct {
	*pb.AnnotateImageResponse
}

// toJSON returns a JSON representation of a response
func (o *visionResponse) toJSON() string {
	b, err := json.Marshal(o)
	if err != nil {
		return "## error marshalling data ##"
	}
	return string(b)
}

// byScore implements sort.Interface based on the Score field
type byScore []*pb.EntityAnnotation

func (o byScore) Len() int           { return len(o) }
func (o byScore) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o byScore) Less(i, j int) bool { return o[i].Score > o[j].Score }

// getLabels returns the labels found in the response ordered by descending score
func (o *visionResponse) getLabels() (labels []string) {
	sort.Sort(byScore(o.LabelAnnotations))
	for _, label := range o.LabelAnnotations {
		labels = append(labels, label.Description)
	}
	return
}

// getDominantColor returns a Hex representation of the dominant color in the image
func (o *visionResponse) getDominantColor() (hex string) {
	var bestScore float32
	var bestColor *pb.ColorInfo
	for _, color := range o.ImagePropertiesAnnotation.DominantColors.Colors {
		if color.Score > bestScore {
			bestScore = color.Score
			bestColor = color
		}
	}
	if bestColor == nil {
		return "#ffffff"
	}
	return fmt.Sprintf("#%02x%02x%02x", int(bestColor.Color.Red), int(bestColor.Color.Green), int(bestColor.Color.Blue))
}

// isSafe returns true if no field of SafeSearchAnnotation is LIKELY or more
func (o *visionResponse) isSafe() bool {
	safe := o.SafeSearchAnnotation
	for _, value := range []*pb.Likelihood{&safe.Adult, &safe.Medical, &safe.Racy, &safe.Spoof, &safe.Violence} {
		if *value == pb.Likelihood_LIKELY || *value == pb.Likelihood_VERY_LIKELY {
			return false
		}
	}
	return true
}
