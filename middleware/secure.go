package middleware

import (
	"github.com/iris-contrib/middleware/secure"
)

// GetSecurity adds to the response some headers to increase protection.
func GetSecurity() *secure.Secure {
	return secure.New(secure.Options{
		STSSeconds:              315360000,
		STSIncludeSubdomains:    true,
		STSPreload:              true,
		ForceSTSHeader:          false,
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXSSFilter:        true,
	})
}
