package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/digineo/go-uci"
	"github.com/pkg/errors"
)

const (
	upnpConfigName  = "upnpd"
	upnpSectionName = "upnpd"
)

func main() {
	u := uci.NewTree("/etc/config")
	err := u.LoadConfig(upnpConfigName, false)
	if err != nil {
		log.Fatal(err)
	}

	ifaces, err := getSections(u, upnpConfigName, upnpSectionName)
	if err != nil {
		log.Fatal(err)
	}

	enabled, ok := u.GetBool(upnpConfigName, ifaces[0], "enabled")
	if !ok {
		log.Fatal(err)
	}

	c := Upnp{
		Enabled: enabled,
	}

	cc, _ := json.Marshal(c)
	fmt.Println(string(cc))
}

func getSections(u uci.Tree, configName string, section string) ([]string, error) {
	ifaces, ifOK := u.GetSections(configName, section)
	if !ifOK {
		return nil, errors.Errorf("No %s in %s found", section, configName)
	}

	return ifaces, nil
}

type Upnp struct {
	Enabled bool `json:"enabled"`
}
