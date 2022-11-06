package service

import (
	"microservicesAPIDevInGolang/domain"
	"microservicesAPIDevInGolang/errs"
)

// ServiceインタフェースとRepositoryインターフェース間のBusiness Logicの構築
// Primary portのServiceインターフェースを作成　メソッドを入れる
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
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

func (s DefaultCustomService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomService {
	return DefaultCustomService{repository}
}
