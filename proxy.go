package pwd

import (
	"database/sql"
	"log"
	"net/url"
)

// NewProxies returns a data feed with links to proxy servers from the database
func NewProxies(db *sql.DB, type_ string) *Proxies {
	prxs := make(Proxies)
	go func() {
		for {
			rows, err := db.Query("select `id`, `ip`, `port` from `proxies` where `type` = ? && `working` = 'U' or `working` = 'Y' order by `dt_last_used`, `dt_create` limit 100", type_)
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

// Proxies channel with proxy servers
type Proxies chan *Proxy

// Proxy data type containing information about the proxy server and managing it
type Proxy struct {
	id   uint
	ip   string
	port uint16
	db   *sql.DB
}

// SetID set proxy server id
func (p *Proxy) SetID(id uint) {
	p.id = id
}

// SetIP set the IP address of the proxy server
func (p *Proxy) SetIP(ip string) {
	p.ip = ip
}

// SetPort set proxy server port
func (p *Proxy) SetPort(port uint16) {
	p.port = port
}

// SetDB establish a link to the database connection that contains the proxy server entry
func (p *Proxy) SetDB(db *sql.DB) {
	p.db = db
}

// Log write to the log data about the work of the proxy server
func (p *Proxy) Log(link string, statusCode int, duration float64) {
	parseUrl, err := url.Parse(link)
	if err != nil {
		log.Println(err)
	}
	_, err = p.db.Exec("insert into `proxies_logs` (`id_proxy`,`url`,`domain`,`code`,`duration`) values(?, ?, ?, ?, ?)", p.id, parseUrl.String(), parseUrl.Hostname(), statusCode, duration)
	if err != nil {
		log.Println(err)
	}
}

// UpdateLastUsedTime updates the time when the proxy server was last used
func (p *Proxy) UpdateLastUsedTime() {
	_, err := p.db.Exec("update `proxies` set `dt_last_used` = CURRENT_TIMESTAMP where `id` = ?", p.id)
	if err != nil {
		log.Println(err)
	}
}
