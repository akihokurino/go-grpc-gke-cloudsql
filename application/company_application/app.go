package company_application

import (
	"gke-go-recruiting-server/adapter"
	"gke-go-recruiting-server/domain/account_domain"
)

type app struct {
	tx          adapter.TX
	companyRepo adapter.CompanyRepo
}

func NewApp(
	tx adapter.TX,
	companyRepo adapter.CompanyRepo) adapter.CompanyApp {
	return &app{
		tx:          tx,
		companyRepo: companyRepo,
	}
}

func (a *app) OperationBuild(me *account_domain.AgencyAccount) adapter.CompanyAppForOperation {
	return &operationApp{
		me:  me,
		app: a,
	}
}
