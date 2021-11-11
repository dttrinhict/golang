package storages

import (
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type GCS struct {
	Client *storage.Client
	ProjectId string
	Ctx context.Context
}

var Gcs *GCS

func GCSInit(credentials, projectID string) (*GCS, error) {
	if Gcs == nil {
		ctx := context.Background()
		//client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentials))
		client, err := storage.NewClient(ctx)
		//client, err := storage.NewClient(ctx, option.ImpersonateCredentials())
		if err != nil {
			return nil, err
		}
		Gcs = &GCS{
			Client: client,
			ProjectId: projectID,
			Ctx: ctx,
		}
		return Gcs, nil
	}
	return Gcs, nil
}

func (gcs *GCS) CreateBucket(name string) (*string, error) {
	// Creates a Bucket instance.
	bucket := gcs.Client.Bucket(name)
	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(gcs.Ctx, time.Second*15)
	defer cancel()
	if err := bucket.Create(ctx, gcs.ProjectId, &storage.BucketAttrs{
		Location: "Asia",
	}); err != nil {
		return nil, err
	}
	return &name, nil
}

func (gcs *GCS) WriteFileToBucket(filename, bucketname string) error {
	bucket := gcs.Client.Bucket(bucketname)
	writer := bucket.Object(filename).NewWriter(gcs.Ctx)
	writer.ContentType = "text/plain"
	if _, err := writer.Write([]byte("hello world")); err != nil {
		return err
		// TODO: handle error.
		// Note that Write may return nil in some error situations,
		// so always check the error from Close.
	}
	if err := writer.Close(); err != nil {
		// TODO: handle error.
		return err
	}
	fmt.Println("updated object:", writer.Attrs())
	return nil
}

func (gcs *GCS) ReadFileFromBucket(filename, bucketname string) ([]byte, error) {
	bucket := gcs.Client.Bucket(bucketname)
	reader, err := bucket.Object(filename).NewReader(gcs.Ctx)
	defer  reader.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}


func (gcs *GCS) ShowFileFromBucket(filename, bucketname string) (*string, error) {
	contentByte, err := gcs.ReadFileFromBucket(filename, bucketname)
	if err != nil {
		// TODO: handle error.
		return nil, err
	}
	content := string(contentByte)
	return &content, nil
}

// downloadFile downloads an object.
func (gcs *GCS) DownloadFile(output, filename, bucketname string) {
	data, err := gcs.ReadFileFromBucket(filename,bucketname)
	if err != nil {
		fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	err = os.WriteFile(output, data, 0644)
	if err != nil {
		fmt.Errorf("ioutil.WriteFile: %v", err)
	}
}
