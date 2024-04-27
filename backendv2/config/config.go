package config

//Properties Configuration properties based on env variables.
type Properties struct {
	Port              string `env:"MY_APP_PORT" env-default:"8080"`
	Host              string `env:"HOST" env-default:"localhost"`
	DBHost            string `env:"DB_HOST" env-default:"localhost"`
	DBPort            string `env:"DB_PORT" env-default:"27017"`
	DBName            string `env:"DB_NAME" env-default:"HDB"`
	ProductCollection string `env:"DB_COL_NAME" env-default:"records"`
	UsersCollection   string `env:"USERS_COL_NAME" env-default:"users"`
	JwtTokenSecret    string `env:"JWT_TOKEN_SECRET" env-default:"abrakadabra"`
}