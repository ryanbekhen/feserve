package proxy

import (
	pathutil "path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/ryanbekhen/feserve/internal/balancer"
	"github.com/ryanbekhen/feserve/internal/httputils"
)

type Proxy struct {
	proxypaths   []proxypath
	proxydomains []proxydomain
}

type proxypath struct {
	path    string
	rewrite bool
	servers *balancer.RoundRobin
}

type proxydomain struct {
	domain  string
	servers *balancer.RoundRobin
}

func New() *Proxy {
	return &Proxy{}
}

func (p *Proxy) AddForwardToPath(path string, rewrite bool, servers []string) {
	p.proxypaths = append(p.proxypaths, proxypath{
		path, rewrite, balancer.NewRoundRobin(servers),
	})
}

func (p *Proxy) AddForwardToDomain(domain string, servers []string) {
	p.proxydomains = append(p.proxydomains, proxydomain{
		domain, balancer.NewRoundRobin(servers),
	})
}

func (p *Proxy) Routing(r *fiber.App) {
	if len(p.proxydomains) > 0 {
		r.Use(func(c *fiber.Ctx) error {
			host := string(c.Request().Host())
			for _, proxydomain := range p.proxydomains {
				if proxydomain.domain == host {
					server := proxydomain.servers.Get()
					requestPath := c.Path()
					query := string(c.Request().URI().QueryString())
					if query != "" {
						query = "?" + query
					}
					httputils.ForwardUserIP(c)
					return proxy.Do(c, server+requestPath+query)
				}
			}

			return c.Next()
		})
	}

	for _, proxypath := range p.proxypaths {
		if proxypath.path == "*" {
			panic("balancer using path cannot use path *")
		}
		r.All(pathutil.Join(proxypath.path, "*"), func(c *fiber.Ctx) error {
			server := proxypath.servers.Get()
			var requestPath string
			if proxypath.rewrite {
				requestPath = strings.Replace(c.Path(), proxypath.path, "", 1)
			} else {
				requestPath = c.Path()
			}
			query := string(c.Request().URI().QueryString())
			if query != "" {
				query = "?" + query
			}
			httputils.ForwardUserIP(c)
			err := proxy.Do(c, server+requestPath+query)
			if err != nil {
				return err
			}
			c.Response().Header.Del(fiber.HeaderServer)
			return nil
		})
	}
}
