package file

import (
	"io/ioutil"
	"os"

	"github.com/gocarina/gocsv"
)

type csvHandler struct{}

func NewCSVHandler() *csvHandler {
	return &csvHandler{}
}

func (h *csvHandler) ReadCSVFile(fileName string, outData interface{}) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return gocsv.UnmarshalBytes(bytes, outData)
}

func (h *csvHandler) WriteCSVFile(fileName string, inData interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = gocsv.MarshalFile(inData, file)
	if err != nil {
		return err
	}

	return nil
}
