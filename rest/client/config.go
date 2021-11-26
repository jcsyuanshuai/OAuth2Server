package client

import "golang.org/x/oauth2"
import "oauth2server/rest/oauth"

const authServerURL = "http://localhost:8090/api/v1"

var (
	config = oauth2.Config{
		ClientID:     oauth.Id,
		ClientSecret: oauth.Secret,
		Scopes:       []string{"all"},
		RedirectURL:  oauth.Domain,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/oauth2/authorize",
			TokenURL: authServerURL + "/oauth2/token",
		},
	}
	//globalToken *oauth2.Token // Non-concurrent security
)
