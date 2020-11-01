package db

import (
	"errors"
	"strconv"

	"github.com/mkhstar/inmemdb/result"
)

func get(key string) *result.Info {
	if len(inMemDB.data) == 0 {
		return &result.Info{
			Result: nil,
		}
	}
	val, ok := inMemDB.data[key]
	if ok == false || isExpired(key) {
		return &result.Info{
			Result: nil,
		}
	}
	if stringVal, stringValOk := val.(string); stringValOk {
		return &result.Info{
			Result: stringVal,
		}
	}
	return &result.Info{
		Error: errors.New("WRONGTYPE Operation against a key holding the wrong kind of value"),
	}

}

func smembers(key string) *result.Info {
	if v, ok := inMemDB.data[key]; ok {
		if setMap, keyOk := v.(dset); keyOk {
			setSize := len(setMap)
			setSlice := make([]string, 0, setSize)
			for mapKey := range setMap {
				setSlice = append(setSlice, mapKey)
			}
			return &result.Info{
				Result: setSlice,
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

func llen(key string) *result.Info {
	if v, ok := inMemDB.data[key]; ok {
		if list, keyOk := v.(dlist); keyOk {
			return &result.Info{
				Result: len(list),
			}
		}
		return &result.Info{
			Error: errors.New("WRONGTYPE Operation against a key holding the wrong kind of value"),
		}
	}
	return &result.Info{
		Result: 0,
	}
}

func ping() *result.Info {
	return &result.Info{Result: "PONG"}
}

func lrange(args []string) *result.Info {
	key := args[0]
	start, _ := strconv.Atoi(args[1])
	end, _ := strconv.Atoi(args[2])

	if v, ok := inMemDB.data[key]; ok {
		if list, keyOk := v.(dlist); keyOk {
			sliceLength := len(list)
			if sliceLength == 0 {
				return &result.Info{
					Result: [0]string{},
				}
			}
			if end < 0 {
				end = (sliceLength + end) + 1
			}

			if start < 0 {
				start = sliceLength + start
			}

			if end > sliceLength {
				end = sliceLength
			}
			if start >= end {
				return &result.Info{
					Result: [0]string{},
				}
			}

			return &result.Info{
				Result: []string(list[start:end]),
			}

		}
		return &result.Info{
			Error: errors.New("WRONGTYPE Operation against a key holding the wrong kind of value"),
		}
	}
	return &result.Info{
		Result: [0]string{},
	}
}
