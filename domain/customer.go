package domain

import "microservicesAPIDevInGolang/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// Secondary portのRepositoryインターフェースを作成　メソッドを入れる
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
