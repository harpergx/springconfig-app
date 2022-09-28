package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harpergx/springconfig-app-in-go/config"
	"github.com/harpergx/springconfig-app-in-go/controller"
	"github.com/spf13/viper"
)

var appName = "app"

func init() {
	profile := flag.String("profile", "dev", "Environment profile, something similar to spring profiles")
	configServerUrl := flag.String("configServerUrl", "http://localhost:8888", "Address to config server")
	configBranch := flag.String("configBranch", "master", "git branch to fetch configuration from")
	flag.Parse()

	// Pass the flag values into viper.
	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)
}

func main() {
	config.LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))

	r := mux.NewRouter()
	err := controller.Controller(r)

	if err != nil {
		panic("서버 실행에 실패했습니다.")
	}

	http.ListenAndServe(":8080", r)

}
