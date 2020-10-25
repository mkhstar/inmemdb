package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mkhstar/inmemdb/result"
)

func expire(key string, seconds string) *result.Info {
	if _, ok := inMemDB.data[key]; ok {
		secondsInt, _ := strconv.Atoi(seconds)
		inMemDB.keysToExpire[key] = time.Now().Add(time.Duration(secondsInt) * time.Second)
		return &result.Info{
			Result: "OK",
		}
	}
	return &result.Info{
		Error: fmt.Errorf("Key %s was not found", key),
	}
}

func isExpired(key string) bool {
	if val, ok := inMemDB.keysToExpire[key]; ok {
		if int(time.Since(val)) > 0 {
			delete(inMemDB.keysToExpire, key)
			delete(inMemDB.data, key)
			return true
		}
		return false
	}
	return false
}

func ttl(key string) *result.Info {
	if val, ok := inMemDB.keysToExpire[key]; ok {
		value := int(val.Sub(time.Now()).Seconds())
		if value <= -1 {
			return &result.Info{
				Result: -1,
			}
		}
		return &result.Info{
			Result: value,
		}

	}
	return &result.Info{
		Result: -2,
	}
}
