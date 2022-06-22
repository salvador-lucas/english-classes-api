package services

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/salvador-lucas/english-classes-api/models"
	"github.com/salvador-lucas/english-classes-api/utils"
	"github.com/salvador-lucas/english-classes-api/views"
)

type StudentsService interface {
	GetAllStudents() (*views.GetAllStudentsResponse, *utils.Error)
	AddStudent(request *views.AddStudentRequest) (*views.AddStudentResponse, *utils.Error)
	AddStudentPayment(studentID int, request *views.AddPaymentRequest) (*views.AddPaymentResponse, *utils.Error)
	GetStudentPayments(studentID int) (*views.GetStudentPaymentsResponse, *utils.Error)
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

func (s *studentService) AddStudent(request *views.AddStudentRequest) (*views.AddStudentResponse, *utils.Error) {
	student, err := s.createStudent(request)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	if err := s.db.Create(student).Error; err != nil {
		s.logger.Error(err)
		if err, ok := err.(*mysql.MySQLError); ok && err != nil && err.Number == 1062 {
			return nil, utils.ErrorInvalid("unable to add new student")
		}
		return nil, utils.ErrorInvalid("internal server error")
	}

	return &views.AddStudentResponse{
		Name:     student.Name,
		Lastname: student.Lastname,
	}, nil
}

func (s *studentService) AddStudentPayment(studenID int, request *views.AddPaymentRequest) (*views.AddPaymentResponse, *utils.Error) {
	payment, err := s.buildPayment(studenID, request)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	if err := s.db.Create(&payment).Error; err != nil {
		s.logger.Error(err)
		if err, ok := err.(*mysql.MySQLError); ok && err != nil {
			return nil, utils.ErrorInvalid("unable to add new payment")
		}
		return nil, utils.ErrorInternal("internal server error")
	}
	return &views.AddPaymentResponse{
		Amount: payment.Amount,
		Period: payment.Period,
	}, nil
}

func (s *studentService) GetStudentPayments(studentID int) (*views.GetStudentPaymentsResponse, *utils.Error) {
	var studentPayments []models.Payment
	if err := s.db.Where(&models.Payment{
		StudentID: uint(studentID),
	}).Find(&studentPayments).Error; err != nil {
		s.logger.Error(err)
		return nil, utils.ErrorInternal(err.Error())
	}
	var studentPaymentsResponse []views.AddPaymentResponse
	for _, payment := range studentPayments {
		paymentRes := views.AddPaymentResponse{
			Amount: payment.Amount,
			Period: payment.Period,
		}
		studentPaymentsResponse = append(studentPaymentsResponse, paymentRes)
	}
	return &views.GetStudentPaymentsResponse{Payments: studentPaymentsResponse}, nil
}

// ===//=/=//=== private functions ===//=/=//===

func (s *studentService) createStudent(request *views.AddStudentRequest) (*models.Student, *utils.Error) {
	return &models.Student{
		Name:     request.Name,
		Lastname: request.Lastname,
		Enabled:  true,
	}, nil
}

func (s *studentService) buildPayment(studenID int, request *views.AddPaymentRequest) (*models.Payment, *utils.Error) {

	if request.Amount <= 0 {
		err := errors.New("Invalid amount")
		return nil, utils.ErrorInternal(err.Error())
	}

	return &models.Payment{
		StudentID: uint(studenID),
		Amount:    request.Amount,
		Period:    request.Period,
	}, nil
}
