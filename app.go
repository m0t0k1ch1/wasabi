package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/codegangsta/negroni"
	"github.com/garyburd/redigo/redis"
	"github.com/lestrrat/go-server-starter/listener"
	"github.com/m0t0k1ch1/ksatriya"
	"github.com/m0t0k1ch1/slackbot"
)

type Wasabi struct {
	*negroni.Negroni
	actions map[string]Action
	conf    *Config
}

func (wasabi *Wasabi) NewContext(w http.ResponseWriter, req *http.Request, args ksatriya.Args) ksatriya.Ctx {
	redisAddr := fmt.Sprintf("%s:%s", wasabi.conf.Redis.Host, wasabi.conf.Redis.Port)
	redisConn, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		log.Fatal(err)
	}

	slackConn := slackbot.NewClient(wasabi.conf.Slack.Token)
	slackConn.SetAsUser(true)

	return &Context{
		Ctx:       ksatriya.NewContext(w, req, args),
		redisConn: redisConn,
		slackConn: slackConn,
		actions:   wasabi.actions,
		conf:      wasabi.conf,
	}
}

func (wasabi *Wasabi) Run() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM)
	go func() {
		for {
			s := <-signalChan
			if s == syscall.SIGTERM {
				manners.Close()
			}
		}
	}()

	var l net.Listener
	listeners, err := listener.ListenAll()
	if err != nil {
		if err == listener.ErrNoListeningTarget {
			l, err = net.Listen("tcp", fmt.Sprintf(":%s", wasabi.conf.Port))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		l = listeners[0]
	}

	manners.Serve(l, wasabi)
}

func New(confPath string) *Wasabi {
	wasabi := &Wasabi{
		Negroni: negroni.Classic(),
		conf:    NewConfig(confPath),
		actions: NewActionMap(),
	}

	k := ksatriya.New()
	k.SetCtxBuilder(wasabi.NewContext)

	k.POST("/", ActionHandler)

	wasabi.UseHandler(k)

	return wasabi
}

func ActionHandler(kctx ksatriya.Ctx) {
	action(convertContext(kctx))
}
func action(ctx *Context) {
	text := ctx.ParamSingle("text")
	trigger := ctx.ParamSingle("trigger_word")

	cmd := NewCommand(text, trigger)
	res := ctx.actions[cmd.Name](ctx, cmd.Args)

	ctx.slackConn.SendMessage(res.Channel, res.Text)
}
