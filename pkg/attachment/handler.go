package attachment

import "github.com/spf13/viper"


//NewAttachmentHandler will create a new handler based on the given type
func NewAttachmentHandler(hostType string) Handler {

	if hostType == "" {
		return nil
	}

	if hostType != LocalHostType {
		panic("only '" + LocalHostType + "' driver is supported")
	}

	localStoragePath := viper.GetString("app.attachments.storage_path")
	if localStoragePath == "" {
		localStoragePath = "cryptletter-uploads"
	}

	return NewLocalTempHandler(30 * 24 * 60 * 60, localStoragePath) // 30 days
}
