package db

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"thuanle/cse-mark/internal/models"
)

func (db *Db) UpdateUserSettings(u *models.UserSettingsModel) error {
	update := bson.M{"$set": u}
	_, err := db.settingsUsers.UpdateByID(context.Background(), u.UserId, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) GetUserByIntId(chatId int64) (*models.UserSettingsModel, error) {
	filter := bson.M{"_id": chatId}
	var result models.UserSettingsModel
	err := db.settingsUsers.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Error().Err(err).Msg("Get user error")
		return nil, err
	}
	return &result, nil
}

func (db *Db) GetUserById(user string) (*models.UserSettingsModel, error) {
	filter := bson.M{"_id": user}
	var result models.UserSettingsModel
	err := db.settingsUsers.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Error().Err(err).Msg("Get user error")
		return nil, err
	}
	return &result, nil
}

func (db *Db) AppendUserHistory(chatId int64, course string, studentId string) error {
	update := bson.M{
		"$set": bson.M{
			"query_id": studentId,
		},
		"$addToSet": bson.M{
			"query_courses": course,
		},
		"$setOnInsert": bson.M{
			"_id": chatId,
		},
	}
	_, err := db.settingsUsers.UpdateByID(context.Background(), chatId, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) ClearUserHistory(chatId int64) error {
	update := bson.M{
		"$unset": bson.M{
			"query_id":      "",
			"query_courses": []string{},
		},
	}
	_, err := db.settingsUsers.UpdateByID(context.Background(), chatId, update, options.Update().SetUpsert(true))
	if err != nil {
		log.Error().Err(err).Msg("Clear user history error")
	}
	return nil
}
