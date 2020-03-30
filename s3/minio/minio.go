package server

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/minio/minio-go/v6"
	err2 "github.com/tradingAI/go/error"
)

type MinioConf struct {
	AccessKey string
	SecretKey string
	Host      string
	Port      int
	Secure    bool
}

func (m MinioConf) Validate() (err error) {
	if m.AccessKey == "" {
		err = err2.ErrEmptyMinioAccessKey
		glog.Error(err)
		return
	}

	if m.SecretKey == "" {
		err = err2.ErrEmptyMinioSecretKey
		glog.Error(err)
		return
	}

	if m.Host == "" {
		err = err2.ErrEmptyMinioHost
		glog.Error(err)
		return
	}

	if m.Port <= 1024 || m.Port >= 65535 {
		err = err2.ErrInvalidMinioPort
		glog.Error(err)
		return
	}

	return
}

func NewMinioClient(conf MinioConf) (client *minio.Client, err error) {
	client, err = minio.New(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		conf.AccessKey,
		conf.SecretKey,
		conf.Secure)

	if err != nil {
		glog.Error(err)
		return
	}

	return
}

func MinioUpload(client *minio.Client, bucket string, fp string, objName string, contentType string) (err error) {
	location := "us-east-1"
	exists, err := client.BucketExists(bucket)
	if err != nil {
		glog.Error(err)
		return
	}

	if !exists {
		err = client.MakeBucket(bucket, location)
		if err != nil {
			glog.Error(err)
			return
		}

		glog.Infof("Successfully created bucket [%s]", bucket)
	}

	n, err := client.FPutObject(bucket, objName, fp, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Infof("Successfully uploaded %s of size %d", objName, n)

	return
}

func MinioDownload(client *minio.Client, bucket string, fp string, objName string) (err error) {
	err = client.FGetObject(bucket, objName, fp, minio.GetObjectOptions{})
	if err != nil {
		glog.Error(err)
		return
	}

	return
}
