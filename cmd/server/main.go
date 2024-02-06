package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/sfomuseum/go-macaroon/sfomuseum"
	"github.com/superfly/macaroon"
)

func handler(key macaroon.SigningKey, loc string) http.Handler {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		kid := make([]byte, 0)

		m, err := macaroon.New(kid, loc, key)

		if err != nil {
			slog.Error("Failed to create macaroon", "error", err)
			http.Error(rsp, "Failed to create macaroon", http.StatusInternalServerError)
			return
		}

		// Max lifetime

		max_d := 10 * time.Second

		c := &macaroon.ValidityWindow{
			NotBefore: time.Now().Unix(),
			NotAfter:  time.Now().Add(max_d).Unix(),
		}

		err = m.Add(c)

		if err != nil {
			slog.Error("Failed to add caveats", "error", err)
			http.Error(rsp, "Failed to add caveats", http.StatusInternalServerError)
			return
		}

		enc, err := sfomuseum.EncodeMacaroonAsBase64(m)

		if err != nil {
			slog.Error("Failed to encode macaroon", "error", err)
			http.Error(rsp, "Failed to encode macaroon", http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-type", "text/plain")
		rsp.Write([]byte(enc))
	}

	return http.HandlerFunc(fn)
}

func main() {

	key := macaroon.NewSigningKey()
	loc := "localhost"

	mux := http.NewServeMux()
	mux.Handle("/", handler(key, loc))

	addr := "localhost:8080"
	slog.Info("Listening for requests", "address", addr)

	http.ListenAndServe(addr, mux)
}
