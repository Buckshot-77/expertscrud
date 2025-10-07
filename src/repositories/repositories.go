package repositories

import (
	"github.com/Buckshot-77/expertscrud/src/interfaces"
	"github.com/Buckshot-77/expertscrud/src/repositories/installment"
	"gorm.io/gorm"
)

type RepositoriesContainer struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func GetRepositories(db *gorm.DB) RepositoriesContainer {
	return RepositoriesContainer{
		InstallmentRepository: installment.NewInstallmentRepository(db),
	}
}
