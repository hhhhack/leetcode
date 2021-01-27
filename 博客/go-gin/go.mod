module go-gin

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	gopkg.in/ini.v1 v1.62.0
	routes v0.0.0-00010101000000-000000000000
)

replace routes => ./routes
