package main

//go:generate go run gen.go
import (
	"PwnRecoli/pkg/utils"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"text/template"
	"time"
)

type Settings struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	AllOutputStr  string `yaml:"allOutputStr"`
	AllOutputByte []byte `yaml:"allOutputByte"`
}

func main() {
	var settings = Settings{}
	conf, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal("fail to read file:", err)
	}

	err = yaml.Unmarshal(conf, &settings)

	f, err := os.Create("./pkg/settings/settings.go")
	if err != nil {
		log.Fatal("fail to create file:", err)
	}

	AllOutputByteHex := string(utils.HexByte(settings.AllOutputByte))

	err = packageTemplate.Execute(f, struct {
		Timestamp     time.Time
		Host          string
		Port          string
		AllOutputByte string
		AllOutputStr  string
	}{
		Timestamp:     time.Now(),
		Host:          settings.Host,
		Port:          settings.Port,
		AllOutputByte: AllOutputByteHex,
		AllOutputStr:  settings.AllOutputStr,
	})
	if err != nil {
		return
	}
}

var packageTemplate = template.Must(template.New("").Parse(
	`// Package settings Code generated .* DO NOT EDIT\.$
// This file was generated by robots at
// {{ .Timestamp }}
package settings

var ServerHost = "{{ .Host }}"
var ServerPort = "{{ .Port }}"
var AllOutputByte  = []byte("{{ .AllOutputByte }}")
var AllOutputStr  = ` + "`" + `{{ .AllOutputStr }}` + "`"))
