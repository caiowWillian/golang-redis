package main

import (
	"module/src/helpers/redisHelper"
)

func main() {

	for i := 0; i < 10; i++ {
		redisHelper.Set()
	}
}
