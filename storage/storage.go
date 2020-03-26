package storage

type Order uint8


// defines an interface for storing images
type Storage interface {
	Create(imageData []byte) (id string, err error)
	Retrieve(id string) (imageData []byte, err error)
	List() ([]string, error)
}
