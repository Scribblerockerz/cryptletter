package attachment

type Handler interface {
	HostType() string
	Put(fileData string) (string, error)
	Get(identifier string) (string, error)
	Delete(identifier string) error
	SetTTL(identifier string, ttl int64) error
	Cleanup() error
	DropAll() error
	ListTimetable() ([]string, error)
}
