// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type Config struct {
	Beat-mqtt Beat-mqttConfig
}

type Beat-mqttConfig struct {
	Period string `yaml:"period"`
}
