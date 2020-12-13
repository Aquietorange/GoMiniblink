package GoMiniblink

type ProxyType int

const (
	ProxyType_NONE ProxyType = iota
	ProxyType_HTTP
	ProxyType_SOCKS4
	ProxyType_SOCKS4A
	ProxyType_SOCKS5
	ProxyType_SOCKS5HOSTNAME
)

type ProxyInfo struct {
	Type     ProxyType
	HostName string
	Port     int
	UserName string
	Password string
}
