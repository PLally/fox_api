package foxes

import (
	"bytes"
	"encoding/json"
	"github.com/plally/random_fox_api/storage"
	"image"
)


func jsonFormat(s storage.Storage, id string, format string) ([]byte, error) {
	return json.Marshal(map[string]string{
		"id": id,
	})
}

func imageFormat(s storage.Storage, id string, format string) ([]byte, error) {
	data, err := s.Retrieve(id)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(data)
	_, storedFormat, err := image.DecodeConfig(r)
	if err != nil { return nil, err }

	if format == storedFormat {
		return data, nil
	}

	r = bytes.NewReader(data)
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	encoder, ok := encoders[format]
	if !ok {
		return nil, InvalidFormatErr
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	encoder(buf, img)

	return buf.Bytes(), nil
}