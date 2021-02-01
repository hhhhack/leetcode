module routes

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/spf13/viper v1.7.1 // indirect
	models v0.0.0-00010101000000-000000000000
)

replace models => ../models
