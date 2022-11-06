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
	// status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
