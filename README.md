# docker-compose-log-parser

A simple parser written in go for parsing docker-compose log running in foreground. It is designed to use in conjunction with the jq utility. You can compile this on windows yourself. However, it is designed with linux/unix OSes in mind.

## Usage Example:

`go build`

`docker-compose up | parser | jq -R "fromjson?"`
