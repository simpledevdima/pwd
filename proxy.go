package pwd

func NewProxy(id uint, ip string, port uint16) *Proxy {
	p := &Proxy{}
	p.SetID(id)
	p.SetIP(ip)
	p.SetPort(port)
	return p
}

type Proxy struct {
	id   uint
	ip   string
	port uint16
}

func (p *Proxy) SetID(id uint) {
	p.id = id
}

func (p *Proxy) SetIP(ip string) {
	p.ip = ip
}

func (p *Proxy) SetPort(port uint16) {
	p.port = port
}
