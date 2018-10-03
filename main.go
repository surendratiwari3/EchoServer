package main
import (
	"github.com/labstack/echo"
	logger "echoServer/logger"
)
func main(){
	log := logger.NewLogger("info.log","info")
	e := echo.New()
	if err := e.Start("0.0.0.0:10000"); err != nil {
		log.WithError(err).Fatal("echo server not able to start")
	}
}
