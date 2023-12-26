package db

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

func (db *Db) GetMark(course string, id string) (string, error) {
	filter := bson.M{"_id": id}
	var result bson.M
	err := db.mark.Collection(course).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Error().Err(err).Msg("Get mark error")
		return "", err
	}

	jsonStr, err := bson.MarshalExtJSON(result, true, false)

	return string(jsonStr), nil
}

var (
	storeMu sync.Mutex
)

func (db *Db) StoreMarks(sub string, marks []map[string]string) error {
	var bulkWrites []mongo.WriteModel
	for _, mark := range marks {
		filter := bson.M{"_id": mark["_id"]}
		update := bson.M{"$set": mark}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		bulkWrites = append(bulkWrites, model)
	}
	bulkWriteOptions := options.BulkWrite().SetOrdered(false)

	storeMu.Lock()
	defer storeMu.Unlock()
	result, err := db.mark.Collection(sub).BulkWrite(context.Background(), bulkWrites, bulkWriteOptions)
	if err != nil {
		log.Error().Err(err).Msg("Bulk write error")
		return err
	}
	// Print the number of modified documents
	log.Info().Interface("result", result).Msg("Store marks")
	return nil
}

func (db *Db) ClearMarks(sub string) error {
	storeMu.Lock()
	defer storeMu.Unlock()
	err := db.mark.Collection(sub).Drop(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Clear marks error")
		return err
	}
	return nil
}
