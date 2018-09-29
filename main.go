package main

import (
	"context"
	"encoding/json"
	"flag"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tukejonny/go-oauth2-example/internal/auth"
	"github.com/tukejonny/go-oauth2-example/internal/config"
)

var (
	confPath = flag.String("config", "./conf/app.yml", "Path for config file")

	authClient *auth.GoogleAuth
	appConf    *config.AppConfig
)

func handleGoogleAuthorize(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Redirect(w, r, authClient.Config().AuthCodeURL(appConf.Nonce), http.StatusFound)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.URL.Query().Get("nonce") != appConf.Nonce {
		http.Error(w, "Invalid nonce value", http.StatusBadRequest)
		return
	}

	userInfo, err := authClient.FetchUserInfo(r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.MarshalIndent(userInfo, "", "   ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func main() {
	flag.Parse()
	var (
		ctx     = context.Background()
		initErr error
	)
	appConf, initErr = config.LoadConfig(*confPath)
	if initErr != nil {
		panic(initErr)
	}

	authClient, initErr = auth.NewGoogleAuth(ctx, &appConf.AuthConf)
	if initErr != nil {
		panic(initErr)
	}

	router := httprouter.New()
	router.GET("/auth/google", handleGoogleAuthorize)
	router.GET("/auth/google/callback", handleGoogleCallback)

	if err := http.ListenAndServe(":8000", router); err != nil {
		panic(err)
	}
}
