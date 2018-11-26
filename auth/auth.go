package auth

import (
	"net/http"
	"strings"

	auth "github.com/abbot/go-http-auth"
	"github.com/Go-SIP/gosip/users"
)

type AuthHandler interface {
	BasicAuth(h http.Handler) http.Handler
	TokenAuth(h http.Handler) http.Handler
}

type authHandler struct {
	db                 users.UserDatabase
	basicAuthenticator auth.AuthenticatorInterface
	tokenAuthenticator auth.AuthenticatorInterface
}

func authenticatedHandler(a auth.AuthenticatorInterface, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := a.NewContext(r.Context(), r)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func NewAuthHandler(db users.UserDatabase) AuthHandler {
	ah := &authHandler{db: db}
	basicAuthenticator := auth.NewBasicAuthenticator("example.com", ah.usernamePassword)
	ah.basicAuthenticator = basicAuthenticator
	return ah
}

func (ah *authHandler) usernamePassword(username, realm string) string {
	user, err := ah.db.GetByUsername(username)
	if err != nil {
		return ""
	}

	return user.Password()
}

func (ah *authHandler) BasicAuth(h http.Handler) http.Handler {
	return authenticatedHandler(
		ah.basicAuthenticator,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			authInfo := auth.FromContext(ctx)
			authInfo.UpdateHeaders(w.Header())
			if authInfo == nil || !authInfo.Authenticated {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			h.ServeHTTP(w, r)
		}))
}

func (ah *authHandler) TokenAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		splitHeader := strings.Split(tokenHeader, "Bearer")
		if len(splitHeader) != 2 {
			http.Error(w, "Bad request.", 400)
			return
		}
		token := splitHeader[1]

		_, err := ah.db.GetByToken(token)
		if err != nil {
			http.Error(w, "Unauthorized.", 401)
			return
		}

		h.ServeHTTP(w, r)
	})
}
