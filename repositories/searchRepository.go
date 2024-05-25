package repositories

import (
	"context"
	"fmt"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/responses"
)

func addCondition(query []string, conditionColumn string) string {
	conditions := fmt.Sprintf("%s ILIKE '%%%s%%'", conditionColumn, query[0])
	size := len(query)
	fmt.Println("Length:", len(query))
	for i := 1; i < size; i++ {
		fmt.Print(query[i])
		conditions += fmt.Sprintf(" OR %s ILIKE '%%%s%%'", conditionColumn, query[i])
	}
	return conditions
}

func GetMentor(ctx context.Context, query []string) ([]responses.SearchResponseDTO, error) {
	db := configs.GetDB()
	var listMentor []responses.SearchResponseDTO = nil

	err := db.WithContext(ctx).
		Raw("? UNION ?",
			db.WithContext(ctx).
				Table("mentors m").
				Select("m.id AS mentor_id, m.rating AS rating, u.username AS username, u.university AS university, u.profile_picture AS profile_picture").
				Joins("JOIN users u ON m.user_id = u.id").
				Where(addCondition(query, "username")),
			db.WithContext(ctx).
				Table("mentors m").
				Select("m.id AS mentor_id, m.rating AS rating, u.username AS username, u.university AS university, u.profile_picture AS profile_picture").
				Joins("JOIN users u ON m.user_id = u.id").
				Joins("JOIN mentor_courses mc ON mc.mentor_id = m.id").
				Joins("JOIN courses c ON c.id = mc.course_id").
				Where(addCondition(query, "c.name")),
		).
		Group("mentor_id").
		Scan(&listMentor).
		Error

	return listMentor, err
}
