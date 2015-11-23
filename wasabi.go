package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fukata/golang-stats-api-handler"
	"github.com/garyburd/redigo/redis"
	"github.com/m0t0k1ch1/potto"
)

type Wasabi struct {
	*potto.Potto
	conf *Config
}

func (wsb *Wasabi) NewContext(w http.ResponseWriter, req *http.Request, args potto.Args) potto.Ctx {
	pctx := wsb.Potto.NewContext(w, req, args)

	redisAddr := fmt.Sprintf("%s:%s", wsb.conf.Redis.Host, wsb.conf.Redis.Port)
	redisConn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatal(err)
	}

	return &Context{
		Context: pctx.(*potto.Context),
		conf:    wsb.conf,
		redis:   NewRedis(redisConn),
	}
}

func New(confPath string) *Wasabi {
	wsb := &Wasabi{
		Potto: potto.New(),
		conf:  NewConfig(confPath),
	}

	wsb.SetCtxBuilder(wsb.NewContext)
	wsb.AddRoute("GET", "/stats", Stats)

	wsb.AddAction("ping", Ping)
	wsb.AddAction("init", Initialize)
	wsb.AddAction("show", Show)
	wsb.AddAction("add", Add)
	wsb.AddAction("del", Remove)
	wsb.AddAction("pick", Pick)

	return wsb
}

func Stats(pctx potto.Ctx) {
	pctx.RenderJSON(http.StatusOK, stats_api.GetStats())
}
