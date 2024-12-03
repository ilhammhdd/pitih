package enum

import (
	"database/sql"
	"encoding/json"
	"errors"
)

type EnumSerder interface {
	sql.Scanner
	json.Unmarshaler
	json.Marshaler
}

func ScanEnum[K ~uint](src any, key *K, keyValue map[K]string) error {
	var valueStr string
	value, ok := src.([]byte)
	if !ok {
		valueStr, ok = src.(string)
		if !ok {
			valueInt64, ok := src.(int64)
			if !ok {
				return errors.New("not an enum value of type []byte, string, or int64")
			} else {
				*key = K(valueInt64)
				return nil
			}
		}
	} else {
		valueStr = string(value)
	}

	gotKey, ok := GetEnumKey(valueStr, keyValue)
	if !ok {
		return errors.New("enum key not found")
	}

	*key = gotKey
	return nil
}

func UnmarshalJSONEnum[K ~uint](b []byte, key *K, keyValue map[K]string) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	gotStatus, ok := GetEnumKey(value, keyValue)
	if !ok {
		return errors.New("unknown enum value")
	}

	*key = gotStatus
	return nil
}

func MarshalJSONEnum[K ~uint](key K, keyValue map[K]string) ([]byte, error) {
	value, ok := keyValue[key]
	if !ok {
		return nil, errors.New("unknown enum key")
	}

	return json.Marshal(value)
}

func GetEnumKey[K ~uint](value string, keyValue map[K]string) (K, bool) {
	for key, val := range keyValue {
		if value == val {
			return key, true
		}
	}
	return 0, false
}

func GetEnumValueFromPtr[K ~uint](key *K, keyValues map[K]string) string {
	if key == nil {
		return ""
	}
	return keyValues[*key]
}
