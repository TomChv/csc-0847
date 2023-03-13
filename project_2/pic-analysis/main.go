// Get from https://github.com/GoogleCloudPlatform/serverless-photosharing-workshop/blob/master/functions/image-analysis/go/response.go
package pic_analysis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	vision "cloud.google.com/go/vision/apiv1"

	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
)

// GCSEvent is the payload of a GCS event. Please refer to the docs for
// additional information regarding GCS events.
type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// firestorePicture is the structure of a `pictures` document in Firestore
type firestorePicture struct {
	Labels []string `firestore:"labels"`
	Color  string   `firestore:"color"`
	// Created value will be set with the time of creation on the server (firestore) side
	// see https://godoc.org/cloud.google.com/go/firestore#DocumentRef.Create
	Created time.Time `firestore:"created,serverTimestamp"`
}

func VisionAnalysis(ctx context.Context, e GCSEvent) error {
	log.Printf("Event: %#v", e)

	filename := e.Name
	filebucket := e.Bucket
	log.Printf("New picture uploaded %s in %s", filename, filebucket)

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return errors.New("failed to create CloudVision client")
	}
	defer client.Close()

	request := &pb.AnnotateImageRequest{
		Image: &pb.Image{
			Source: &pb.ImageSource{
				ImageUri: fmt.Sprintf("gs://%s/%s", filebucket, filename),
			},
		},
		Features: []*pb.Feature{
			{Type: pb.Feature_LABEL_DETECTION},
			{Type: pb.Feature_IMAGE_PROPERTIES},
			{Type: pb.Feature_SAFE_SEARCH_DETECTION},
		},
	}

	r, err := client.AnnotateImage(ctx, request)
	if err != nil {
		log.Printf("Failed annotate image: %v", err)
		return fmt.Errorf("vision API error: code %d, message: '%s'", r.Error.Code, r.Error.Message)
	}

	resp := visionResponse{r}
	log.Printf("Raw vision output for: %s: %s", filename, resp.toJSON())

	labels := resp.getLabels()
	log.Printf("Labels: %s", strings.Join(labels, ", "))

	if !resp.isSafe() {
		return nil
	}

	metadata, err := getMetadata(filebucket, filename)
	if err != nil {
		log.Printf("Failed to get image's metadata: %v", err)
		return errors.New("failed to get image's metadata")
	}

	metadata["Label"] = strings.Join(labels, ",")
	if err := UpdateMetadata(filebucket, filename, metadata); err != nil {
		log.Printf("failed to add labels in metadata: %v", err)
		return err
	}
	return nil
}
