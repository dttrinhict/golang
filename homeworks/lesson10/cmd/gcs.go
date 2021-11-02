package main

import (
	"fmt"
	"golang/homeworks/lesson10/util/storages"
)

func main() {
	// Sets your Google Cloud Platform project ID.
	projectID := "vinid-playground"
	gcs, err := storages.GCSInit("/Users/trinhdt2/learn/golang-techmaster/golang/homeworks/lesson10/util/storages/vinid-playground-5afeaf8166fa.json", projectID)
	if err != nil {
		fmt.Println("Create GCS instance is failed")
	}
	fmt.Println(gcs.Client)
	bucketName := "trinhdt2-test"
	result, err := gcs.CreateBucket(bucketName)
	if err != nil {
		fmt.Println("Error while create bucket ", bucketName)
	}else{
		fmt.Printf("Bucket: %v already created.\n", *result)
	}

	fileName := "test.txt"
	errWriter := gcs.WriteFileToBucket(fileName, bucketName)
	if errWriter != nil {
		fmt.Println("Error while write to bucket "+ errWriter.Error())
	}
	content, errReader := gcs.ShowFileFromBucket(fileName, bucketName)
	if errReader != nil {
		fmt.Println("Error while read file from bucket "+ errWriter.Error())
	}
	fmt.Println("Content file: ", *content)


	fmt.Println("Download file")
	object := "file_example_MP4_1920_18MG.mp4"
	gcs.DownloadFile("/Users/trinhdt2/learn/golang-techmaster/golang/homeworks/lesson10/tmp/file_example_MP4_1920_18MG.mp4",object, bucketName)
}
