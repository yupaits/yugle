package appconfig

type Config struct {
	Port            string
	HtmlTemplateDir string
	StaticFolder    string
	FaviconPath     string
}

func DefaultConfig() *Config {
	config := Config{}
	config.Port = "8888"
	config.HtmlTemplateDir = "ui/dist/*.html"
	config.StaticFolder = "./ui/dist/static"
	config.FaviconPath = "./ui/dist/favicon.ico"
	return &config
}
