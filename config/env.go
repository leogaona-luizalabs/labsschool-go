package config

import (
	"os"
	"strconv"
)

// BurzumToken ...
var BurzumToken string

// Port ...
var Port int

func init() {
	BurzumToken = os.Getenv("BURZUM_TOKEN")
	Port, _ = strconv.Atoi(os.Getenv("PORT"))
}
