package kafka

import (
	"fmt"
	"log"
	"strconv"

	"github.com/awesome-sphere/as-seating-consumer/kafka/interfaces"
	"github.com/awesome-sphere/as-seating-consumer/utils"
	"github.com/segmentio/kafka-go"
)

var TOPIC string
var PARTITION int
var KAFKA_LOCATION string

func listTopic(connector *kafka.Conn) map[string]*interfaces.TopicInterface {
	partitions, err := connector.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}
	m := make(map[string]*interfaces.TopicInterface)
	for _, p := range partitions {
		if _, ok := m[p.Topic]; ok {
			m[p.Topic].Partition += 1
		} else {
			m[p.Topic] = &interfaces.TopicInterface{Partition: 1}
		}
	}
	return m

}

func doesTopicExist(connector *kafka.Conn, topic_name string) bool {
	topics := listTopic(connector)
	_, ok := topics[topic_name]
	return ok
}

func connectKafka() *kafka.Conn {
	KAFKA_LOCATION = fmt.Sprintf(
		"%s:%s",
		utils.GetenvOr("KAFKA_HOST", "localhost"),
		utils.GetenvOr("KAFKA_PORT", "9092"),
	)

	conn, err := kafka.Dial("tcp", KAFKA_LOCATION)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

func InitKafkaTopic() {
	TOPIC = utils.GetenvOr("KAFKA_TOPIC", "seating")
	var err error
	PARTITION, err = strconv.Atoi(utils.GetenvOr("KAFKA_TOPIC_PARTITION", "5"))
	if err != nil {
		panic(err.Error())
	}

	conn := connectKafka()
	defer conn.Close()

	if !doesTopicExist(conn, TOPIC) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             TOPIC,
				NumPartitions:     PARTITION,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}
	log.Printf("Reading topic %s", TOPIC)
	ReadMessage(TOPIC)
}
