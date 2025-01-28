# ClouDNS Module for Caddy

This package provides a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It allows you to manage DNS records with ClouDNS.

## Caddy Module Name

```
dns.providers.cloudns
```

## Installation

To install this module, add it to your Caddy configuration.

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
tls {
  dns cloudns {
    auth_id "<auth_id>"
    sub_auth_id "<sub_auth_id>"
    auth_password "<auth_password>"
  }
}
```

## Environment Variables

You can also set the following environment variables to configure the module:

- `CLOUDNS_AUTH_ID`
- `CLOUDNS_SUB_AUTH_ID`
- `CLOUDNS_AUTH_PASSWORD`

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## Acknowledgements

This module is based on the [libdns](https://github.com/libdns/libdns) and [Caddy](https://github.com/caddyserver/caddy) projects.