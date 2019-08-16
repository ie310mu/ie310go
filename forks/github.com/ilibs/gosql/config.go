package gosql

//Config is database connection configuration
type Config struct {
	Enable       bool   `json:"enable"`
	Driver       string `json:"driver"`
	Dsn          string `json:"dsn"`
	MaxOpenConns int    `toml:"max_open_conns" json:"max_open_conns"` //最大打开的连接数
	MaxIdleConns int    `toml:"max_idle_conns" json:"max_idle_conns"` //最大闲置连接数
	MaxLifetime  int    `toml:"max_life_time" json:"max_life_time"`   //最大闲置时间，单位s，一般设置为mysql waittime的一半
	ShowSql      bool   `toml:"show_sql" json:"show_sql"`
}
