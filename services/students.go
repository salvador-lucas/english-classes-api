package services

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/salvador-lucas/english-classes-api/models"
	"github.com/salvador-lucas/english-classes-api/utils"
	"github.com/salvador-lucas/english-classes-api/views"
)

type StudentsService interface {
	GetAllStudents() (*views.GetAllStudentsResponse, *utils.Error)
}

type studentService struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewStudentService(deps utils.Dependencies) *studentService {
	return &studentService{
		db:     deps.Db,
		logger: deps.Logger,
	}
}

func (s *studentService) GetAllStudents() (*views.GetAllStudentsResponse, *utils.Error) {
	var students []models.Student
	if err := s.db.Find(&students).Error; err != nil {
		return nil, utils.ErrorInternal(err.Error())
	}

	var studentsResponse []views.GetStudentResponse
	for _, student := range students {
		studentResponse := views.GetStudentResponse{
			Name:     student.Name,
			Lastname: student.Lastname,
		}
		studentsResponse = append(studentsResponse, studentResponse)
	}
	return &views.GetAllStudentsResponse{Students: studentsResponse}, nil
}
