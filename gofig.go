package gofig

import (
	"fmt"
	"syscall"
)

func Raw(name string) (value string, found bool) {
	value, found = syscall.Getenv(name)
	return
}

func String(name string) (value string, err error) {
	value, found := Raw(name)
	if found != true {
		err = fmt.Errorf("Configuration parameter %s not provided", name)
	}
	return
}

func StringDefault(name, defaultValue string) (value string) {
	value, err := String(name)
	if err != nil {
		value = defaultValue
	}
	return
}

func StringVar(target *string, name string) (err error) {
	*target, err = String(name)
	return
}

func StringVarDefault(target *string, name, defaultValue string) {
	*target = StringDefault(name, defaultValue)
}

func Int(name string) (value int, err error) {
	return 1, nil
}
