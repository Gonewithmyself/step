package auth

import "net/http"

func router() {
	http.HandleFunc("/sign", sign)
	http.HandleFunc("/refresh", sign)
}

func sign(w http.ResponseWriter, r *http.Request) {

}

func refresh(w http.ResponseWriter, r *http.Request) {

}
