package config

type Config struct {
	USERNAME string
	PASS     string
	DBNAME   string
}

func NewConfig(USERNAME, PASS, DATABASE string) (*Config, error) {
	cfg := &Config{USERNAME, PASS, DATABASE}
	return cfg, nil
}
