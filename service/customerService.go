package service

import (
	"microservicesAPIDevInGolang/domain"
	"microservicesAPIDevInGolang/dto"
	"microservicesAPIDevInGolang/errs"
)

// ServiceインタフェースとRepositoryインターフェース間のBusiness Logicの構築
// Primary portのServiceインターフェースを作成　メソッドを入れる
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

// Secondary portのRepositoryとの依存性を定義
type DefaultCustomService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()
	return response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomService {
	return DefaultCustomService{repository}
}
