package main

import (
	"fmt"
	"golang/homeworks/lesson12/util/storages"
)

func main() {
	// Sets your Google Cloud Platform project ID.
	//projectID := "vinid-playground"
	projectID := "onemg-general-np"
	//gcs, err := storages.GCSInit("/Users/trinhdt2/learn/golang-techmaster/golang/homeworks/lesson10/util/storages/vinid-playground-5afeaf8166fa.json", projectID)
	gcs, err := storages.GCSInit("", projectID)
	if err != nil {
		fmt.Println("Create GCS instance is failed")
	}
	fmt.Println(gcs.Client)
	bucketName := "om-chatbot-fpt-np"
	//result, err := gcs.CreateBucket(bucketName)
	//if err != nil {
	//	fmt.Println("Error while create bucket ", bucketName)
	//}else{
	//	fmt.Printf("Bucket: %v already created.\n", *result)
	//}
	//
	//fileName := "test.txt"
	//errWriter := gcs.WriteFileToBucket(fileName, bucketName)
	//if errWriter != nil {
	//	fmt.Println("Error while write to bucket "+ errWriter.Error())
	//}
	//content, errReader := gcs.ShowFileFromBucket(fileName, bucketName)
	//if errReader != nil {
	//	fmt.Println("Error while read file from bucket "+ errWriter.Error())
	//}
	//fmt.Println("Content file: ", *content)


	fmt.Println("Download file")
	object := "voice/2021/07/26/recording/061598ae-0276-42a5-ab6d-723fbecfcf6c.wav.gpg"
	gcs.DownloadFile("/Users/trinhdt2/learn/golang-techmaster/golang/homeworks/lesson10/tmp/061598ae-0276-42a5-ab6d-723fbecfcf6c.wav.gpg",object, bucketName)
}
