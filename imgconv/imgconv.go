package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// Convert image extension
func Convert(src, dst string, rmsrc bool) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	df, err := os.Create(dst)

	if err != nil {
		return err
	}
	defer df.Close()

	switch filepath.Ext(dst) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	case ".gif":
		err = gif.Encode(df, img, nil)
	}
	if err != nil {
		return err
	}

	if rmsrc {
		if err = Remove(src); err != nil {
			return err
		}
	}
	return nil
}

// Remove file
func Remove(src string) error {
	err := os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}
