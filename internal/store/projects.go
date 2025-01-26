package store

import (
	"portfolio/data"
	"portfolio/types"
)

type ProjectStore interface {
	GetProjects(count int) []types.Project
}

type Store struct{}

func (s *Store) GetProjects(count int) {
	finalProj := make([]types.Project, count)
	if count <= len(data.Projects) {
		return data.Projects[0 : count+1]
	}

	idx := 0
	for len(finalProj) < count {
		
	}
}
