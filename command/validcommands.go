package command

// ValidCommands shows the valid commands of inMemDB
var ValidCommands = map[string][]string{
	"PING": nil,
	"DEL": {
		"string",
	},
	"EXPIRE": {
		"string",
		"int",
	},
	"GET": {
		"string",
	},
	"SET": {
		"string",
		"string",
	},
	"SETEX": {
		"string",
		"int",
		"string",
	},
	"TTL": {
		"string",
	},
	"SADD": {
		"string",
		"string",
	},
	"SMEMBERS": {
		"string",
	},
	"SREM": {
		"string",
		"string",
	},
	"LPUSH": {
		"string",
		"string",
	},
	"RPUSH": {
		"string",
		"string",
	},
	"LLEN": {
		"string",
	},
	"LPOP": {
		"string",
	},
	"LRANGE": {
		"string",
		"int",
		"int",
	},
}
