// Copyright 2021 Quanhui Li. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// apiserver is the api server for filestore service.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/spf13/viper"

	_ "filestore/internal/apiserver/config"
	op "filestore/internal/apiserver/options"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	fmt.Println(viper.GetInt("redis.maxactive"))
	fmt.Println(op.NewOptions().RedisOptions.NewClient().Dial())
}
