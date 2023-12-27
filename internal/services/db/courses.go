package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"thuanle/cse-mark/internal/models"
	"time"
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

func (db *Db) GetCourseSettingsById(course string) (*models.CourseSettingsModel, error) {
	var res models.CourseSettingsModel
	err := db.settingsCourses.FindOne(context.Background(), bson.M{"_id": course}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (db *Db) SetCourseSettings(course *models.CourseSettingsModel) error {
	update := bson.M{
		"$set": bson.M{
			"link":       course.Link,
			"updated_at": time.Now().Unix(),
		},
		"$setOnInsert": bson.M{
			"_id":     course.Course,
			"course":  course.Course,
			"by_id":   course.ById,
			"by_user": course.ByUser,
		},
	}
	_, err := db.settingsCourses.UpdateByID(context.Background(), course.Course, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}
