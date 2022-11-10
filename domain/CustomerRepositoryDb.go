package domain

import (
	"database/sql"
	"microservicesAPIDevInGolang/errs"
	"microservicesAPIDevInGolang/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers`
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?`
		err = d.client.Select(&customers, findAllSql, status)
		// rows, err = d.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpetd database error")
	}

	// sqlx Start
	/*
		err = sqlx.StructScan(rows, &customers)
		if err != nil {
			logger.Error("Error while scanning customers table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpetd database error")
		}
	*/
	// sqlx End
	/*
		for rows.Next() {
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
			if err != nil {
				logger.Error("Error while scanning customers table " + err.Error())
				return nil, errs.NewUnexpectedError("Unexpetd database error")
			}
			customers = append(customers, c)
		}
	*/
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?`

	// row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customers table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
