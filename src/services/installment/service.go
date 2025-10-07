package installment

import (
	"errors"
	"time"

	"github.com/Buckshot-77/expertscrud/src/interfaces"
	"github.com/Buckshot-77/expertscrud/src/repositories"
	"github.com/Buckshot-77/expertscrud/src/structs"
)

type InstallmentService struct {
	InstallmentRepository interfaces.InstallmentRepository
}

func NewInstallmentService(repositoryContainer repositories.RepositoriesContainer) InstallmentService {
	return InstallmentService{
		InstallmentRepository: repositoryContainer.InstallmentRepository,
	}
}

func (is InstallmentService) Create(installment structs.Installment) error {

	if installment.DueDay < time.Now().Day() {
		return errors.New("installment overdue")
	}

	return is.InstallmentRepository.Create(installment)
}
