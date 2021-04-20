# inmemdb  [![Coverage Status](https://coveralls.io/repos/github/mkhstar/inmemdb/badge.svg?branch=master)](https://coveralls.io/github/mkhstar/inmemdb?branch=master) [![GoDoc](https://godoc.org/github.com/mkhstar/inmemdb?status.svg)](http://godoc.org/github.com/mkhstar/inmemdb)

inmemdb is a redis API compatible in-memory database heavily inspired by redis.


## Usage

**inmemdb** runs on an HTTP server on the default port **9005**. The default port can be changed with an additional command-line argument. Example below


```bash
./inmemdb

# ./inmemdb 5000 - runs on port 5000
```

## HTTP Request

```
POST / HTTP/1.1
Host: http://localhost:9005
Content-Type: application/json

{"outputFormat": "json", "command": ["get", "name"]}
```

## Commands Supported

Commands supported are listed below. More to come!

- PING
- GET
- SET
- SETEX
- DEL
- EXPIRE
- TTL
- SADD
- SMEMBERS
- SREM
- LPUSH
- RPUSH
- LLEN
- LPOP
- LRANGE
- EXIT


## Client

Currently, there are no clients available for inmemdb. In a short period of time, clients for Nodejs, Go, PHP, and Python will be released.

- [x] [Nodejs](https://github.com/mkhstar/inmemdb-client)
- [ ] Go
- [ ] Python
- [ ] PHP


## Outputs

Inmemdb can return its response in two formats (JSON and text). Response formats such as XML may be supported later. You can specify the format using the outputFormat argument. The default format is in JSON.


## Examples

As already discussed above, the keys supported (See above) are compatible with redis.

Note: The commands are **case-insensitive** and are specified in an array of strings (See above http request)


|             Command   |Normal Output                          |JSON Output                         |
|----------------|-------------------------------|-----------------------------|
|`PING`|`OK`           |`{"result":"PONG","status":"success"}`
|`SET name musah`|`OK`           |`{"result":"OK","status":"success"}`
|`GET name`          |`musah`            |`{"result":"musah","status":"success"}` 
|`EXPIRE name 100`          |`OK`|`{"result":"OK","status":"success"}`
|`TTL name`          |`55`|`{"result":55","status":"success"}`
|`DEL name`          |`OK`|`{"result":"OK","status":"success"}`
|`SADD names musah`          |`OK`|`{"result":"OK","status":"success"}`
|`SADD names musah`          |`OK`|`{"result":"OK","status":"success"}`
|`SADD names kusi`          |`OK`|`{"result":"OK","status":"success"}`
|`SADD names hussein`          |`OK`|`{"result":"OK","status":"success"}`
|`SMEMBERS names`          |`1) musah` <br/> `2) kusi` <br> `3) hussein`|`{"result":["musah","kusi","hussein"],"status":"success"}`
|`RPUSH sentence works`          |`OK`|`{"result":"OK","status":"success"}`
|`LPUSH sentence just`          |`OK`|`{"result":"OK","status":"success"}`
|`LPUSH sentence it`          |`OK`|`{"result":"OK","status":"success"}`
|`LRANGE sentence 0 -1`          |`1) it` <br/> `2) just` <br> `3) works`|`{"result":["it","just","works"],"status":"success"}`
|`LPOP sentence`          |`works`|`{"result":"just","status":"success"}`
|`LRANGE sentence 0 -1`          |`1) it` <br/> `2) just`|`{"result":["it","just"],"status":"success"}`




## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## License
[MIT](https://choosealicense.com/licenses/mit/)
