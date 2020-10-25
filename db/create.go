package db

import (
	"errors"

	"github.com/mkhstar/inmemdb/result"
)

func set(args []string) *result.Info {

	key, value := args[0], args[1]
	inMemDB.mu.Lock()
	inMemDB.data[key] = value
	inMemDB.mu.Unlock()

	return &result.Info{
		Result: "OK",
	}
}

func setExpire(args []string) *result.Info {
	key, seconds, value := args[0], args[1], args[2]
	inMemDB.mu.Lock()
	inMemDB.data[key] = value
	expire(key, seconds)
	inMemDB.mu.Unlock()
	return &result.Info{
		Result: "OK",
	}
}

func sadd(args []string) *result.Info {
	key, setKey := args[0], args[1]

	if v, ok := inMemDB.data[key]; ok {
		if setMap, keyOk := v.(dset); keyOk {
			inMemDB.mu.Lock()
			setMap[setKey] = struct{}{}
			inMemDB.mu.Unlock()
			return &result.Info{
				Result: "OK",
			}
		}
		return &result.Info{
			Error: errors.New("WRONGTYPE Operation against a key holding the wrong kind of value"),
		}
	}

	inMemDB.mu.Lock()
	inMemDB.data[key] = dset{setKey: struct{}{}}
	inMemDB.mu.Unlock()
	return &result.Info{
		Result: "OK",
	}

}

func rpush(args []string) *result.Info {
	key, value := args[0], args[1]

	if v, ok := inMemDB.data[key]; ok {
		if list, keyOk := v.(dlist); keyOk {
			list = append(list, value)
			inMemDB.mu.Lock()
			inMemDB.data[key] = list
			inMemDB.mu.Unlock()
			return &result.Info{
				Result: "OK",
			}
		}
		return &result.Info{
			Error: errors.New("WRONGTYPE Operation against a key holding the wrong kind of value"),
		}
	}

	return initializeList(key, value)
}

func lpush(args []string) *result.Info {
	key, value := args[0], args[1]

	if v, ok := inMemDB.data[key]; ok {
		if list, keyOk := v.(dlist); keyOk {
			list = append(dlist{value}, list...)
			inMemDB.mu.Lock()
			inMemDB.data[key] = list
			inMemDB.mu.Unlock()
			return &result.Info{
				Result: "OK",
			}
		}
		return &result.Info{
			Error: errors.New("WRONGTYPE Operation against a key holding the wrong kind of value"),
		}
	}

	return initializeList(key, value)
}

func initializeList(key string, value string) *result.Info {
	list := dlist{value}
	inMemDB.mu.Lock()
	inMemDB.data[key] = list
	inMemDB.mu.Unlock()
	return &result.Info{
		Result: "OK",
	}
}
