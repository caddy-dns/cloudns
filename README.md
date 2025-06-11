# ClouDNS Module for Caddy

This package provides a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It allows you to manage DNS records with [ClouDNS](https://www.cloudns.net/) for automated TLS certificate issuance via the ACME DNS challenge.

## Caddy Module Name

```
dns.providers.cloudns
```

## Config Parameters

- `auth_id` - ClouDNS Auth ID (required if `sub_auth_id` is not provided)
- `sub_auth_id` - ClouDNS Sub Auth ID (required if `auth_id` is not provided)
- `auth_password` - ClouDNS Auth Password (required)

## Installation

To use this module, you need to build Caddy with this module included. The simplest way is to use [xcaddy](https://github.com/caddyserver/xcaddy):

```bash
xcaddy build --with github.com/caddy-dns/cloudns
```

## Configuration Examples

### JSON Configuration

To use this module for the ACME DNS challenge, configure the ACME issuer in your Caddy JSON like so:

```json
{
  "apps": {
    "tls": {
      "automation": {
        "policies": [
          {
            "issuers": [
              {
                "module": "acme",
                "challenges": {
                  "dns": {
                    "provider": {
                      "name": "cloudns",
                      "auth_id": "CLOUDNS_AUTH_ID",
                      "sub_auth_id": "CLOUDNS_SUB_AUTH_ID",
                      "auth_password": "CLOUDNS_AUTH_PASSWORD"
                    }
                  }
                }
              }
            ]
          }
        ]
      }
    }
  }
}
```

### Caddyfile Configuration

You can also configure the module using the Caddyfile.

#### Global Configuration

```
{
  acme_dns cloudns {
    auth_id "<auth_id>"
    sub_auth_id "<sub_auth_id>"
    auth_password "<auth_password>"
  }
}
```

#### Per-Site Configuration

```
example.com {
  tls {
    dns cloudns {
      auth_id "<auth_id>"
      sub_auth_id "<sub_auth_id>"
      auth_password "<auth_password>"
    }
  }
}
```

## Environment Variables

You can also set the following environment variables to configure the module:

- `CLOUDNS_AUTH_ID` - ClouDNS Auth ID
- `CLOUDNS_SUB_AUTH_ID` - ClouDNS Sub Auth ID
- `CLOUDNS_AUTH_PASSWORD` - ClouDNS Auth Password

Note: Configuration via environment variables will be used as a fallback if the corresponding parameters are not specified in the Caddy configuration.

## Advanced Configuration

The module includes built-in retry mechanisms with exponential backoff for all DNS operations:

- Default operation retries: 5
- Default initial backoff: 1 second
- Default maximum backoff: 30 seconds

These values are automatically configured and do not need to be specified in your configuration.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgements

This module is based on the [libdns/cloudns](https://github.com/libdns/cloudns) package and the [Caddy](https://github.com/caddyserver/caddy) project.