package config

//HttpConfig contains all configuration of Gin
type HttpConfig struct {
	BindAddress string `default:":8080" env:"BIND_ADDRESS"`
	Mode        string `default:"debug" env:"MODE"`
}
