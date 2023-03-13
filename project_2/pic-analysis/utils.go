package pic_analysis

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

func UpdateMetadata(bucket, object string, metadata map[string]string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	_, err = client.Bucket(bucket).Object(object).Update(ctx, storage.ObjectAttrsToUpdate{
		Metadata: metadata,
	})

	if err != nil {
		return err
	}

	return nil
}

// getMetadata prints and returns object attributes.
func getMetadata(bucket, object string) (map[string]string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(bucket).Object(object)
	attrs, err := o.Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).Attrs: %v", object, err)
	}

	return attrs.Metadata, nil
}
