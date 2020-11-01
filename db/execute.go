package db

import (
	"errors"

	"github.com/mkhstar/inmemdb/command"
	"github.com/mkhstar/inmemdb/result"
)

// Execute executues a given command and returns the responses in slice of string or error
func Execute(cmd *command.Command, resulter chan result.Resulter) {
	switch cmd.GetType() {
	case "SET":
		resulter <- set(cmd.Args)
	case "SETEX":
		resulter <- setExpire(cmd.Args)
	case "DEL":
		resulter <- deleteKey(cmd.Args[0])
	case "EXPIRE":
		resulter <- expire(cmd.Args[0], cmd.Args[1])
	case "GET":
		resulter <- get(cmd.Args[0])
	case "TTL":
		resulter <- ttl(cmd.Args[0])
	case "SADD":
		resulter <- sadd(cmd.Args)
	case "SMEMBERS":
		resulter <- smembers(cmd.Args[0])
	case "SREM":
		resulter <- srem(cmd.Args)
	case "LPUSH":
		resulter <- lpush(cmd.Args)
	case "RPUSH":
		resulter <- rpush(cmd.Args)
	case "LLEN":
		resulter <- llen(cmd.Args[0])
	case "LPOP":
		resulter <- lpop(cmd.Args[0])
	case "LRANGE":
		resulter <- lrange(cmd.Args)
	case "PING":
		resulter <- ping()

	default:
		resulter <- &result.Info{Error: errors.New("Coming Soon")}
	}
}
