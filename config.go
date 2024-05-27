package auth

type LdapAuth struct {
	ServerURL      string `mapstructure:"server-url" json:"serverURL" yaml:"server-url"`                  // ldap服务器: 'ldap://xxx.xxx.xxx.xxx:389'
	BindDN         string `mapstructure:"bind-dn" json:"bindDN" yaml:"bind-dn"`                           // ldap通用账号dn: 'cn=xxx,ou=Account_Users,dc=int,dc=xxx,dc=com'
	BindPassword   string `mapstructure:"bind-password" json:"bindPassword" yaml:"bind-password"`         // ldap通用账号密码: 'xxxxxx'
	SearchBaseDN   string `mapstructure:"search-base-dn" json:"searchBaseDN" yaml:"search-base-dn"`       // ldap搜索开始DN: 'ou=Account_Users,dc=int,dc=xxx,dc=com'
	SearchStandard string `mapstructure:"search-standard" json:"searchStandard" yaml:"search-standard"`   // ldap搜索过滤: '&(objectClass=user)(sAMAccountName=%s)'
	EmailSuffix    string `mapstructure:"email-suffix" json:"emailSuffix" yaml:"email-suffix"`            // email后缀: '@xxx.com'
	EmailKey       string `mapstructure:"email-key" json:"emailKey" yaml:"email-key"`                     // ldap校验email属性: 'mail'
	UsernameKey    string `mapstructure:"username-key" json:"usernameKey" yaml:"username-key"`            // ldap鉴权用户名: 'sAMAccountName'
	DisplayNameKey string `mapstructure:"display-name-key" json:"displayNameKey" yaml:"display-name-key"` // ldap显示用户名: 'displayName'
}
