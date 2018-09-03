package appconfig

type Config struct {
	Port            string
	HtmlTemplateDir string
	StaticFolder    string
	FaviconPath     string
}

func DefaultConfig() *Config {
	config := Config{}
	config.Port = "8000"
	config.StaticFolder = "./static"
	config.HtmlTemplateDir = "templates/*"
	config.FaviconPath = "./favicon.ico"
	return &config
}
