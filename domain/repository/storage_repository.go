package repository

type Storage interface {
	BasePath() string
	Store(content []byte, fileName string) error
	StoreStream(content chan []byte, fileName string) error
}
