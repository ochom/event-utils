package queue

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// Queue ...
type Queue struct {
	db    *redis.Client
	name  string
	items []string
}

// NewQueue creates a new instance of Queue
func NewQueue(client redis.Client) *Queue {
	return &Queue{
		db: &client,
	}
}

// SetName sets the name of the queue
func (q *Queue) SetName(name string) {
	q.name = name
}

// Load loads queue items from redis by name/key
func (q *Queue) Load(ctx context.Context) {
	items := []string{}

	data, err := q.db.Get(ctx, q.name).Bytes()
	if err == redis.Nil {
		q.items = items
		return
	}

	if data == nil {
		q.items = items
		return
	}

	if err = json.Unmarshal(data, &items); err != nil {
		q.items = items
		log.Println(err.Error())
		return
	}

	q.items = items
}

// LoadOtherQueues loads all the other keys in this db
func (q *Queue) LoadOtherQueues(ctx context.Context) []string {
	items := []string{}

	iter := q.db.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		items = append(items, iter.Val())
	}

	return items
}

// Len return count of items in queue
func (q *Queue) Len() int {
	return len(q.items)
}

// IsEmpty checks if current queue is empty
func (q *Queue) IsEmpty() bool {
	if len(q.items) == 0 {
		return true
	}

	return false
}

// GetAllItems get all items in the queue
func (q *Queue) GetAllItems() []string {
	return q.items
}

// Enqueue add items to the queue
func (q *Queue) Enqueue(ctx context.Context, item string) {
	q.items = append(q.items, item)
	q.Save(ctx)
}

// Save updates queue values in redis db
func (q *Queue) Save(ctx context.Context) {
	data, err := json.Marshal(q.items)
	if err != nil {
		log.Println(err)
	}

	// unset the queue after total * minutes assuming each item takes on minute on the queue
	expires := time.Minute * time.Duration(q.Len())

	_, err = q.db.Set(ctx, q.name, data, expires).Result()
	if err != nil {
		log.Println(err)
	}
}

// Dequeue remove first item from queue
func (q *Queue) Dequeue(ctx context.Context) *string {

	defer q.Save(ctx)

	if len(q.items) > 0 {
		item := q.items[0]
		q.items = q.items[1:]
		return &item
	}

	return nil
}

// Empty add items to the queue
func (q *Queue) Empty(ctx context.Context) {
	q.items = []string{}
	q.Save(ctx)
}
