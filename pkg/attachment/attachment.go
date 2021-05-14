package attachment

type Handler interface {
	Put(fileData string) (string, error)
	Get(identifier string) (string, error)
	Delete(identifier string) error
}
