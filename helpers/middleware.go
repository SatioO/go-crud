package helpers

import "net/http"

type ApiFunc func(w http.ResponseWriter, r *http.Request) error

func Register(h ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiError, ok := err.(ApiError); ok {
				WriteToJSON(w, apiError.StatusCode, apiError)
			} else {
				errResponse := ApiError{
					StatusCode: http.StatusInternalServerError,
					Msg: "internal server error",
				}

				WriteToJSON(w, http.StatusInternalServerError, errResponse)
			}
		}
	}
}
