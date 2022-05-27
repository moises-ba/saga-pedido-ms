package repository

type Storage interface {
	BasePath() string
	Store(content []byte, fileName string) error
}

type StorageStream interface {
	Storage
	StoreStream(content chan []byte, fileName string) error
}
