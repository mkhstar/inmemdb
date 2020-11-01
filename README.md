# inmemdb  [![Coverage Status](https://coveralls.io/repos/github/mkhstar/inmemdb/badge.svg?branch=master)](https://coveralls.io/github/mkhstar/inmemdb?branch=master) [![GoDoc](https://godoc.org/github.com/mkhstar/inmemdb?status.svg)](http://godoc.org/github.com/mkhstar/inmemdb)

inmemdb is an redis API compatible in-memory database heavily inspired by redis.

In contrast with redis, inmemdb has 2 formats of output. The normal output and the json output. The normal output is similar to the standard output of redis. The json output returns the response in json format, **This helps clients to connect easily without the need of any sort of parsing.**


## Usage

**inmemdb** runs on a TCP server on the default port **9005**. The default port can be changed with an additional command-line argument. Example below


```bash
./inmemdb

# ./inmemdb 5000 - runs on port 5000
```

## TCP Dialing

When the tcp server starts running, clients can start dailing to interact with the server. The most popular client I recommended is **_telnet_**.

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


## Programming Languages Client

Currently, there are no clients available for inmemdb. In a short period of time, clients for Nodejs, Go, PHP, and Python will be released.

- [ ] Nodejs
- [ ] Go
- [ ] Python
- [ ] PHP


## Flags

For now only `--json` flag is supported. if the client include this flag the output will be in json format, else it will be in the default format

Example:

```bash
    get name
    musah

    get name --json
    {"result":"musah","status":"success"}
```


## Examples

As already discussed above, the keys supported (See above) are compatible with redis.

Note: The commands are **case-insensitive**


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