package auth

import (
	"crypto/tls"

	"github.com/go-ldap/ldap/v3"
)

type conn interface {
	Bind(username, password string) error
	Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error)
	StartTLS(config *tls.Config) error
	UnauthenticatedBind(username string) error
	Close()
}

func dial(cfg *LdapAuth) (conn, error) {
	var opts []ldap.DialOpt
	return ldap.DialURL(cfg.ServerURL, opts...)
}

type LdapClient struct {
	dial func(cfg *LdapAuth) (conn, error)
	cfg  *LdapAuth
}

func NewLdapClient(options ...ClientOptionFunc) (*LdapClient, error) {
	client := &LdapClient{
		dial: dial,
	}
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(client); err != nil {
			return nil, err
		}
	}
	return client, nil
}
