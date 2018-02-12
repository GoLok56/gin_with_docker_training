package models

type Controller func(data ...interface{}) (int, map[string]interface{})
