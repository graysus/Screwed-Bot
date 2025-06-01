package common

import (
	"encoding/json"
	"log"
	"os"
)

type Mutation struct {
	Probability float64
	Filter      string
	Output      string
}

type Config struct {
	Gifs      []string
	Mutations []Mutation
}

var Conf Config

func init() {
	f, err := os.Open("assets/config/config.json")
	if err != nil {
		log.Fatalln("Failed to load config: " + err.Error())
	}

	defer f.Close()
	dec := json.NewDecoder(f)

	if err := dec.Decode(&Conf); err != nil {
		log.Fatalln("Failed to decode config: " + err.Error())
	}
}
