package conf

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"

	"gopkg.in/yaml.v2"
)

// Configuration is the struct representing a configuration.
type Configuration struct {
	Service string `yaml:"service"`
}

// C is the exported global configuration variable
var C Configuration

// Load loads the given fp (file path) to the C global configuration variable.
func Load() error {
	var err error
	var hd string

	if hd, err = homedir.Dir(); err != nil {
		return err
	}

	cdir := hd + "/.config/"
	cf := cdir + "goploader.conf.yml"
	if _, err = os.Stat(cdir); os.IsNotExist(err) {
		log.Printf("Creating %v directory.\n", cdir)
		os.Mkdir(cdir, 0700)
	} else if err != nil {
		return err
	}
	if _, err = os.Stat(cf); os.IsNotExist(err) {
		log.Printf("Configuration file %v not found. Writing default configuration.\n", cf)
		C.Service = "https://up.depado.eu/"
		d, err := yaml.Marshal(C)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(cf, d, 0644)
	} else if err != nil {
		return err
	}

	conf, err := ioutil.ReadFile(cf)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(conf, &C)
}
