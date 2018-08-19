package services

import (
	"company_info/config"
	"company_info/models/request"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	config      *config.Config
	headerServ *HeaderService
}

func NewKafkaConsumer(config *config.Config, ps *HeaderService) *KafkaConsumer {
	return &KafkaConsumer{
		config:      config,
		headerServ: ps,
	}
}

func (kc *KafkaConsumer) Run() {

	configConsumer := kafka.ConfigMap{
		"bootstrap.servers":       kc.config.BootstrapServers,
		"group.id":                kc.config.GroupID,
		"auto.offset.reset":       kc.config.AutoOffsetReset,
		"auto.commit.enable":      kc.config.AutoCommitEnable,
		"auto.commit.interval.ms": kc.config.AutoCommitInterval,
	}

	c, err := kafka.NewConsumer(&configConsumer)

	if err != nil {
		panic(err)
	}

	topicsSubs := kc.config.TopicsSubscribed
	err = c.SubscribeTopics(topicsSubs, nil)

	if err != nil {
		panic(err)
	}

	log.Println("Establishing connection with Kafka")

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {

			topic := *msg.TopicPartition.Topic

			switch topic {
			case "header":
				log.Println(`Reading a header topic message`)
				header, err := kc.parseHeaderMessage(msg.Value)
				if err != nil {
					log.Printf("Error parsing event message value. Message %v \n Error: %s\n", msg.Value, err.Error())
					break
				}

				// save header to database
				_, e := kc.headerServ.CreateOne(header)
				if e != nil {
					log.Printf("Error saving header to database\n Error: %s\n", e.Response)
					break
				}
			default: //ignore any other topics
			}
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}

func (kc *KafkaConsumer) parseHeaderMessage(messageValue []byte) (*mrequest.HeaderCreate, error) {
	header := mrequest.HeaderCreate{}
	err := json.Unmarshal(messageValue, &header)

	if err != nil {
		return nil, err
	}

	return &header, nil
}
