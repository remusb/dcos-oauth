package main

import "github.com/codegangsta/cli"

var (
	// TODO can we proxy this?
	flIssuerURL = cli.StringFlag{
		Name:   "issuer-url",
		Usage:  "JWT Issuer URL",
		Value:  "https://dcos.auth0.com/",
		EnvVar: "OAUTH_ISSUER_URL",
	}

	flClientID = cli.StringFlag{
		Name:   "client-id",
		Usage:  "JWT Client ID",
		Value:  "3yF5TOSzdlI45Q1xspxzeoGBe9fNxm9m",
		EnvVar: "OAUTH_CLIENT_ID",
	}

	flSecretKeyPath = cli.StringFlag{
		Name:   "secret-key-path",
		Usage:  "Secret key file path",
		Value:  "/var/lib/dcos/auth-token-secret",
		EnvVar: "SECRET_KEY_FILE_PATH",
	}

	flSegmentKey = cli.StringFlag{
		Name:  "segment-key",
		Usage: "Segment key",
		Value: "39uhSEOoRHMw6cMR6st9tYXDbAL3JSaP",
	}

	flAllowLocalUsers = cli.StringFlag{
		Name:  keyAllowLocalUsers,
		Usage: "Allow local users",
		Value: "false",
		EnvVar: "OAUTH_ALLOW_LOCAL_USERS",
	}

	flDefaultLocalUser = cli.StringFlag{
		Name:  keyDefaultLocalUser,
		Usage: "Default local user",
		Value: "",
		EnvVar: "OAUTH_DEFAULT_LOCAL_USER",
	}

	flDefaultLocalUserHash = cli.StringFlag{
		Name:  keyDefaultLocalUserHash,
		Usage: "Initial password hash for the default local user",
		Value: "",
		EnvVar: "OAUTH_DEFAULT_LOCAL_USER_HASH",
	}

	flAllowLdapUsers = cli.StringFlag{
		Name:  keyAllowLdapUsers,
		Usage: "Allow LDAP users",
		Value: "false",
		EnvVar: "OAUTH_ALLOW_LDAP_USERS",
	}

	flLdapConfigFile = cli.StringFlag{
		Name:  keyLdapConfigFile,
		Usage: "LDAP config file",
		Value: "/etc/ethos/ldap.toml",
		EnvVar: "OAUTH_LDAP_CONFIG_FILE",
	}

	flLdapWhitelistOnly = cli.StringFlag{
		Name:  keyLdapWhitelistOnly,
		Usage: "LDAP user allowed only when on whitelist",
		Value: "false",
		EnvVar: "OAUTH_LDAP_WHITELIST_ONLY",
	}

	flLdapGroupsOnly = cli.StringFlag{
		Name:  keyLdapGroupsOnly,
		Usage: "LDAP user allowed only when in LDAP group with admin role",
		Value: "false",
		EnvVar: "OAUTH_LDAP_GROUPS_ONLY",
	}

	flLdapCheckOnOauth = cli.StringFlag{
		Name:  keyLdapCheckOnOauth,
		Usage: "Perform LDAP group check during Oauth login",
		Value: "false",
		EnvVar: "OAUTH_LDAP_CHECK_ON_OAUTH",
	}

	flOauthAdminGroupsFile = cli.StringFlag{
		Name: keyAdminGroupsFile,
		Usage: "CSV file enumerating Oauth groups with admin role",
		Value: "",
		EnvVar: "OAUTH_ADMIN_GROUPS_FILE",
	}
)
