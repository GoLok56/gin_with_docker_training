package models

type Controller func(data map[string]interface{}) (int, map[string]interface{})
