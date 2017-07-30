package main

import (
	"fmt"
	"log"
  "github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile("./conf/init.json")
  //alternatives
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("$HOME/conf")
	viper.SetConfigName("init")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	values := viper.GetStringMapString("clusters")
	for key, val := range values {
		fmt.Printf("Key: %s, Value: %s", key, val)
	}

	/** Unmarshal into struct
	type Cluster struct {
		ID                    string
		Name                  string
		NodeInstanceType      string
		NodeInstanceSpotPrice string
		NodeMin               int
		NodeMax               int
	}
	
  var C Cluster

	err := prod.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Println(C)
  **/
}
