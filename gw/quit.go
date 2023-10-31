package gw

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Moonyongjung/xns-gateway/util"
)

type Quit struct {
	Code int
}

func (q Quit) Error() string {
	return strconv.Itoa(q.Code)
}

func GatewayQuit() Quit {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	stopSign := <-stop
	util.LogInfo("shutting down the gateway...")
	util.LogInfo("gateway gracefully stopped")
	return Quit{Code: int(stopSign.(syscall.Signal))}
}
