package exception

import (
	"belajar-golang-restful-api/model/web"
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{})  {
	if notFoundError(w, r, err) {
		return
	}
	if validationErrors(w, r, err) {
		return
	}
	internalServerError(w, r, err)
}
func validationErrors(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	webResponse := web.WebResponse{
		Code: http.StatusBadRequest,
		Status: "BAD REQUEST",
		Data: exception.Error(),
	}
	w.Header().Add("content-type","application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	return true
	} else {
		return false
	}
}
func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	webResponse := web.WebResponse{
		Code: http.StatusNotFound,
		Status: "NOT FOUND",
		Data: exception.Error,
	}
	w.Header().Add("content-type","application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	return true
	} else {
		return false
	}
}
func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code: http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data: err,
	}
	w.Header().Add("content-type","application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
}