package main

import (
	"os"
	"github.com/subosito/gotenv"
)

var Server, Admin string
var C_key, C_secret, C_token string
var BotName, User, Pass string
var Timer, Source string

func Init() {

	gotenv.Load()

	Server = os.Getenv("server")
	Admin = os.Getenv("admin")
	C_key = os.Getenv("c_key")
	C_secret = os.Getenv("c_secret")
	C_token = os.Getenv("c_token")
	BotName = os.Getenv("botName")
	User = os.Getenv("user")
	Pass = os.Getenv("pass")
	Timer = os.Getenv("timer")
	Source = os.Getenv("source")


}


