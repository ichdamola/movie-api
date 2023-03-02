package app

import (
	"database/sql"

	"github.com/ichdamola/movie-api/models"
)

func cmToFeetInches(int) (string, float64) {
	return "", 0
}
func addComment(db *sql.DB, comment models.Comment) error {

	//TODO
	stmt, err := db.Prepare("INSERT")
	if err != nil {
		return err
	}

	//TODO
	if _, er := stmt.Exec(comment.ID, comment.Comment); er != nil {
		return er
	}

	return nil
}

func getComments(db *sql.DB, movieID int) ([]models.Comment, error) {

	var comments []models.Comment

	//TODO
	stmt, err := db.Prepare("FIND")
	if err != nil {
		return nil, err
	}

	//TODO
	rows, er := stmt.Query(movieID)
	if er != nil {
		return nil, er
	}

	var comment models.Comment

	for rows.Next() {
		_ = rows.Scan(&comment)
		comments = append(comments, comment)
	}

	return comments, nil
}

func getCharacters(db *sql.DB, movieID int, query any) (*models.CharacterList, error) {

	var list models.CharacterList
	var Characters []models.Character

	//TODO
	stmt, err := db.Prepare("FIND")
	if err != nil {
		return nil, err
	}

	//TODO
	rows, er := stmt.Query(movieID)
	if er != nil {
		return nil, er
	}

	var character models.Character

	for rows.Next() {
		_ = rows.Scan(&character)
		Characters = append(Characters, character)
	}

	for _, character := range Characters {
		list.TotalCm += character.HeightCm
		list.TotalCount++
	}

	for _, character := range Characters {
		list.Characters = append(list.Characters, models.CharacterList_Character{
			Name: character.Name,
		})
	}

	list.TotalFt, list.TotalIn = models.FeetsInches(list.TotalCm)

	return &list, nil

}
