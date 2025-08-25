package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func MustLoad(val interface{}, p string) {
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, val)
	if err != nil {
		panic(err)
	}
}
