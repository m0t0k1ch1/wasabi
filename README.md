# wasabi

wasabi sushi :ghost:

ref. [なんでもルーレットしてくれる Slack bot](http://m0t0k1ch1st0ry.com/blog/2015/11/23/wasabi)

![wasabi](http://m0t0k1ch1st0ry.com/my-images/entry/wasabi.png)

``` sh
$ go get github.com/m0t0k1ch1/wasabi
```

## Run

``` sh
$ wasabi --conf=/path/to/your/config.tml
```

Using [go-server-starter](https://github.com/lestrrat/go-server-starter), you can restart the application gracefully.

``` sh
$ start_server --port 8080 -- wasabi --conf=/path/to/your/config.tml
```

## Usage

* `init`：delete all members in a set
* `show`：get all members in a set
* `add member [member...]`：add one or more members to a set
* `del member [member...]`：delete one or more members
* `pick`：get random member from a set
