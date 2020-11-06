package main

import (
    "encoding/json"
    "log"
    "os"
)

type Config struct {
    Token         string
    CommandPrefix string
    DbName        string
}

func ParseConfig() (Config) {
    file, err := os.Open("config.json")

    if err != nil {
        log.Fatal("can't open config file: ", err)
    }

    defer file.Close()
    decoder := json.NewDecoder(file)
    Config := Config{}
    err = decoder.Decode(&Config)

    if err != nil {
        log.Fatal("can't decode config JSON: ", err)
    }

    return Config
}
