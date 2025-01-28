package cloudns

import (
	"github.com/caddyserver/caddy/v2"
	libdnsClouDns "github.com/libdns/cloudns"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdnsClouDns.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information for the DNS provider.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.cloudns",
		New: func() caddy.Module { return &Provider{new(libdnsClouDns.Provider)} },
	}
}

// Provision prepares the Provider by replacing placeholders in authentication fields with their actual values.
func (p *Provider) Provision(ctx caddy.Context) error {
	replacer := caddy.NewReplacer()
	p.Provider.AuthId = replacer.ReplaceAll(p.Provider.AuthId, "")
	p.Provider.SubAuthId = replacer.ReplaceAll(p.Provider.SubAuthId, "")
	p.Provider.AuthPassword = replacer.ReplaceAll(p.Provider.AuthPassword, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	cloudns {
//	    auth_id "<auth_id>"
//	    sub_auth_id "<sub_auth_id>"
//	    auth_password "<auth_password>"
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "auth_id":
				if d.NextArg() {
					p.Provider.AuthId = d.Val()
				} else {
					return d.ArgErr()
				}
			case "sub_auth_id":
				if d.NextArg() {
					p.Provider.SubAuthId = d.Val()
				} else {
					return d.ArgErr()
				}
			case "auth_password":
				if d.NextArg() {
					p.Provider.AuthPassword = d.Val()
				} else {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.AuthId == "" && p.Provider.SubAuthId == "" {
		return d.Err("missing auth id or sub auth id")
	}
	if p.Provider.AuthPassword == "" {
		return d.Err("missing auth password")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
