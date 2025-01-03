package config

import (
    "os"
    "strconv"
)

var WGpath string
var Address string
var Port string
var Servername string
var Token string
var Debug bool

func init() {
    loadenv()
    exists := false
    Port, exists = os.LookupEnv("PORT")
    if !exists {
        Port = "3000"
    }
    WGpath, exists = os.LookupEnv("WGPATH")
    if !exists {
        WGpath = "/etc/wireguard"
    }
    Address = os.Getenv("ADDRESS")
    Servername = os.Getenv("SERVERNAME")
    Token = os.Getenv("TOKEN")
    debugstr, exists := os.LookupEnv("DEBUG")
    if !exists {
        Debug = false
    } else {
        var err error
        Debug, err = strconv.ParseBool(debugstr)
        if err != nil {
            Debug = false
        }
    }
}
