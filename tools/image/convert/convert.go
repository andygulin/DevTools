package convert

import (
	"github.com/chai2010/webp"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type Convert interface {
	Convert(filename string) ([]string, error)
}

type ConvertImpl struct {
}

const PNGFileExt = ".png"
const WebPFileExt = ".webp"

func (obj *ConvertImpl) Convert(filename string) ([]string, error) {
	jpgFile, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer func(jpgFile *os.File) {
		_ = jpgFile.Close()
	}(jpgFile)

	img, err := jpeg.Decode(jpgFile)
	if err != nil {
		return []string{}, err
	}

	fileName := filepath.Base(filename)

	// PNG
	outputFileName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + PNGFileExt
	pngFile, err := os.Create(outputFileName)
	if err != nil {
		return []string{}, err
	}
	defer func(pngFile *os.File) {
		_ = pngFile.Close()
	}(pngFile)

	err = png.Encode(pngFile, img)
	if err != nil {
		return []string{}, err
	}

	// WebP
	outputFileName2 := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + WebPFileExt
	webPFile, err := os.Create(outputFileName2)
	if err != nil {
		return []string{}, err
	}
	defer func(webPFile *os.File) {
		_ = webPFile.Close()
	}(webPFile)

	err = webp.Encode(webPFile, img, &webp.Options{Lossless: true})
	if err != nil {
		return []string{}, err
	}

	var output []string
	output1, err := filepath.Abs(outputFileName)
	output = append(output, output1)
	output2, err := filepath.Abs(outputFileName2)
	output = append(output, output2)
	return output, nil
}
