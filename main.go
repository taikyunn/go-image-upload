package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1")},
	)
	if err != nil {
		log.Fatal(err)
	}

	// S3クライアントを作成
	svc := s3.New(sess)

	// バケットの一覧を取得：ListBuckets
	result, _ := svc.ListBuckets(nil)
	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	//バケットの新規作成
	// bucket := "new-bucket-nam-taichi-1"
	// resp, err := svc.CreateBucket(&s3.CreateBucketInput{
	// 	Bucket: aws.String(bucket),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)

	// log.Println("処理開始。")
	// // sessionの作成
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	Profile:           "di",
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	// バケットへアップロード
	// ファイルを開く
	// targetFilePath := "./upload.jpeg"
	// file, err := os.Open(targetFilePath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// bucketName := "new-bucket-nam-taichi-1"
	// objectKey := "test-key-1"

	// // Uploaderを作成し、ローカルファイルをアップロード
	// uploader := s3manager.NewUploader(sess)
	// _, err = uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(objectKey),
	// 	Body:   file,
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("done")

	// S3からのダウンロード
	// S3オブジェクトを書き込むファイルの作成
	f, err := os.Create("download.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	bucketName := "new-bucket-nam-taichi-1"
	objectKey := "test-key-1"

	// Downloaderを作成し、S3オブジェクトをダウンロード
	downloader := s3manager.NewDownloader(sess)
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DownloadedSize: %d byte", n)

}
