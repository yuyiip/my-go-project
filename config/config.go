package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// Config type
type Config struct {
	Name string
}

// Init the config file
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// init config file
	if err := c.initConfig(); err != nil {
		return err
	}

	// init log
	c.initLog()

	// watch the config file and hot reload
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // if there is config file, resolve it
	} else {
		viper.AddConfigPath("conf") // if no file, resolve default file
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") // set the type
	viper.AutomaticEnv()        // read the env variables
	viper.SetEnvPrefix("MY_GO_PROJECT")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // resolve config file
		return err
	}
	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	log.InitWithConfig(&passLagerCfg)
}

// watch config file changes and hot reload
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
