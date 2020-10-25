package db

import (
	"errors"

	"github.com/mkhstar/inmemdb/result"
)

func deleteKey(key string) *result.Info {
	delete(inMemDB.data, key)
	delete(inMemDB.keysToExpire, key)
	return &result.Info{
		Result: "OK",
	}
}

func srem(args []string) *result.Info {
	key, setKey := args[0], args[1]

	if v, ok := inMemDB.data[key]; ok {
		if setMap, keyOk := v.(dset); keyOk {
			delete(setMap, setKey)
			return &result.Info{
				Result: "OK",
			}
		}
		return &result.Info{
			Error: errors.New("WRONGTYPE Operation against a key holding the wrong kind of value"),
		}
	}
	return &result.Info{
		Result: "OK",
	}
}

func lpop(key string) *result.Info {
	if v, ok := inMemDB.data[key]; ok {
		if list, keyOk := v.(dlist); keyOk {
			if len(list) == 0 {
				return &result.Info{
					Result: nil,
				}
			}

			lastElement := list[len(list)-1]

			inMemDB.mu.Lock()
			inMemDB.data[key] = list[:(len(list) - 1)]
			inMemDB.mu.Unlock()

			return &result.Info{
				Result: lastElement,
			}
		}
		return &result.Info{
			Error: errors.New("WRONGTYPE Operation against a key holding the wrong kind of value"),
		}
	}
	return &result.Info{
		Result: nil,
	}
}
