package services

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/ronanvirmani/event-management-system/backend/config"
)

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
    bucket := os.Getenv("AWS_S3_BUCKET")
    region := os.Getenv("AWS_REGION")

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region)},
    )
    if err != nil {
        return "", err
    }

    uploader := s3.New(sess)

    size := fileHeader.Size
    buffer := make([]byte, size)
    file.Read(buffer)

    fileName := GenerateID() + "_" + fileHeader.Filename
    _, err = uploader.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(fileName),
        Body:   bytes.NewReader(buffer),
    })
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, fileName), nil
}
