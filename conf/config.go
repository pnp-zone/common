package conf

type Database struct {
	Driver   string
	Port     uint16
	Host     string
	Name     string
	User     string
	Password string
}

type AllowedHost struct {
	Host  string
	Https bool
}

type Server struct {
	ListenAddress           string
	AllowedHosts            []AllowedHost
	UseForwardedProtoHeader bool
	TemplatePath            string
	StaticPath              string
	PluginPath              string
}

type Config struct {
	Server   Server
	Database Database
}

func (c *Config) CheckConfig() error {
	return nil
}
