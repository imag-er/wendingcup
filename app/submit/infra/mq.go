package infra

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

var (
	config = &golang.Config{
		Endpoint:      "localhost:8081",
		ConsumerGroup: "wendingcup-submit",
		Credentials: &credentials.SessionCredentials{
			AccessKey:    "xxxxxx",
			AccessSecret: "xxxxxx",
		},
	}
	topic = "wendingcup-submit"

	Producer golang.Producer
	Consumer golang.SimpleConsumer

	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
	err error
)

func InitMQ() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()

	initProducer()
	initConsumer()
}
func initProducer() {
	// new producer instance
	Producer, err = golang.NewProducer(
		config,
		golang.WithTopics(topic),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = Producer.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func initConsumer() {
	// new consumer instance
	Consumer, err := golang.NewSimpleConsumer(
		config,
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			topic: golang.SUB_ALL,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = Consumer.Start()
	if err != nil {
		log.Fatal(err)
	}
	// gracefule stop Consumer
	defer Consumer.GracefulStop()

	go func() {
		for {
			fmt.Println("start recevie message")
			mvs, err := Consumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
			if err != nil {
				fmt.Println(err)
			}
			// ack message
			for _, mv := range mvs {
				Consumer.Ack(context.TODO(), mv)
				fmt.Println(string(mv.GetBody()))
			}
			fmt.Println("wait a moment")
			fmt.Println()
			time.Sleep(time.Second * 3)
		}
	}()
	// run for a while
	time.Sleep(time.Minute)
}

func SendMQ(submitId uint32) (err error) {
	msg := &golang.Message{
		Topic: topic,
		Body:  []byte("this is a message : " + strconv.Itoa(i)),
	}
	// set keys and tag
	msg.SetKeys("a", "bhhjhjhn")
	msg.SetTag("ab")

	// send message in sync
	resp, err := Producer.Send(context.TODO(), msg)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(resp); i++ {
		fmt.Printf("%#v\n", resp[i])
	}
	// wait a moment
	time.Sleep(time.Second * 1)
}
