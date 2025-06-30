package main

import (
    "bytes"
    "fmt"
    "mime/multipart"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/joho/godotenv"
)

var s3Session *s3.S3
var bucketName string

func InitB2() error {
    _ = godotenv.Load()

    keyID := os.Getenv("B2_KEY_ID")
    appKey := os.Getenv("B2_APP_KEY")
    region := os.Getenv("B2_REGION")
    endpoint := os.Getenv("B2_ENDPOINT")
    bucketName = os.Getenv("B2_BUCKET")

    sess, err := session.NewSession(&aws.Config{
        Region:           aws.String(region),
        Endpoint:         aws.String(endpoint),
        S3ForcePathStyle: aws.Bool(true),
        Credentials:      credentials.NewStaticCredentials(keyID, appKey, ""),
    })
    if err != nil {
        return err
    }

    s3Session = s3.New(sess)
    return nil
}

func UploadToB2(file multipart.File, filename string, contentType string) (string, error) {
    buf := new(bytes.Buffer)
    _, err := buf.ReadFrom(file)
    if err != nil {
        return "", err
    }

    _, err = s3Session.PutObject(&s3.PutObjectInput{
        Bucket:      aws.String(bucketName),
        Key:         aws.String(filename),
        Body:        bytes.NewReader(buf.Bytes()),
        ContentType: aws.String(contentType),
        ACL:         aws.String("public-read"),
    })
    if err != nil {
        return "", err
    }

    endpoint := os.Getenv("B2_ENDPOINT")
    url := fmt.Sprintf("https://%s.%s/%s", bucketName,endpoint, filename)
    return url, nil
}

func DeleteFromB2(filename string) error {
    fmt.Println("刪除檔案請求：", bucketName, filename)

    _, err := s3Session.DeleteObject(&s3.DeleteObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(filename),
    })
    return err
}