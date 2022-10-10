package filesprovider

import (
	"fmt"
	"io"
	"sample_app/model"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileProvider interface {
	UploadFile(name string, body io.Reader) error
	GetDownloadUrl(name string) (string, error)
	StatFile(path string) (*model.FileProviderInfo, error)
}

type S3FileProvider struct {
	client   *s3.S3
	uploader *s3manager.Uploader
	basePath string
	bucket   string
}

func NewS3FileProvider(sess *session.Session, basePath string, bucket string) *S3FileProvider {
	s3 := s3.New(sess)
	uploader := s3manager.NewUploader(sess)
	return &S3FileProvider{
		client:   s3,
		uploader: uploader,
		basePath: basePath,
		bucket:   bucket,
	}
}

func (s3p *S3FileProvider) UploadFile(name string, body io.Reader) error {
	_, err := s3p.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3p.bucket),
		Key:    aws.String(getFullPath(s3p.basePath, name)),
		Body:   body,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s3p *S3FileProvider) GetDownloadUrl(name string) (string, error) {
	req, _ := s3p.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3p.bucket),
		Key:    aws.String(getFullPath(s3p.basePath, name)),
	})
	return req.Presign(5 * time.Minute)
}

func (s3p *S3FileProvider) StatFile(file string) (*model.FileProviderInfo, error) {
	res, err := s3p.client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(s3p.bucket),
		Key:    aws.String(getFullPath(s3p.basePath, file)),
	})
	if err != nil {
		return nil, err
	}
	return &model.FileProviderInfo{Name: file, At: *res.LastModified, Size: *res.ContentLength, Uid: *res.ETag}, nil
}

func getFullPath(base string, filename string) string {
	return fmt.Sprintf("%s/%s", base, filename)
}
