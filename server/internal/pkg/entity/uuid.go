package entity

import (
	"errors"

	uuidPkg "github.com/google/uuid"
)

const lenWithDoubleQuotes = 38

type UUIDBytes []byte

func (uuidBytes UUIDBytes) MarshalJSON() ([]byte, error) {
	if len(uuidBytes) != 16 {
		return []byte{'"', '"'}, nil
	}
	uuidStr := uuidPkg.UUID(uuidBytes).String()

	result := make([]byte, lenWithDoubleQuotes)
	result[0] = '"'
	result[lenWithDoubleQuotes-1] = '"'

	copied := copy(result[1:lenWithDoubleQuotes-1], []byte(uuidStr))
	if copied != lenWithDoubleQuotes-2 {
		return nil, errors.New("copied != lenWithDoubleQuotes")
	}

	return result, nil
}

func (uuidBytes *UUIDBytes) UnmarshalJSON(uuidStr []byte) error {
	uuid, err := uuidPkg.Parse(string(uuidStr))
	if err != nil {
		return err
	}
	(*uuidBytes) = uuid[:]
	return nil
}
