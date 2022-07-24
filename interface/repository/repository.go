package repository

type CSVHandler interface {
	ReadCSVFile(fileName string, outData interface{}) error
	WriteCSVFile(fileName string, inData interface{}) error
}
