package store

import (
	"portfolio/data"
	"portfolio/internal/utils"
	"portfolio/types"
)

type ProjectStore interface {
	GetProjects() []types.Project
	GetRepos(count int) []types.Repository
}

func New() *Store {
	return &Store{}
}

type Store struct{}

func (s *Store) GetProjects() []types.Project {
	return data.Projects
}

func (s *Store) GetRepos(count int) []types.Repository {
	rep, err := utils.FetchRepositories("coleYab")
    if err != nil {
        return []types.Repository{}
    }

	return rep[0:12]
}
