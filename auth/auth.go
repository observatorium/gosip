package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/Go-SIP/gosip/users"
	"github.com/abbot/go-http-auth"
)

type UserDatabase interface {
	GetByUsername(username string) (users.User, error)
	GetByToken(token string) (users.User, error)
}

type Handler struct {
	users              UserDatabase
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

func NewHandler(users UserDatabase) *Handler {
	ah := &Handler{users: users}

	ah.basicAuthenticator = auth.NewBasicAuthenticator("example.com", ah.usernamePassword)

	return ah
}

func (ah *Handler) usernamePassword(username, realm string) string {
	user, err := ah.users.GetByUsername(username)
	if err != nil {
		return ""
	}

	return user.Password()
}

func (ah *Handler) Basic(h http.Handler) http.Handler {
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

type username string

var authUsername username = "username"

func (ah *Handler) Token(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		splitHeader := strings.Split(tokenHeader, "Bearer")
		if len(splitHeader) != 2 {
			http.Error(w, "Bad request.", 400)
			return
		}
		token := strings.TrimSpace(splitHeader[1])

		user, err := ah.users.GetByToken(token)
		if err != nil {
			http.Error(w, "Unauthorized.", 401)
			return
		}

		ctx := context.WithValue(r.Context(), authUsername, user.Username())
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}

func Username(ctx context.Context) string {
	return ctx.Value(authUsername).(string)
}
