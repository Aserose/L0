package errorInfo

import "L0/internal/repository"

type ErrorInfo struct {
	repo repository.Repository
}

func NewErrInfo(repo repository.Repository) ErrorInfo {
	return ErrorInfo{
		repo: repo,
	}
}

func (d ErrorInfo) GetLast() string {
	return d.repo.ErrorCache.GetLast()
}
