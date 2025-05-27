package mongo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"thuanle/cse-mark/internal/configs"
	"thuanle/cse-mark/internal/domain/mark"
	"time"
)

type MarkRepo struct {
	db      *mongo.Database
	mu      sync.Mutex
	timeout time.Duration
}

func NewMarkRepo(client *Client, config *configs.Config) mark.Repository {
	db := client.mgClient.Database(config.DbMark)
	return &MarkRepo{
		db:      db,
		timeout: client.Timeout,
		mu:      sync.Mutex{},
	}
}

func (r *MarkRepo) GetMark(courseId string, studentId string) (string, error) {
	filter := bson.M{"_id": studentId}
	var result map[string]interface{}

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	err := r.db.Collection(courseId).FindOne(ctx, filter).Decode(&result)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		err = mark.ErrNotFound
		return "", err
	}

	jsonStr, err := json.MarshalIndent(result, "", " ")
	return string(jsonStr), nil
}

func (r *MarkRepo) RemoveMarksByCourseId(courseId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	err := r.db.Collection(courseId).Drop(ctx)
	if err != nil {
		log.Error().
			Str("course", courseId).
			Err(err).
			Msg("ClearCourse")
		return err
	}
	return nil
}

func (r *MarkRepo) AddCourseMarks(courseId string, marks []map[string]string) error {
	var bulkWrites []mongo.WriteModel
	for _, mark := range marks {
		// Ensure _id is set, if not use id
		if len(mark["_id"]) == 0 {
			mark["_id"] = mark["id"]
		}

		filter := bson.M{"_id": mark["_id"]}
		update := bson.M{"$set": mark}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		bulkWrites = append(bulkWrites, model)
	}
	bulkWriteOptions := options.BulkWrite().SetOrdered(false)

	r.mu.Lock()
	defer r.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	result, err := r.db.Collection(courseId).BulkWrite(ctx, bulkWrites, bulkWriteOptions)
	if err != nil {
		log.Error().
			Str("course", courseId).
			Err(err).
			Msg("Bulk write error")
		return err
	}

	// Print the number of modified documents
	log.Info().
		Str("course", courseId).
		Interface("result", result).
		Msg("Store marks")

	return nil
}

func (r *MarkRepo) RemoveCourseMarks(courseId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	err := r.db.Collection(courseId).Drop(ctx)
	if err != nil {
		log.Error().
			Str("courseId", courseId).
			Err(err).
			Msg("ClearCourse")

		return err
	}
	return nil
}
