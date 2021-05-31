module main

go 1.15

replace (
	common => ./common
	myxml => ./myxml
)

require (
	common v0.0.0-00010101000000-000000000000 // indirect
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	myxml v0.0.0-00010101000000-000000000000
)
