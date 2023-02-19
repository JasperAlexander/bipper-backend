package database

import (
	"bipper-backend/graph/model"
	"bipper-backend/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to db")

	return &DB{
		client: client,
	}
}

func (db *DB) SaveUser(input model.CreateUser) *models.User {
	collection := db.client.Database("bipper").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := models.User{
		Email:        input.Email,
		Name:         input.Name,
		Username:     *input.Username,
		CreatedAt:    time.Now(),
		Description:  *input.Description,
		ProfileImg:   *input.ProfileImg,
		BannerImg:    *input.BannerImg,
		Dob:          *input.Dob,
		Location:     *input.Location,
		Url:          *input.URL,
		PinnedBeepID: *input.PinnedBeepID,
		Protected:    *input.Protected,
		Config:       *input.Config,
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	return &models.User{
		ID:           res.InsertedID.(primitive.ObjectID).Hex(),
		Email:        input.Email,
		Name:         input.Name,
		Username:     *input.Username,
		CreatedAt:    time.Now(),
		Description:  *input.Description,
		ProfileImg:   *input.ProfileImg,
		BannerImg:    *input.BannerImg,
		Dob:          *input.Dob,
		Location:     *input.Location,
		Url:          *input.URL,
		PinnedBeepID: *input.PinnedBeepID,
		Protected:    *input.Protected,
		Config:       *input.Config,
	}
}

func (db *DB) SaveBeep(input model.CreateBeep) *models.Beep {
	collection := db.client.Database("bipper").Collection("beeps")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	beep := models.Beep{
		UserID: input.UserID,
		Text:   input.Text,
		// ConversationID:    *input.ConversationID,
		// ReferenceID:       *input.ReferenceID,
		CreatedAt: time.Now(),
		// Sensitive:         *input.Sensitive,
		// Config:            *input.Config,
		// Location:          *input.Location,
		// ImpressionCount:   0,
		// RebeepCount:      0,
		// QuoteCount:        0,
		// LikeCount:         0,
		// ReplyCount:        0,
		// UrlClickCount:     0,
		// ProfileClickCount: 0,
		// DetailsClickCount: 0,
		// VideoViewCount:    0,
	}

	log.Print("before inserting")
	res, err := collection.InsertOne(ctx, beep)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("after inserting")

	return &models.Beep{
		ID:     res.InsertedID.(primitive.ObjectID).Hex(),
		UserID: input.UserID,
		Text:   input.Text,
		// ConversationID: *input.ConversationID,
		// ReferenceID:    *input.ReferenceID,
		CreatedAt: time.Now(),
		// Sensitive:      *input.Sensitive,
		// Config:         *input.Config,
		// Location:       *input.Location,
	}
}

func (db *DB) UpdateBeep(input model.UpdateBeep) *models.Beep {
	ObjectID, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		log.Fatal(err)
	}

	collection := db.client.Database("bipper").Collection("beeps")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	beep := models.Beep{}
	if input.Text != nil {
		beep.Text = *input.Text
	}
	if input.Sensitive != nil {
		beep.Sensitive = *input.Sensitive
	}

	_, updateErr := collection.UpdateOne(ctx, bson.M{"_id": ObjectID}, bson.M{"$set": beep})
	if updateErr != nil {
		log.Fatal(updateErr)
	}

	return &models.Beep{
		ID: input.ID,
	}
}

func (db *DB) UpdateBeepMetrics(input model.UpdateBeepMetrics) *models.Beep {
	ObjectID, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		log.Fatal(err)
	}

	collection := db.client.Database("bipper").Collection("beeps")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	beep := models.Beep{}
	if input.ImpressionCount != nil {
		beep.ImpressionCount = int32(*input.ImpressionCount)
	}
	if input.RebeepCount != nil {
		beep.RebeepCount = int32(*input.RebeepCount)
	}
	if input.QuoteCount != nil {
		beep.QuoteCount = int32(*input.QuoteCount)
	}
	if input.LikeCount != nil {
		beep.LikeCount = int32(*input.LikeCount)
	}
	if input.ReplyCount != nil {
		beep.ReplyCount = int32(*input.ReplyCount)
	}
	if input.URLClickCount != nil {
		beep.UrlClickCount = int32(*input.URLClickCount)
	}
	if input.ProfileClickCount != nil {
		beep.ProfileClickCount = int32(*input.ProfileClickCount)
	}
	if input.DetailsClickCount != nil {
		beep.DetailsClickCount = int32(*input.DetailsClickCount)
	}
	if input.VideoViewCount != nil {
		beep.VideoViewCount = int32(*input.VideoViewCount)
	}

	_, updateErr := collection.UpdateOne(ctx, bson.M{"_id": ObjectID}, bson.M{"$set": beep})
	if updateErr != nil {
		log.Fatal(updateErr)
	}

	return &models.Beep{
		ID: input.ID,
	}
}

func (db *DB) FindUserByID(ID string) *models.User {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}

	collection := db.client.Database("bipper").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	user := models.User{}
	res.Decode(&user)

	return &user
}

func (db *DB) FindUserByUsername(Username string) *models.User {
	collection := db.client.Database("bipper").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("finding user")
	res := collection.FindOne(ctx, bson.M{"username": Username})
	user := models.User{}
	res.Decode(&user)
	log.Printf("After searching user: %v", user)

	return &user
}

func (db *DB) FindBeepByID(ID string) *models.Beep {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}

	collection := db.client.Database("bipper").Collection("beeps")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	beep := models.Beep{}
	res.Decode(&beep)

	return &beep
}

func (db *DB) AllUsers() []*models.User {
	collection := db.client.Database("bipper").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var users []*models.User
	for cur.Next(ctx) {
		var user *models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}

func (db *DB) AllBeeps() []*models.Beep {
	collection := db.client.Database("bipper").Collection("beeps")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{}, options.Find().SetSort(bson.D{primitive.E{Key: "createdat", Value: -1}}))
	// cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var beeps []*models.Beep
	for cur.Next(ctx) {
		var beep *models.Beep
		err := cur.Decode(&beep)
		if err != nil {
			log.Fatal(err)
		}
		beeps = append(beeps, beep)
	}

	return beeps
}

func (db *DB) AllBeepsOfUserByUsername(username string) []*models.Beep {
	user := db.FindUserByUsername(username)

	collection := db.client.Database("bipper").Collection("beeps")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{"userid": user.ID}, options.Find().SetSort(bson.D{primitive.E{Key: "createdat", Value: -1}}))
	if err != nil {
		log.Fatal(err)
	}

	var beeps []*models.Beep
	for cur.Next(ctx) {
		var beep *models.Beep
		err := cur.Decode(&beep)
		if err != nil {
			log.Fatal(err)
		}
		beeps = append(beeps, beep)
	}

	return beeps
}
