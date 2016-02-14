package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fukata/golang-stats-api-handler"
	"github.com/garyburd/redigo/redis"
	"github.com/m0t0k1ch1/potto"
)

type Action func(*Context, ActionArgs) (*Response, error)
type ActionArgs potto.ActionArgs

type Wasabi struct {
	*potto.Potto
	Conf *Config
}

func (wsb *Wasabi) AddAction(name string, action Action) {
	wsb.Potto.AddAction(name, func(pctx potto.Ctx, pargs potto.ActionArgs) (*potto.Response, error) {
		ctx := pctx.(*Context)
		args := ActionArgs(pargs)

		res, err := action(ctx, args)

		return res.Response, err
	})
}

func (wsb *Wasabi) NewContext(w http.ResponseWriter, req *http.Request, args potto.Args) potto.Ctx {
	pctx := wsb.Potto.NewContext(w, req, args)

	redisAddr := fmt.Sprintf("%s:%s", wsb.Conf.Redis.Host, wsb.Conf.Redis.Port)
	redisConn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatal(err)
	}

	return &Context{
		Context: pctx.(*potto.Context),
		Redis:   NewRedis(redisConn),
		Conf:    wsb.Conf,
	}
}

func New(confPath string) *Wasabi {
	wsb := &Wasabi{
		Potto: potto.New(),
		Conf:  NewConfig(confPath),
	}

	wsb.SetCtxBuilder(wsb.NewContext)
	wsb.AddRoute("GET", "/stats", Stats)

	wsb.AddAction("ping", Ping)
	wsb.AddAction("init", Init)
	wsb.AddAction("show", Show)
	wsb.AddAction("add", Add)
	wsb.AddAction("del", Del)
	wsb.AddAction("pick", Pick)

	return wsb
}

func Stats(pctx potto.Ctx) {
	pctx.RenderJSON(http.StatusOK, stats_api.GetStats())
}
