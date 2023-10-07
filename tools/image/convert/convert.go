package convert

import (
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type Convert interface {
	Convert(filename string) (string, error)
}

type ConvertImpl struct {
}

func (obj *ConvertImpl) Convert(filename string) (string, error) {
	jpgFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func(jpgFile *os.File) {
		_ = jpgFile.Close()
	}(jpgFile)

	img, err := jpeg.Decode(jpgFile)
	if err != nil {
		return "", err
	}

	fileName := filepath.Base(filename)
	outputFileName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".png"
	pngFile, err := os.Create(outputFileName)
	if err != nil {
		return "", err
	}
	defer func(pngFile *os.File) {
		_ = pngFile.Close()
	}(pngFile)

	err = png.Encode(pngFile, img)
	if err != nil {
		return "", err
	}

	output, _ := filepath.Abs(outputFileName)
	return output, nil
}
