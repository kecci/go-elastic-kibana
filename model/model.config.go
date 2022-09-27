package model

type Config struct {
	Env    string `json:"env"`
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
	Elasticsearch struct {
		Addresses             []string `json:"addresses"`
		Username              string   `json:"username"`
		Password              string   `json:"password"`
		MaxIdleConnsPerHost   int      `json:"max_idle_conns_per_host"`
		ResponseHeaderTimeout int      `json:"response_header_timeout"`
		DialContextTimeout    int      `json:"dial_context_timeout"`
	} `json:"elasticsearch"`
	Tmdb struct {
		ApiKey string `json:"api_key"`
	} `json:"tmdb"`
}
