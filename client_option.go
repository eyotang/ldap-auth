package auth

type ClientOptionFunc func(*LdapClient) error

func WithConf(conf *LdapAuth) ClientOptionFunc {
	return func(client *LdapClient) error {
		client.cfg = conf
		return nil
	}
}

func WithConfProps(serverURL, bindDN, bindPasswd, searchBaseDN, searchStandard, emailSuffix, emailKey, usernameKey, displayNameKey string) ClientOptionFunc {
	return func(client *LdapClient) error {
		client.cfg = &LdapAuth{
			ServerURL:      serverURL,
			BindDN:         bindDN,
			BindPassword:   bindPasswd,
			SearchBaseDN:   searchBaseDN,
			SearchStandard: searchStandard,
			EmailSuffix:    emailSuffix,
			EmailKey:       emailKey,
			UsernameKey:    usernameKey,
			DisplayNameKey: displayNameKey,
		}
		return nil
	}
}
