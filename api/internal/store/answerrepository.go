package store

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/Kirillznkv/new_year/api/internal/model"
	"log"
	"os"
	"strings"
)

type AnswersRepository struct {
	store *Store
}

func (r *AnswersRepository) Create(a *model.Answer) error {
	if err := r.store.db.QueryRow("INSERT INTO answers (lvl, user_id, image) VALUES ($1, $2, $3) RETURNING id",
		a.Lvl,
		a.UserId,
		a.Image,
	).Scan(&a.ID); err != nil {
		return err
	}

	return r.saveImage(a)
}

func (r *AnswersRepository) saveImage(a *model.Answer) error {
	u, err := r.store.Users().FindById(a.UserId)
	if err != nil {
		return err
	}

	savePath := fmt.Sprintf("./images/lvl_%d/%s.jpg", a.Lvl, u.FirstName+"_"+u.SecondName)

	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	b64Parts := strings.Split(a.Image, ",")

	var idx int
	if l := len(b64Parts); l == 1 {
		idx = 0
	} else if l == 2 {
		idx = 1
	} else {
		return errors.New("invalid image")
	}

	data, err := base64.StdEncoding.DecodeString(b64Parts[idx])
	if err != nil {
		return err
	}

	_, err = file.Write(data)

	return err
}

func (r *AnswersRepository) GetAnswers() []*model.Answer {
	var answers []*model.Answer

	rows, err := r.store.db.Query("SELECT id, lvl, user_id, image FROM answers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		ans := &model.Answer{}

		if err := rows.Scan(&ans.ID, &ans.Lvl, &ans.UserId, &ans.Image); err != nil {
			log.Fatal(err)
		}

		answers = append(answers, ans)
	}

	return answers
}
