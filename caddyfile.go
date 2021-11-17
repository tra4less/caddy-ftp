package ftp

import (
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("ftp", parseCaddyfile)
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	h4f := new(HTTP4Ftp)
	err := h4f.UnmarshalCaddyfile(h.Dispenser)
	return h4f, err
}

// UnmarshalCaddyfile sets up the handler from Caddyfile tokens. Syntax:
// Specifying the formats on the first line will use those formats' defaults.
func (h4f *HTTP4Ftp) UnmarshalCaddyfile(d *caddyfile.Dispenser) (err error) {
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "addr":
				if !d.NextArg() {
					return d.ArgErr()
				}
				h4f.Addr = d.Val()
			case "pass":
				if !d.NextArg() {
					return d.ArgErr()
				}
				h4f.Pass = d.Val()
			case "user":
				if !d.NextArg() {
					return d.ArgErr()
				}
				h4f.User = d.Val()
			case "dial_timeout":
				if !d.NextArg() {
					return d.ArgErr()
				}

				if h4f.DialTimeout, err = time.ParseDuration(d.Val()); err != nil {
					return err
				}
			case "disable_epsv":
				h4f.DisabledEPSV = true

			case "disable_mlsd":
				h4f.DisabledMLSD = true

			case "disable_utf8":
				h4f.DisableUTF8 = true
			}
		}
	}
	return nil
}

// Interface guard
var (
	_ caddyfile.Unmarshaler       = (*HTTP4Ftp)(nil)
	_ caddyhttp.MiddlewareHandler = (*HTTP4Ftp)(nil)
	_ caddy.Validator             = (*HTTP4Ftp)(nil)
)
