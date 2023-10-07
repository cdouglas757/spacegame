package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewSpaceMenStore() (s *SpaceMenStore, err error) {
	s = &SpaceMenStore{}

	if s.db == nil {
		db, err := sql.Open("sqlite3", "sqlite\\spacegame.db")

		if err != nil {
			return s, err
		}

		s.db = db
	}
	return
}

type SpaceMenStore struct {
	db *sql.DB
}

type SpaceManDTO struct {
	Name string
	Rank string
	xp   int
}

func (s *SpaceMenStore) GetSpaceManDetails(id int) *SpaceManDTO {
	sdto := &SpaceManDTO{}
	query := `
		SELECT s.name, r.displaytext, s.xp FROM spacemen s 
		JOIN ranks r on s.currentrank = r.id
		WHERE s.id =?`
	stmt, err := s.db.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRow(id)

	err = row.Scan(&sdto.Name, &sdto.Rank, &sdto.xp)

	if err != nil {
		log.Fatal(err)
	}

	return sdto

}
