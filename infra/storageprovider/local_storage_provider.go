package storageprovider

import (
	"fmt"
	"os"

	"github.com/moises-ba/saga-pedido-ms/domain/repository"
)

type localFileStorage struct {
	basePath string
}

func NewLocalFileStorage(basePath string) repository.Storage {
	return &localFileStorage{
		basePath: basePath,
	}
}

func (ls *localFileStorage) BasePath() string {
	return ls.basePath
}

func (ls *localFileStorage) Store(content []byte, fileName string) error {

	contentChan := make(chan []byte)
	go func() {
		defer close(contentChan)
		contentChan <- content
	}()

	return ls.StoreStream(contentChan, fileName)
}

func (ls *localFileStorage) StoreStream(contentChan chan []byte, fileName string) error {

	fullFileName := fmt.Sprintf("/%v/%v", ls.basePath, fileName)
	file, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	for content := range contentChan {
		if _, err := file.WriteString(string(content)); err != nil {
			return err
		}
	}

	return nil
}
