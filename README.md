# docker-compose-log-parser

A simple parser written in go for parsing docker-compose log running in foreground. It is designed to use in conjunction with the jq utility.

**Usage Example **

docker-compose up | parser | jq .
