package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func DoubleHandler(w http.ResponseWriter, r *http.Request) {
	str := r.FormValue("val")
	if str == "" {
		http.Error(w, "value is missing", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(str)
	if err != nil {
		http.Error(w, "invalid number", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, num*2)
}
