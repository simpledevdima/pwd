package pwd

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"os"
	"time"
)

func NewSelenium() *Selenium {
	s := &Selenium{}
	s.paths.selenium = os.Getenv("GOPATH") + "/src/github.com/tebeka/selenium/vendor/selenium-server.jar"
	s.paths.chromeDriver = os.Getenv("GOPATH") + "/src/github.com/tebeka/selenium/vendor/chromedriver"
	s.port = 4444
	s.timeout = 60
	s.debug = false
	return s
}

type Selenium struct {
	paths struct {
		selenium     string
		chromeDriver string
	}
	port      uint16
	timeout   time.Duration
	debug     bool
	service   *selenium.Service
	caps      selenium.Capabilities
	WebDriver selenium.WebDriver
}

func (s *Selenium) Start() {
	fmt.Printf("start selenium service with chrome driver\n")
	options := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),
		selenium.Output(os.Stderr),
		selenium.ChromeDriver(s.paths.chromeDriver),
	}
	selenium.SetDebug(s.debug)
	service, err := selenium.NewSeleniumService(s.paths.selenium, int(s.port), options...)
	if err != nil {
		log.Println(err)
	} else {
		s.service = service
		s.caps = selenium.Capabilities{
			"browserName": "chrome",
		}
	}
}

func (s *Selenium) Stop() {
	fmt.Printf("stop selenium service\n")
	err := s.service.Stop()
	if err != nil {
		log.Println(err)
	}
}

func (s *Selenium) SetProxy(ip string, port uint16) {
	s.caps.AddProxy(selenium.Proxy{
		Type: selenium.Manual,
		SSL:  fmt.Sprintf("%s:%d", ip, port),
	})
}

func (s *Selenium) Connect() {
	wd, err := selenium.NewRemote(s.caps, fmt.Sprintf("http://localhost:%d/wd/hub", s.port))
	if err != nil {
		log.Println(err)
	}
	s.WebDriver = wd
	s.setup()
}

func (s *Selenium) setup() {
	err := s.WebDriver.SetPageLoadTimeout(s.timeout * time.Second)
	if err != nil {
		log.Println(err)
	}
}

func (s *Selenium) Disconnect() {
	err := s.WebDriver.Quit()
	if err != nil {
		log.Println(err)
	}
}
