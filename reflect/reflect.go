package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Name    string `json:"server-name"` // CONFIG_SERVER_NAME
	IP      string `json:"server-ip"`   // CONFIG_SERVER_IP
	URL     string `json:"server-url"`  // CONFIG_SERVER_URL
	Timeout string `json:"timeout"`     // CONFIG_TIMEOUT
}

func readConfig() *Config {
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exists := os.LookupEnv(key); exists {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}

func main() {
	/*
		os.Setenv("CONFIG_SERVER_IP", "127.0.0.1")
		os.Setenv("CONFIG_SERVER_URL", "http://127.0.0.1:8080")
		os.Setenv("CONFIG_TIMEOUT", "30")
		os.Setenv("CONFIG_SERVER_NAME", "server")
		c := readConfig()
		fmt.Printf("%+v", c)
	*/
	config := Config{
		IP:  "127.0.0.1",
		URL: "http://127.0.0.1:8080",
	}
	t := reflect.TypeOf(config)
	v := reflect.ValueOf(config)
	fmt.Printf("Object number elements: %d\n", t.NumField())
	fmt.Printf("TypeOf : %v\n", t)
	fmt.Printf("ValueOf IP field: %s\n", v.FieldByName("IP"))
	fmt.Printf("First Fields: %+v\n", t.Field(0))
	fmt.Printf("Get First Field tag: %+v\n", t.Field(0).Tag.Get("json"))

}
