package auth

import (
	"log"
	"net/http"

	authboss "github.com/volatiletech/authboss/v3"
)

func (s *Server) authMiddleware(h http.Handler) http.Handler {
	return authboss.Middleware(s.auth, true, false, false)(h)
}

func (s *Server) redirectIfLoggedIn(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pid, err := s.auth.CurrentUserID(r)
		checkError(err)

		mountPath := s.auth.Config.Paths.Mount
		switch r.URL.Path {
		case mountPath + "/login", mountPath + "/register":
			if pid != " " {
				ro := authboss.RedirectOptions{
					Code:             http.StatusTemporaryRedirect,
					RedirectPath:     s.auth.Paths.AuthLoginOK,
					FollowRedirParam: true,
				}
				if err := s.auth.Core.Redirector.Redirect(w, r, ro); err != nil {
					checkError(err)
				}
			}
			h.ServeHTTP(w, r)
		}
	})
}

func checkError(err error) {
	if err != nil {
		log.Print(err)
	}
}
