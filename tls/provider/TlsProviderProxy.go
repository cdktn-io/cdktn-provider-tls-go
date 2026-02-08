// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type TlsProviderProxy struct {
	// When `true` the provider will discover the proxy configuration from environment variables.
	//
	// This is based upon [`http.ProxyFromEnvironment`](https://pkg.go.dev/net/http#ProxyFromEnvironment) and it supports the same environment variables (default: `true`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.2.1/docs#from_env TlsProvider#from_env}
	FromEnv interface{} `field:"optional" json:"fromEnv" yaml:"fromEnv"`
	// Password used for Basic authentication against the Proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.2.1/docs#password TlsProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// URL used to connect to the Proxy. Accepted schemes are: `http`, `https`, `socks5`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.2.1/docs#url TlsProvider#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Username (or Token) used for Basic authentication against the Proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tls/4.2.1/docs#username TlsProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

