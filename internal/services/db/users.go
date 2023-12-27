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
