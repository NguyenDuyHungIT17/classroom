package api

//khởi tạo service
import (
	"backend/service/classroom/api/internal/config"
	"backend/service/classroom/api/internal/handler"
	"backend/service/classroom/api/internal/svc"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("classroom-config", "etc/classroom.yaml", "the config file")

// classroom service
type ClassroomService struct {
	C      config.Config       // cấu hình yaml sau khi load
	Server *rest.Server        //server http từ gozero
	Ctx    *svc.ServiceContext // context
}

// tạo classroom service từ rest server
func NewClassroomService(server *rest.Server) *ClassroomService {
	flag.Parse()

	var c config.Config

	//load config
	conf.MustLoad(*configFile, &c)

	//tạo servicecontext
	ctx := svc.NewServiceContext(c)

	//đăng kí api endpoint vào rest server
	handler.RegisterHandlers(server, ctx)

	return &ClassroomService{
		C:      c,
		Server: server,
		Ctx:    ctx,
	}
}

// chạy
func (s *ClassroomService) Start() error {
	return nil
}
