package handler

import "errors"

func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("User service --" + err.Error())
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("Task service --" + err.Error())
		panic(err)
	}
}
