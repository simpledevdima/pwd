package pwd

import (
	"database/sql"
	"log"
	"net/url"
)

func NewProxies(db *sql.DB) *Proxies {
	prxs := make(Proxies)
	go func() {
		for {
			rows, err := db.Query("select `id`, `ip`, `port` from `proxies` where `working` = 'U' or `working` = 'Y' order by `dt_last_used`, `dt_create` limit 100")
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				var prx Proxy
				prx.SetDB(db)
				err = rows.Scan(&prx.id, &prx.ip, &prx.port)
				if err != nil {
					log.Println(err)
				}
				prxs <- &prx
				prx.UpdateLastUsedTime()
			}
		}
	}()
	return &prxs
}

type Proxies chan *Proxy

type Proxy struct {
	id   uint
	ip   string
	port uint16
	db   *sql.DB
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

func (p *Proxy) SetDB(db *sql.DB) {
	p.db = db
}

func (p *Proxy) Log(addr string, statusCode int, duration float64) {
	parseUrl, err := url.Parse(addr)
	if err != nil {
		log.Println(err)
	}
	_, err = p.db.Exec("insert into `proxies_logs` (`id_proxy`,`url`,`domain`,`http_code`,`duration`) values(?, ?, ?, ?, ?)", p.id, parseUrl, parseUrl.Hostname(), statusCode, duration)
	if err != nil {
		log.Println(err)
	}
}

func (p *Proxy) UpdateLastUsedTime() {
	_, err := p.db.Exec("update `proxies` set `dt_last_used` = CURRENT_TIMESTAMP where `id` = ?", p.id)
	if err != nil {
		log.Println(err)
	}
}
