package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/m0t0k1ch1/potto"
)

type Wasabi struct {
	*potto.Potto
	conf *Config
}

func (wsb *Wasabi) NewContext(w http.ResponseWriter, req *http.Request, args potto.Args) potto.Ctx {
	redisAddr := fmt.Sprintf("%s:%s", wasabi.conf.Redis.Host, wasabi.conf.Redis.Port)
	redisConn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatal(err)
	}

	return &Context{
		Ctx:       wsb.Potto.NewContext(w, req, args),
		conf:      wsb.conf,
		redisConn: redisConn,
	}
}

func New(confPath string) *Wasabi {
	wsb := &Wasabi{
		Potto: potto.New(),
		conf:  NewConfig(confPath),
	}

	wsb.SetCtxBuilder(wsb.NewContext)
	wsb.AddAction("ping", PingAction)

	return wsb
}