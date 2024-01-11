package db

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *Db) GetMark(course string, id string) (string, error) {
	filter := bson.M{"_id": id}
	var result map[string]interface{}
	err := db.mark.Collection(course).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Error().
			Str("course", course).
			Err(err).
			Msg("Get marks error")
		return "", err
	}

	jsonStr, err := json.MarshalIndent(result, "", " ")

	return string(jsonStr), nil
}
