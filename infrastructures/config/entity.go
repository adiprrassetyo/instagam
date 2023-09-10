package config

type App struct {
	Mode       string
	Name       string
	Port       string
	Url        string
	Secret_key string
}

type Database struct {
	Host     string
	Name     string
	Username string
	Password string
	Port     string
}

type LoadConfig struct {
	App      App
	Database Database
}
