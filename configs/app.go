package configs

import (
	"flag"
	"github.com/spf13/viper"
	"godesktop/utils/log"
)

// Configuration .
type Configuration interface {
	Get(string) (string, error)
	Init(set *flag.FlagSet)
}

// ViperConfiguration .
type ViperConfiguration struct {
}

// Init .
func (vc *ViperConfiguration) Init() {

	// config paths
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.SetConfigName("config_app")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired")
		} else {
			log.Fatalf("Config file error %s", err.Error())
		}
	}

}

// CheckExists .
func (vc *ViperConfiguration) CheckExists(param string) bool {
	return viper.IsSet(param)
}

// Get .
func (vc *ViperConfiguration) Get(param string) string {
	return viper.GetString(param)
}

// GetInt .
func (vc *ViperConfiguration) GetInt(param string) int {
	return viper.GetInt(param)
}

// GetBool .
func (vc *ViperConfiguration) GetBool(param string) bool {
	return viper.GetBool(param)
}

// NewConfiguration .
func NewConfiguration() (cfg *ViperConfiguration) {
	cfg = &ViperConfiguration{}
	return

}

// Set .
func (vc *ViperConfiguration) Set(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}
