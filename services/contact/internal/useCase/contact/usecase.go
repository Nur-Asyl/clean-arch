package contact

func NewUseCase(repo repository.Repository) usecase.UseCase {
	return usecase.NewUseCase(repo)
}
