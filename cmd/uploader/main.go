package main

import (
	"bytes"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("test")
	opts := options.GridFSBucket().SetName("custom name")
	bucket, err := gridfs.NewBucket(db, opts)
	if err != nil {
		panic(err)
	}
	var name string
	for i := 0; i < 1000000; i++ {
		uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"metadata tag", "first"}})
		name = fmt.Sprintf("test%d.json", i)
		objectID, err := bucket.UploadFromStream(name, bytes.NewReader([]byte(fmt.Sprintf(`{"name":"%s"}`, name))), uploadOpts)
		if err != nil {
			panic(err)
		}
		fmt.Printf("New file uploaded with ID %s", objectID)
	}
}
