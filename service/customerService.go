package service

import "microservicesAPIDevInGolang/domain"

// ServiceインタフェースとRepositoryインターフェース間のBusiness Logicの構築
// Primary portのServiceインターフェースを作成　メソッドを入れる
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// Secondary portのRepositoryとの依存性を定義
type DefaultCustomService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomService {
	return DefaultCustomService{repository}
}
