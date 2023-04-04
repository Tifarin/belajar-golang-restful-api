package middleware

import (
	"belajar-golang-restful-api/model/web"
	"encoding/json"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}
func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}
func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(w,r)
	} else {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	webResponse := web.WebResponse{
		Code: http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}
	w.Header().Add("content-type","application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
	}
}