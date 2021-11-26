package oauth

import (
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/go-session/session"
	"github.com/golang-jwt/jwt"
	"net/http"
)

const (
	Id     = "1024"
	Secret = "0403"
	Domain = "http://localhost:8080/client/oauth2"
)

var srv *server.Server

var clientStore *store.ClientStore

func init() {
	clientStore, _ = NewClientStore()
	srv, _ = NewServer()
}

func GetServer() *server.Server {
	return srv
}

func GetClientStore() *store.ClientStore {
	return clientStore
}

func NewClientStore() (cs *store.ClientStore, err error) {
	cs = store.NewClientStore()
	_ = cs.Set(Id, &models.Client{
		ID:     Id,
		Secret: Secret,
		Domain: Domain,
	})
	print(cs)
	return cs, nil
}

func NewServer() (srv *server.Server, err error) {
	mgr := manage.NewDefaultManager()
	mgr.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	mgr.MustTokenStorage(store.NewMemoryTokenStore())
	mgr.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodES256))
	mgr.MapClientStorage(GetClientStore())

	cfg := server.NewConfig()
	srv = server.NewServer(cfg, mgr)

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})

	srv.SetUserAuthorizationHandler(func(wtr http.ResponseWriter, req *http.Request) (userID string, err error) {
		store, err := session.Start(req.Context(), wtr, req)
		if err != nil {
			return
		}

		uid, ok := store.Get("LoggedInUserID")
		if !ok {
			if req.Form == nil {
				req.ParseForm()
			}

			store.Set("ReturnUri", req.Form)
			store.Save()

			wtr.Header().Set("Location", "/server/login")
			wtr.WriteHeader(http.StatusFound)

			return
		}

		userID = uid.(string)
		store.Delete("LoggedInUserID")
		store.Save()
		return
	})

	//srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
	//	return
	//})
	//
	//srv.SetResponseErrorHandler(func(re *errors.Response) {
	//	return
	//})
	return srv, nil
}
