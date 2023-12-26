package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"thuanle/cse-mark/internal/models"
)

func (db *Db) GetAllCourses(updatedAfter int64) ([]*models.CourseSettingsModel, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"updated_at": bson.M{"$gt": updatedAfter}},
			{"link": bson.M{"$ne": ""}},
		},
	}

	cur, err := db.settingsCourses.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var courses []*models.CourseSettingsModel
	err = cur.All(context.Background(), &courses)
	if err != nil {
		return nil, err
	}

	return courses, err
}
