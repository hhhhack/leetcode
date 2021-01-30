module go_tool

go 1.15

replace (
	initConf => ./initConf
	logger => ./logger
	sendRequest => ./sendRequest
)

require (
	initConf v0.0.0-00010101000000-000000000000
	logger v0.0.0-00010101000000-000000000000
	sendRequest v0.0.0-00010101000000-000000000000
)
