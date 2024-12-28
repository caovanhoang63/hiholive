package uploadprovider

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"log"
	"net/http"
)

type s3Provider struct {
	id         string
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
	s3Client   *s3.S3
}

func (provider *s3Provider) ID() string {
	return provider.id
}

// NewS3Provider creates a new s3Provider
func NewS3Provider(id string) *s3Provider {
	return &s3Provider{
		id: id,
	}
}

func (provider *s3Provider) InitFlags() {
	flag.StringVar(&provider.bucketName, "s3-upload-bucket-name", "", "bucket name ")
	flag.StringVar(&provider.region, "s3-region", "", "aws region")
	flag.StringVar(&provider.apiKey, "s3-api-key", "", "aws api key")
	flag.StringVar(&provider.secret, "s3-secret", "", "aws secret key")
	flag.StringVar(&provider.domain, "cdn-domain", "", "aws cdn domain")
}

func (provider *s3Provider) Activate(serviceCtx srvctx.ServiceContext) error {
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
	provider.s3Client = s3.New(s3Session)
	return nil
}

func (provider *s3Provider) Stop() error {
	return nil
}

func (provider *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) (*core.File, error) {
	// fileBytes is an io reader to read data
	fileBytes := bytes.NewReader(data)

	fileType := http.DetectContentType(data)

	_, err := provider.s3Client.PutObject(&s3.PutObjectInput{
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

// SaveImageUploaded receives data and stores it into aws s3
func (provider *s3Provider) SaveImageUploaded(ctx context.Context, data []byte, dst string) (*core.Image, error) {
	// fileBytes is an io reader to read data

	fileBytes := bytes.NewReader(data)

	fileType := http.DetectContentType(data)

	_, err := provider.s3Client.PutObject(&s3.PutObjectInput{
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
