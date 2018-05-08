package gowu

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"runtime"
)

var fixturesPath = ""

func init() {
	_, filename, _, _ := runtime.Caller(1)
	fixturesPath = path.Join(path.Dir(filename), "./testdata/fixtures")
}

func GetFuxture(filename string) string {
	return path.Join(fixturesPath, filename)
}

func LoadFixtureFile(file string) []byte {
	if data, err := ioutil.ReadFile(GetFuxture(file)); err == nil {
		return data
	} else {
		log.Error(err)
	}

	var b []byte
	return b
}
