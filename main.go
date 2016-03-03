package main

import (
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	log.Info("Starting PihexUI")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")      // name of config file (without extension)
	viper.AddConfigPath("/etc/pihex/") // path to look for the config file in
	viper.AddConfigPath(".")           // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		log.Fatal("Failed to read config ", err)
	}

	level, _ := log.ParseLevel(strings.ToLower(viper.GetString("loggingLevel")))
	log.SetLevel(level)

	r := mux.NewRouter()

	webroot := viper.GetString("webroot")
	log.Debug("Web Content root " + webroot)
	fs := http.FileServer(http.Dir(webroot))
	//http.Handle("/", http.StripPrefix("/", fs))
	http.Handle("/", fs)

	http.Handle("/api", r)
	http.ListenAndServe(viper.GetString("listenAddress"), nil)
}
