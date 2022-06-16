package services

// import (
// 	"github.com/jinzhu/gorm"
// 	"github.com/salvador-lucas/english-classes-api/utils"
// 	"github.com/salvador-lucas/english-classes-api/views"
// 	"github.com/sirupsen/logrus"
// )

// type PaymentsService interface {
// 	AddPayment(request *views.AddPaymentRequest) (*views.AddPaymentResponse, *utils.Error)
// }

// type paymentsService struct {
// 	db     *gorm.DB
// 	logger *logrus.Logger
// }

// func NewPaymentsService(deps utils.Dependencies) *paymentsService {
// 	return &paymentsService{
// 		db:     deps.Db,
// 		logger: deps.Logger,
// 	}
// }

// func (p *paymentsService) AddPayment(request *views.AddPaymentRequest) (*views.AddPaymentResponse, *utils.Error) {
// 	payment, err := p.createPayment(request)
// 	if err != nil {
// 		p.logger.Error(err)
// 		return nil, err
// 	}

// 	if err := p.db.Create(payment).Error; err != nil {
// 		p.logger.Error(err)
// 		if err, ok := err.(*mysql.MySQLError); ok && err != nil && err.Number == 1062 {
// 			return nil, utils.ErrorInvalid("unable to add new payment")
// 		}
// 		return nil, utils.ErrorInvalid("internal server error")
// 	}

// 	return &views.AddPaymentResponse{
// 		Amount:   payment.Amount,
// 		StudenID: payment.ID,
// 		Period:   payment.Period,
// 	}, nil
// }

// func (p *paymentsService) createPayment(request *views.AddPaymentRequest) (*models.Payment, *utils.Error) {
// 	return &models.Payment{
// 		Amount:    request.Amount,
// 		StudentID: request.StudenID,
// 		Period:    request.Period,
// 	}, nil
// }
