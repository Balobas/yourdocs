package templates

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"os"
)

func CreateTempTemplate(doc []byte) (*os.File, string, error) {
	name, err := uuid.NewV4()
	if err != nil {
		return nil, "", err
	}
	filename := "app/templates/" + name.String() + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		return nil, "", errors.New("Cant create file")
	}
	_, err = file.Write(doc)
	if err != nil {
		return nil, "", errors.New("Cant write file")
	}
	return file, filename, nil
}
