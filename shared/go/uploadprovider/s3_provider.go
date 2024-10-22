package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/caovanhoang63/hiholive/shared/core"
	"log"
	"net/http"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*core.File, error) {
	// fileBytes is an io reader to read data
	fileBytes := bytes.NewReader(data)

	fileType := http.DetectContentType(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return nil, err
	}

	file := &core.File{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return file, nil
}

// NewS3Provider creates a new s3Provider
func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			""),             // Token
	})
	if err != nil {
		log.Fatal(err)
	}
	provider.session = s3Session
	return provider
}

// SaveImageUploaded receives data and stores it into aws s3
func (provider *s3Provider) SaveImageUploaded(ctx context.Context, data []byte, dst string) (*core.Image, error) {
	// fileBytes is an io reader to read data
	fileBytes := bytes.NewReader(data)

	fileType := http.DetectContentType(data)

	_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	//req, _ := s3.New(provider.session).PutObjectRequest(&s3.PutObjectInput{
	//	Bucket: aws.String(provider.bucketName),
	//	Key:    aws.String(dst),
	//	ACL:    aws.String("private"),
	//})
	//
	//req.Presign(time.Second * 5)

	if err != nil {
		return nil, err
	}

	img := &core.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
