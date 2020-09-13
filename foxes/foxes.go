package foxes

import (
	"bytes"
	"errors"
	"github.com/plally/random_fox_api/storage"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"io"
)
// these functions wrap the storage functions to
// validate input, convert images to a standard type, and convert images to a request type

var supportedFormats = map[string]bool {
	"png": true,
	"jpeg": true,
}

func jpegEncoder(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

var encoders = map[string]func(w io.Writer, m image.Image) error {
	"png": png.Encode,
	"jpeg": jpegEncoder,
	"jpg": jpegEncoder,
}

var InvalidFormatErr = errors.New("Image format is not supported")

func normalizeImage(r io.Reader) ([]byte, error) {
	img, format, err := image.Decode(r)
	if err != nil || !supportedFormats[format] {
		return nil, InvalidFormatErr
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	png.Encode(buf, img)

	return buf.Bytes(), nil
}

func CreateFox(s storage.Storage, r io.Reader) (string, error) {
	data, err := normalizeImage(r)
	if err != nil {
		return "", err
	}

	id, err := s.Create(data)
	if err != nil {
		return id, err
	}

	return id, nil
}

func RetrieveFox(s storage.Storage, id string, format string) ([]byte, error) {
	switch format {
	case "jpg":
		fallthrough
	case "png":
		return imageFormat(s, id, format)
	case "json":
		return jsonFormat(s, id, format)
	}
	return nil, errors.New("unsupported format")
}
