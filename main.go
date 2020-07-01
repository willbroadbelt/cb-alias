package main

import (
	"./cbalias"
	"io/ioutil"
	"log"
)


func main() {
	//Example usage
	if err := cbalias.GetConfigRepo(); err != nil {
		log.Printf("Get config failed: %v", err)
	}

	yamlFile, err := ioutil.ReadFile(cbalias.RepoPath + "/" + cbalias.ConfigName)
	if err != nil {
		log.Printf("Read file err   #%v ", err)
	}
	v, err := cbalias.ParseYaml(yamlFile)
	if err != nil {
		log.Printf("Parse yaml err   #%v ", err)
	}

	log.Printf("Got: %s\n", v["couchbase-server"]["7.0"].Stable)
	log.Printf("Got: %s\n", v["couchbase-server"]["7.0"].Release)

	log.Printf("Got: %s\n", v["couchbase-operator"]["1.2"].Release)

}
