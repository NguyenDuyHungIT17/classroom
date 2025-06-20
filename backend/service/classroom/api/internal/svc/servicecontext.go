package svc

import (
	"backend/service/classroom/api/internal/config"
	"backend/service/classroom/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UsersModel              model.UsersModel
	ClassesModel            model.ClassesModel
	EnrollmentsModel        model.EnrollmentsModel
	SubjectsModel           model.SubjectsModel
	LessonsModel            model.LessonsModel
	GradeComponentsModel    model.GradeComponentsModel
	GradesModel             model.GradesModel
	EmailConfirmationsModel model.EmailConfirmationsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config: c,

		UsersModel:              model.NewUsersModel(conn),
		ClassesModel:            model.NewClassesModel(conn),
		EnrollmentsModel:        model.NewEnrollmentsModel(conn),
		SubjectsModel:           model.NewSubjectsModel(conn),
		LessonsModel:            model.NewLessonsModel(conn),
		GradeComponentsModel:    model.NewGradeComponentsModel(conn),
		GradesModel:             model.NewGradesModel(conn),
		EmailConfirmationsModel: model.NewEmailConfirmationsModel(conn),
	}
}
