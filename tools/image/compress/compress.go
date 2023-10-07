package compress

import (
	"image/jpeg"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Compress interface {
	Compress(filename string) (string, error)
}

type CompressImpl struct {
	Quality int
}

func (obj *CompressImpl) Compress(filename string) (string, error) {
	inputFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func(inputFile *os.File) {
		_ = inputFile.Close()
	}(inputFile)

	fileName := filepath.Base(filename)
	outputFileName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "_" + strconv.Itoa(obj.Quality) + filepath.Ext(filename)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer func(outputFile *os.File) {
		_ = outputFile.Close()
	}(outputFile)

	options := jpeg.Options{
		Quality: obj.Quality,
	}

	img, err := jpeg.Decode(inputFile)
	if err != nil {
		return "", err
	}

	err = jpeg.Encode(outputFile, img, &options)
	if err != nil {
		return "", err
	}

	output, _ := filepath.Abs(outputFileName)
	return output, nil
}
