package auth

import (
	"os"
	"testing"
)

func setup(t *testing.T) (*LdapClient, error) {
	cfg := &LdapAuth{
		ServerURL:      os.Getenv("ServerURL"),
		BindDN:         os.Getenv("BindDN"),
		BindPassword:   os.Getenv("BindPassword"),
		SearchBaseDN:   os.Getenv("SearchBaseDN"),
		SearchStandard: "(&(objectClass=user)(userPrincipalName=%s))",
		EmailSuffix:    os.Getenv("EmailSuffix"),
		EmailKey:       "mail",
		UsernameKey:    "sAMAccountName",
		NickNameKey:    "displayName",
	}
	return NewLdapClient(WithConf(cfg))
}
