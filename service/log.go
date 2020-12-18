package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"

	"github.com/dgkg/keypass/queue"
)

type ServiceLog struct {
	kafkaReader func(string, string) *kafka.Reader
	cache       *redis.Client
	elastic     *elasticsearch.Client
}

func NewLog(cache *redis.Client, elastic *elasticsearch.Client, reader func(string, string) *kafka.Reader) *ServiceLog {

	// ss := &ServiceLog{
	// 	kafkaReader: reader,
	// 	cache:       cache,
	// 	elastic:     elastic,
	// }
	// go ss.sendMessageToRedis()
	// go ss.sendMessageToElastic()
	// return ss
	return nil
}

func (ss *ServiceLog) StatLogin(ctx *gin.Context) {
	now := fmt.Sprintf("%v-%v-%v", time.Now().Month, time.Now().Day, time.Now().Hour)
	res := ss.cache.PFCount(ctx, now)
	fmt.Printf("key %v res %v", now, res)
	ctx.JSON(http.StatusOK, gin.H{"stat": res.Val()})
}

func (ss *ServiceLog) sendMessageToElastic() {
	reader := ss.kafkaReader("group-elastic", queue.TopicLog)
	ctx := context.Background()
	for {
		msg, origin, err := ss.parseMessage(ctx, reader)
		if err != nil {
			log.Println(err)
		}
		log.Println(msg, origin)

		req := esapi.IndexRequest{
			Index:      "message_log",
			DocumentID: uuid.New().String(),
			Body:       bytes.NewBuffer(origin.Value),
			Refresh:    "true",
		}

		// Perform the request with the client.
		res, err := req.Do(context.Background(), ss.elastic)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		log.Println(res)
		reader.CommitMessages(ctx, *origin)
	}
}

func (ss *ServiceLog) sendMessageToRedis() {
	reader := ss.kafkaReader("group-redis", queue.TopicLog)
	ctx := context.Background()
	for {

		msg, origin, err := ss.parseMessage(ctx, reader)
		if err != nil {
			log.Println(err)
		}
		now := fmt.Sprintf("%v-%v-%v", msg.EventTime.Month, msg.EventTime.Day, msg.EventTime.Hour)
		pf := ss.cache.PFAdd(ctx, now, string(origin.Key))
		if pf.Err != nil {
			fmt.Println("service/stats: err cache", pf.Err())
		}
		reader.CommitMessages(ctx, *origin)
	}
}

func (ss *ServiceLog) parseMessage(ctx context.Context, reader *kafka.Reader) (*queue.MessageLog, *kafka.Message, error) {
	m, err := reader.FetchMessage(ctx)
	if err != nil {
		return nil, nil, err
	}
	var msg queue.MessageLog
	return &msg, &m, json.Unmarshal(m.Value, &msg)
}
