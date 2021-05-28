package attachment

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

//NewAttachmentHandler will create a new handler based on the given type
func NewAttachmentHandler(hostType string) Handler {

	if hostType == "" {
		return nil
	}

	viper.SetDefault("app.default_message_ttl", 43830)
	defaultTTL := viper.GetInt64("app.default_message_ttl")

	if hostType == S3HostType {
		viper.SetDefault("s3.bucket_name", "cryptletter-attachments")
		viper.SetDefault("s3.bucket_region", "eu-central-1")

		endpoint := viper.GetString("s3.endpoint")
		accessKeyID := viper.GetString("s3.access_id")
		secretAccessKey := viper.GetString("s3.access_secret")
		bucketName := viper.GetString("s3.bucket_name")
		bucketRegion := viper.GetString("s3.bucket_region")
		isSecure := viper.GetBool("s3.secure")

		// Initialize minio client object.
		return NewS3Handler(defaultTTL, bucketName, bucketRegion, endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: isSecure,
		})
	}

	if hostType == LocalHostType {
		viper.SetDefault("app.attachments.storage_path", "cryptletter-uploads")
		localStoragePath := viper.GetString("app.attachments.storage_path")

		return NewLocalTempHandler(defaultTTL, localStoragePath) // 30 days
	}

	panic("only '" + LocalHostType + "' and '" + S3HostType + "' drivers are supported")
}
