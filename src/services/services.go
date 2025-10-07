package services

import (
	"github.com/Buckshot-77/expertscrud/src/interfaces"
	"github.com/Buckshot-77/expertscrud/src/repositories"
	"github.com/Buckshot-77/expertscrud/src/services/installment"
)

type ServiceContainer struct {
	InstallmentService interfaces.InstallmentService
}

func GetServices(repositoryContainer repositories.RepositoriesContainer) ServiceContainer {
	return ServiceContainer{
		InstallmentService: installment.NewInstallmentService(repositoryContainer),
	}
}
