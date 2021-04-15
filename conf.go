package cat

type Conf struct {
	Flag       bool   `json:"flag"`
	AppId      string `json:"app_id"`
	Port       int    `json:"port"`
	HttpPort   int    `json:"http_port"`
	ServerAddr string `json:"server_addr"`
	IsDebug    bool   `json:"is_debug"`
}
