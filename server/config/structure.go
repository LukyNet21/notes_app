package config

type Config struct {
	Host      string         `toml:"host"`
	Port      int            `toml:"port"`
	DebugMode bool           `toml:"debug_mode"`
	Database  DatabaseConfig `toml:"database"`
	JWT       JWTConfig      `toml:"jwt"`
}

type DatabaseConfig struct {
	Type             string `toml:"type" comment:"Chose between 'sqlite' and 'mysql'"`
	DatabaseFileName string `toml:"database_file_name"`
	User             string `toml:"user,commented"`
	Password         string `toml:"password,commented"`
	DBName           string `toml:"db_name,commented"`
	DBHost           string `toml:"db_host,commented"`
	DBPort           int    `toml:"db_port,commented"`
}

type JWTConfig struct {
	Secret string `toml:"secret"`
}
