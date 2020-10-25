package db

import (
	"errors"

	"github.com/mkhstar/inmemdb/command"
	"github.com/mkhstar/inmemdb/result"
)

// Execute executues a given command and returns the responses in slice of string or error
func Execute(cmd *command.Command) result.Resulter {
	switch cmd.GetType() {
	case "SET":
		return set(cmd.Args)
	case "SETEX":
		return setExpire(cmd.Args)
	case "DEL":
		return deleteKey(cmd.Args[0])
	case "EXPIRE":
		return expire(cmd.Args[0], cmd.Args[1])
	case "GET":
		return get(cmd.Args[0])
	case "TTL":
		return ttl(cmd.Args[0])
	case "SADD":
		return sadd(cmd.Args)
	case "SMEMBERS":
		return smembers(cmd.Args[0])
	case "SREM":
		return srem(cmd.Args)
	case "LPUSH":
		return lpush(cmd.Args)
	case "RPUSH":
		return rpush(cmd.Args)
	case "LLEN":
		return llen(cmd.Args[0])
	case "LPOP":
		return lpop(cmd.Args[0])
	case "LRANGE":
		return lrange(cmd.Args)

	default:
		return &result.Info{Error: errors.New("Coming Soon")}
	}
}
