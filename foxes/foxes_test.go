package foxes

import (
	"bytes"
	"github.com/google/uuid"
	"io/ioutil"
	"testing"
)

func readFox(filepath string) []byte {
	filepath = "../testdata/" + filepath
	data, err := ioutil.ReadFile(filepath)
	if err != nil { panic(err) }
	return data
}

func isValidId(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil { return false }
	return true
}
func TestCreateFox(t *testing.T) {
	testCases := []struct{
		name string
		bytes []byte
		expectedErr error
	}{
		{"fox1.png",readFox("fox1.png"), nil},
		{"fox1.jpg",readFox("fox1.jpg"), nil},
		{"fox1.gif",readFox("fox1.gif"), InvalidFormatErr},
	}

	for _, testData := range testCases {
		t.Run(testData.name, func(t *testing.T) {
			s := dummyStore{}
			foxid, err := CreateFox(s, bytes.NewReader(testData.bytes))
			if err != testData.expectedErr {
				t.Errorf("Error: %v \n Expected: %v", err, testData.expectedErr)

			}
			if testData.expectedErr != nil { return }
			if !isValidId(foxid) {
				t.Errorf("Invalid id %v", foxid)
			}
		})
	}

}

func TestRetrieveFox(t *testing.T) {
	testCases := []struct{
		name string
		id string
		format string
		expected []byte
		expectedErr error
	}{
		{ "fox1.png","1", "png", readFox("fox1.png"), nil},
		{ "fox1.png invalid format","1", "gif", readFox("fox1.png"), InvalidFormatErr},
		{"fox1.png as jpeg", "1", "jpeg", readFox("fox1.jpg"), nil},
	}

	for _, testData := range testCases {
		t.Run(testData.name, func(t *testing.T) {
			s := dummyStore{}
			data, err := RetrieveFox(s, testData.id, testData.format)
			if err != testData.expectedErr {
				t.Error(err)

			}
			if testData.expectedErr != nil { return }
			if !bytes.Equal(data, testData.expected) {
				t.Error("Unexpected data")
			}
		})
	}
}

type dummyStore struct {}

func (d dummyStore) Create([]byte) (string, error) {
	id := uuid.New()

	return id.String(), nil
}

func (d dummyStore) Retrieve(id string) ([]byte, error) {
	switch id {
	case "1": return readFox("fox1.png"), nil
	}
	return []byte{}, nil
}

func (d dummyStore) List() ([]string, error) {return nil,nil}