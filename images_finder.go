package awesomeProject

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
)

func findImage(filepath string) (bool, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return false, err
	}

	switch img.(type) {
	case *image.RGBA, *image.NRGBA, *image.NRGBA64, *image.RGBA64:
		return true, nil
	default:
		return false, nil
	}
}

func main() {
	err := filepath.Walk("C:\\", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			findImg, err := findImage(path)
			if err != nil {
				fmt.Println("Error checking file:", err)
				return err
			}
			if findImg {
				fmt.Println("Found image:", path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the file sys:", err)
	}
}
