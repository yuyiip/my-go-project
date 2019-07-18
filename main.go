package main

import (
	"errors"
	"net/http"
	"time"

	"my-go-project/config"
	"my-go-project/router"
	"my-go-project/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg     = pflag.StringP("config", "c", "", "my-go-project config file path.")
	// version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	// if *version {
	// 	v := v.Get()
	// 	marshalled, err := json.MarshalIndent(&v, "", "  ")
	// 	if err != nil {
	// 		log.Error("%v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println(string(marshalled))
	// 	return
	// }

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// int db
	// model.DB.Init()
	// defer model.DB.Close()

	// Set gin mode
	gin.SetMode(viper.GetString("runmode"))
	// create the Gin engine
	g := gin.New()

	// gin middlewares
	// middlewares := []gin.HandlerFunc{}

	// routes
	router.Load(
		g,
		// middlewares
		middleware.Logging(),
	)

	// Ping the server to make sure the router is working
	go func() {
		if err := pingServer(); err != nil {
			log.Error("The router has no response, or it might took too long to start up", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening to the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
