package auth

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/pkg/errors"
)

// ErrEntries is returned by ldap Authenticate function,
// When search result return user DN does not exist or too many entries returned.
var ErrEntries = errors.New("auth/ldap: Search user DN does not exist or too many entries returned")

type Info struct {
	Username string
	Email    string
	NikName  string
}

func (c *LdapClient) Authenticate(userName, password string) (*Info, error) { //nolint:lll
	l, err := c.dial(c.cfg)

	if err != nil {
		return nil, err
	}

	defer l.Close()

	if c.cfg.BindPassword != "" {
		err = l.Bind(c.cfg.BindDN, c.cfg.BindPassword)
	} else {
		err = l.UnauthenticatedBind(c.cfg.BindDN)
	}

	if err != nil {
		return nil, err
	}

	result, err := l.Search(&ldap.SearchRequest{
		BaseDN: c.cfg.SearchBaseDN,
		Scope:  ldap.ScopeWholeSubtree,
		Filter: fmt.Sprintf(c.cfg.SearchStandard, userName),
	})

	if err != nil {
		return nil, err
	}

	if len(result.Entries) != 1 {
		return nil, ErrEntries
	}

	err = l.Bind(result.Entries[0].DN, password)

	if err != nil {
		return nil, err
	}

	name := result.Entries[0].GetAttributeValue(c.cfg.UsernameKey)
	mail := result.Entries[0].GetAttributeValue(c.cfg.EmailKey)
	nikName := result.Entries[0].GetAttributeValue(c.cfg.NikNameKey)

	return &Info{Username: name, Email: mail, NikName: nikName}, nil
}

func (c *LdapClient) GetUserList() (users []*Info, err error) {
	var (
		l      conn
		result *ldap.SearchResult
	)
	if l, err = c.dial(c.cfg); err != nil {
		return
	}
	defer l.Close()

	if c.cfg.BindPassword != "" {
		err = l.Bind(c.cfg.BindDN, c.cfg.BindPassword)
	} else {
		err = l.UnauthenticatedBind(c.cfg.BindDN)
	}

	if err != nil {
		return nil, err
	}

	if result, err = l.Search(&ldap.SearchRequest{
		BaseDN: c.cfg.SearchBaseDN,
		Scope:  ldap.ScopeWholeSubtree,
		Filter: fmt.Sprintf(c.cfg.SearchStandard, "*"),
	}); err != nil {
		return
	}

	users = make([]*Info, 0, len(result.Entries))
	for _, entry := range result.Entries {
		users = append(users, &Info{
			Username: entry.GetAttributeValue(c.cfg.UsernameKey),
			NikName:  entry.GetAttributeValue(c.cfg.NikNameKey),
			Email:    entry.GetAttributeValue(c.cfg.EmailKey),
		})
	}

	return
}
