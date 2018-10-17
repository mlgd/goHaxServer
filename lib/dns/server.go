package dns

import (
	"net"
	"strings"
	"sync"

	"github.com/astaxie/beego"
	libDNS "github.com/miekg/dns"
)

const (
	DefaultExchange = "8.8.4.4:53"
)

var (
	server          *libDNS.Server
	filterEnabled   bool
	filterAddresses []string
	mu              sync.Mutex
	exchange        string
)

func StartServer(filter bool) {
	filterEnabled = filter
	if exchange == "" {
		exchange = DefaultExchange
	}

	libDNS.HandleFunc(".", handleReflect)

	go serve()
}

func StopServer() {
	if server != nil {
		server.Shutdown()
	}
}

func SetExchange(e string) {
	exchange = e
}

func SetFilterAddresses(addresses []string) {
	mu.Lock()
	defer mu.Unlock()
	filterAddresses = make([]string, 0)
	for _, address := range addresses {
		if address != "" {
			if !strings.HasSuffix(address, ".") {
				address += "."
			}
			filterAddresses = append(filterAddresses, address)
		}
	}
}

func EnableFilter() {
	filterEnabled = true
}

func DisableFilter() {
	filterEnabled = false
}

func serve() {
	server = &libDNS.Server{Addr: ":53", Net: "udp", TsigSecret: nil}
	if err := server.ListenAndServe(); err != nil {
		beego.Critical(err)
	}
}

func handleReflect(w libDNS.ResponseWriter, r *libDNS.Msg) {
	var (
		v4 bool
		a  net.IP
		rr libDNS.RR
		m  *libDNS.Msg
	)

	for _, question := range r.Question {
		filter := false

		if filterEnabled {
			switch question.Qtype {
			case libDNS.TypeA, libDNS.TypeAAAA:
				mu.Lock()
				for _, address := range filterAddresses {
					if question.Name == address {
						filter = true
						break
					}
				}
				mu.Unlock()
			}
		}

		if filter {
			m = new(libDNS.Msg)
			m.SetReply(r)
			m.Compress = false
			if ip, ok := w.RemoteAddr().(*net.UDPAddr); ok {
				a = ip.IP
				v4 = a.To4() != nil
			}
			if v4 {
				rr = &libDNS.A{
					Hdr: libDNS.RR_Header{Name: question.Name, Rrtype: libDNS.TypeA, Class: libDNS.ClassINET, Ttl: 0},
					A:   a.To4(),
				}
			} else {
				rr = &libDNS.AAAA{
					Hdr:  libDNS.RR_Header{Name: question.Name, Rrtype: libDNS.TypeAAAA, Class: libDNS.ClassINET, Ttl: 0},
					AAAA: a,
				}
			}
			m.Answer = append(m.Answer, rr)
		} else {
			m, _ = libDNS.Exchange(r, exchange)
		}
	}

	w.WriteMsg(m)
}
