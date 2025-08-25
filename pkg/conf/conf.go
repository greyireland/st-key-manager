package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//必须打yaml标签才能解析

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
