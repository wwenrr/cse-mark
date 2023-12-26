package db

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"thuanle/cse-mark/internal/configs"
)

type Db struct {
	client          *mongo.Client
	mark            *mongo.Database
	settings        *mongo.Database
	settingsUsers   *mongo.Collection
	settingsCourses *mongo.Collection
}

// create singleton Db
var (
	db   *Db
	once sync.Once
)

func Instance() *Db {
	once.Do(func() {
		db = &Db{}
	})
	return db
}

func (db *Db) Load() error {
	connectionString := `mongodb://` + configs.MongoHost + `:` + configs.MongoPort

	ctx, cancel := context.WithTimeout(context.Background(), configs.DbTransactionTimeout)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to MongoDB")
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Error pinging MongoDB")
		return err
	}

	log.Info().Msg("Connected to MongoDB!")

	db.client = client
	db.mark = client.Database(configs.DbMark)

	db.settings = client.Database(configs.DbSettings)
	db.settingsUsers = db.settings.Collection(configs.DbSettingsUsers)
	db.settingsCourses = db.settings.Collection(configs.DbSettingsCourses)
	return nil
}

func (db *Db) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), configs.DbTransactionTimeout)
	defer cancel()

	if err := db.client.Disconnect(ctx); err != nil {
		log.Fatal().Err(err).Msg("Error disconnecting from MongoDB")
	}
}
