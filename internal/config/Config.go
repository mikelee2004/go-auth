package config

type DatabaseConfig struct {
	DBHost string
	DBPort string
	DBName string
	DBPassword string
	DBUser string
}

func LoadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DBHost: "localhost",
		DBPort: "5432",
		DBName: "auth",
		DBPassword: "3791",
		DBUser: "postgres",
	}
}