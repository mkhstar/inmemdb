package command

// ValidCommands shows the valid commands of inMemDB
var ValidCommands = map[string][]string{
	"DEL": []string{
		"string",
	},
	"EXPIRE": []string{
		"string",
		"int",
	},
	"GET": []string{
		"string",
	},
	"SET": []string{
		"string",
		"string",
	},
	"SETEX": []string{
		"string",
		"int",
		"string",
	},
	"TTL": []string{
		"string",
	},
	"SADD": []string{
		"string",
		"string",
	},
	"SMEMBERS": []string{
		"string",
	},
	"SREM": []string{
		"string",
		"string",
	},
	"LPUSH": []string{
		"string",
		"string",
	},
	"RPUSH": []string{
		"string",
		"string",
	},
	"LLEN": []string{
		"string",
	},
	"LPOP": []string{
		"string",
	},
	"LRANGE": []string{
		"string",
		"int",
		"int",
	},
}
