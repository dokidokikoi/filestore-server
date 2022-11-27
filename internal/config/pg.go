package config

type PGConfig struct {
	Host     string
	Port     int
	Database string
	UserName string
	Password string
	TimeZone string
}

const (
	pgKey string = "postgresql"
)

var PgConfig = &PGConfig{
	Port: 5432, Host: "127.0.0.1", Database: "postgres", Password: "postgres", TimeZone: "Asia/Shanghai",
}

func init() {
	pgConfig := GetSpecConfig(pgKey)
	if pgConfig != nil {
		pgConfig.Unmarshal(PgConfig)
	}
}
