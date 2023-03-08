package auth

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

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
			NickName: entry.GetAttributeValue(c.cfg.NickNameKey),
			Email:    entry.GetAttributeValue(c.cfg.EmailKey),
		})
	}

	return
}

// SearchUser 需判断user是否为空
func (c *LdapClient) SearchUser(name string) (user *Info, err error) {
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
		Filter: fmt.Sprintf(c.cfg.SearchStandard, name),
	}); err != nil {
		return
	}

	if len(result.Entries) > 0 {
		entry := result.Entries[0]
		user = &Info{
			Username: entry.GetAttributeValue(c.cfg.UsernameKey),
			Email:    entry.GetAttributeValue(c.cfg.EmailKey),
			NickName: entry.GetAttributeValue(c.cfg.NickNameKey),
		}
	}

	return
}
