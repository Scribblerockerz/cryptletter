package attachment

import (
	"context"
	"errors"
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/logger"
	"github.com/Scribblerockerz/cryptletter/pkg/utils"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type s3Handler struct {
	list       utils.DismissiveList
	defaultTTL int64
	bucketName string
	minio      *minio.Client
}

//Put will place data into the storage directory
func (l s3Handler) Put(fileData string) (string, error) {

	tmpFile, err := ioutil.TempFile(os.TempDir(), "att-")
	if err != nil {
		return "", err
	}

	if _, err = tmpFile.Write([]byte(fileData)); err != nil {
		return "", errors.New("failed to write to temporary file")
	}

	ctx := context.Background()
	identifier := uuid.New().String()

	_, err = l.minio.FPutObject(ctx, l.bucketName, identifier, tmpFile.Name(), minio.PutObjectOptions{}) // ContentType: "blob"
	if err != nil {
		return "", err
	}

	tmpFile.Close()
	os.Remove(tmpFile.Name())

	// Add file to the list of tracked resources
	err = l.list.Add(identifier, l.defaultTTL)
	if err != nil {
		return "", err
	}

	return identifier, nil
}

//Get will retrieve file data by a given identifier
func (l s3Handler) Get(identifier string) (string, error) {
	// Add file to the list of tracked resources
	if !l.list.Has(identifier) {
		fmt.Println("asdasdasd")
		return "", errors.New("unable to read file by identifier")
	}

	fmt.Println("8888")

	info, err := l.minio.GetObject(
		context.Background(),
		l.bucketName,
		identifier,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return "", err
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, info)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

//Delete will remove stored files by identifier
func (l s3Handler) Delete(identifier string) error {
	// Delete identifier from tracked resources
	l.list.Del(identifier)

	err := l.minio.RemoveObject(
		context.Background(),
		l.bucketName,
		identifier,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		// TODO: may be return nil?
		return err
	}

	return nil
}

//SetTTL will update the TTL of an identifier
func (l s3Handler) SetTTL(identifier string, ttl int64) error {
	return l.list.Set(identifier, ttl)
}

//Cleanup will sync the known file identifiers with the unknown and clean them up
func (l s3Handler) Cleanup() error {

	trackedIdentifier, err := l.list.All()
	if err != nil {
		return err
	}

	objectsCh := make(chan minio.ObjectInfo)

	// Send object names that are needed to be removed to objectsCh
	go func() {
		defer close(objectsCh)
		for object := range l.minio.ListObjects(context.Background(), l.bucketName, minio.ListObjectsOptions{}) {
			if object.Err != nil {
				logger.LogError(object.Err)
			}
			if !utils.ListContainsString(trackedIdentifier, object.Key) {
				objectsCh <- object
			}
		}
	}()

	opts := minio.RemoveObjectsOptions{
		GovernanceBypass: true,
	}

	for rErr := range l.minio.RemoveObjects(context.Background(), l.bucketName, objectsCh, opts) {
		logger.LogError(rErr.Err)
	}

	return nil
}

//DropAll will clear the storage dir and the list
func (l s3Handler) DropAll() error {

	objectsCh := make(chan minio.ObjectInfo)

	// Send object names that are needed to be removed to objectsCh
	go func() {
		defer close(objectsCh)
		for object := range l.minio.ListObjects(context.Background(), l.bucketName, minio.ListObjectsOptions{}) {
			if object.Err != nil {
				logger.LogError(object.Err)
			}
			objectsCh <- object
		}
	}()

	opts := minio.RemoveObjectsOptions{
		GovernanceBypass: true,
	}

	for rErr := range l.minio.RemoveObjects(context.Background(), l.bucketName, objectsCh, opts) {
		logger.LogError(rErr.Err)
	}

	err := l.minio.RemoveBucket(context.Background(), l.bucketName)
	if err != nil {
		return err
	}

	return l.list.Drp()
}

//ListTimetable will list all known entries from the dismissive list
func (l s3Handler) ListTimetable() ([]string, error) {
	return l.list.All()
}

const (
	S3HostType = "s3"
)

func (l s3Handler) HostType() string {
	return S3HostType
}

func NewS3Handler(defaultTTL int64, bucketName string, bucketRegion string, endpoint string, options *minio.Options) Handler {

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, options)
	if err != nil {
		panic(err)
	}

	err = initializeBucket(minioClient, bucketName, bucketRegion)
	if err != nil {
		panic(err)
	}

	return &s3Handler{
		list:       utils.NewDismissiveList("s3-files"),
		defaultTTL: defaultTTL,
		bucketName: bucketName,
		minio:      minioClient,
	}
}

//initializeBucket will ensure that the bucket exists and is ready to be used
func initializeBucket(mc *minio.Client, bucketName string, bucketLRegion string) error {
	ctx := context.Background()

	err := mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: bucketLRegion})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := mc.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return nil
		} else {
			return err
		}
	}

	return nil
}
