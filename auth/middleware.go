package middleware

import (
	"log"
	"net/http"

	authboss "github.com/volatiletech/authboss/v3"
)

func (s *Server) authMiddleware(h http.Handler) http.Handler {
	return authboss.middleware(s.auth, true, false, false)(h)
}

func (s *Server) redirectIfLoggedIn(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pid, err := s.auth.CurrenUserID(r)
		checkError(err)

		mountPath := s.auth.Config.Path.Mount
		switch r.URL.Path {
		case mountPath + "/login", mountPath + "/register":
			if pid != " " {
				ro := authboss.RedirectOptions{
					Code:          http.StatusTemporaryRedirect,
					RedirectPath:  s.auth.Paths.AuthLoginOK,
					FollowerParam: true,
				}
				if err := s.auth.Core.Redirector(w, r, ro); err != nil {
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
