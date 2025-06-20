package main

import (
	"backend/config"
	"backend/service/classroom/api"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/server.yaml", "the config file")

func main() {

	flag.Parse()

	//load cấu hình yaml
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//rest server
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//khởi tạo classroom service
	classroomService := api.NewClassroomService(server)
	classroomService.Start()

	fmt.Printf(" ✅✅ Starting classroom server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}
