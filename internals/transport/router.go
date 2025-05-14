package transport

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")

	if key != "" {
		fmt.Fprintln(w, "key: ", key)
	} else {
		fmt.Fprintln(w, "No Quary parm")
	}
}
