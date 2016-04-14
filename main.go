package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/hkchindeko/Beat-MQTT/beater"
)

func main() {
	err := beat.Run("beat-mqtt", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
