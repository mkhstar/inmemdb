package db

import (
	"encoding/json"
	"time"
)

// StoreData type of data
type StoreData map[string]interface{}

// Store represents the data on disk
type Store struct {
	Data         StoreData            `json:"data"`
	KeysToExpire map[string]time.Time `json:"keysToExpire"`
}

// UnmarshalJSON implements Unmarshaler
func (storeData *StoreData) UnmarshalJSON(p []byte) error {
	*storeData = make(map[string]interface{})
	var container = make(map[string]interface{})
	json.Unmarshal(p, &container)

	for k, v := range container {
		switch value := v.(type) {
		case string:
			(*storeData)[k] = value
		case []interface{}:
			listValues := make([]string, 0, len(value))
			for _, iv := range value {
				if strval, strvalOk := iv.(string); strvalOk {
					listValues = append(listValues, strval)
				}
			}
			(*storeData)[k] = dlist(listValues)
		case map[string]interface{}:
			setValue := make(dset)
			for ik := range value {
				setValue[ik] = struct{}{}
			}
			(*storeData)[k] = setValue

		}
	}

	return nil
}
