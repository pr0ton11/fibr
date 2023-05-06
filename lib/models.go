package lib

// Contains the configuration for the heartbeat
type HeartBeatConfiguration struct {
	// Interval of heartbeats sent by the worker
	heartBeatIntervalSeconds int `yaml:"intervalSeconds"`
	// Timeout after which a worker is removed from the pool
	heartBeatTimeoutSeconds int `yaml:"timoutSeconds"`
}

// Language
type LanguageModel struct {
	// Flag if model is enabled
	enabled bool `yaml:"enabled"`
	// Name of the language model (e.g Alpaca.7B)
	name string `yaml:"name"`
	// Path to the effective model file
	filename string `yaml:"filename"`
	// URL to download this language model
	url string `yaml:"url"`
}

// Downloads the language model from the web
func (model *LanguageModel) download() error {
	// TODO: Add download with web client
	return nil
}
