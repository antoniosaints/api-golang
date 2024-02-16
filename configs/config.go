package configs

var cfg *config

type config struct {
	Api ApiStruct
}
type ApiStruct struct {
	Prefix   string
	DbName   string
	Host     string
	Port     string
	Username string
	Password string
}

func init() {
	cfg = new(config)

	cfg.Api = ApiStruct{
		Prefix:   "/api/v1",
		DbName:   "test",
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "",
	}
}

func GetConfig() ApiStruct {
	return cfg.Api
}
