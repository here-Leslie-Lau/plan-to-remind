package db

type DatabaseCfg struct {
	// db username
	Username string `json:"username"`
	// db password
	Password string `json:"password"`
	// db url
	Url string `json:"url"`
	// db port
	Port string `json:"port"`
	// database name
	DatabaseName string `json:"database_name"`
	// db MaxIdleConn
	MaxIdleConn int
	// db MaxOpenConn
	MaxOpenConn int
	// db ConnMaxLifeTime
	ConnMaxLifeTime string
}
