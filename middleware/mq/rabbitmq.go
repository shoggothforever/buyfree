package mq

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type RabbitConfig struct {
	Mqurl            string
	DeadName         string
	DeadQueueName    string
	DeadExchangeName string
	DeadRouting      string
	DeadExchangeType string
	BaseName         string
	BaseExchangeName string
	BaseQueueName    string
}

const (
	mqurl                   = "amqp://dsm:wusa@localhost:5672/"
	deadName                = "dead"
	deadqueuename           = "dlx_queue"
	deadexchangename        = "dlx_exchange"
	deadrouting             = "dream.dead"
	Topic                   = "topic"
	baseName                = "dsm"
	baseexchangename        = "topic_logs"
	basequeuename           = ""
	Publishdeadexchange     = "x-dead-letter-exchange"
	Publishdeadrouter       = "x-dead-letter-routing-key"
	Xmessagettl         int = 1
	Xmaxlength          int = 1000
)

var Rbc RabbitConfig
var smqs sync.Map

// var rq RabbitMQ
var once sync.Once

func init() {
	smqs = sync.Map{}
	Rbc.Mqurl = mqurl
	Rbc.DeadName = deadName
	Rbc.DeadExchangeName = deadexchangename
	Rbc.DeadExchangeType = Topic
	Rbc.BaseQueueName = basequeuename
	Rbc.BaseName = baseName
	Rbc.BaseQueueName = basequeuename
	Rbc.BaseExchangeName = baseexchangename
	Rbc.DeadRouting = deadrouting
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	// 连接名称
	Name string
	// 交换机
	ExchangeName string
	//
	ExchangeType string
	// routing Key
	RoutingKey string
	//MQ链接字符串
	Mqurl string
}

func (rq *RabbitMQ) setMQAtt(name, ename, etype, rkey, mqurl string) {
	rq.Name = name
	rq.ExchangeName = ename
	rq.ExchangeType = etype
	rq.RoutingKey = rkey
	rq.Mqurl = mqurl
}
func NewRabbit(name, ename, etype, rkey, mqurl string) (*RabbitMQ, error) {
	var rq RabbitMQ
	var err error
	rq.setMQAtt(name, ename, etype, rkey, mqurl)
	irq, ok := smqs.Load(name)
	if !ok {
		rq.Conn, err = amqp.Dial(mqurl)
		if err != nil {
			FailOnError(err, "与RabbitMQ建立连接失败")
			return nil, err
		}
		rq.Channel, err = rq.Conn.Channel()
		if err != nil {
			FailOnError(err, "从连接中获取管道失败")
			return nil, err
		}
		err = rq.Channel.ExchangeDeclare(ename, etype, true, true, false, false, nil)
		if err != nil {
			FailOnError(err, "创建Exchanger失败")
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		smqs.Store(rq.Name, rq)
	} else {
		rq = irq.(RabbitMQ)
		fmt.Println("已经存储了该连接", rq)
	}
	return &rq, nil
}
func (rq *RabbitMQ) ReleaseMQ() {
	rq.Conn.Close()
	rq.Channel.Close()
}
func GetRabbitMQ(name string) RabbitMQ {
	irq, ok := smqs.Load(name)
	if !ok {
		logs.Info("获取连接信息失败，连接名%s", name)
		panic(ok)
	}
	return irq.(RabbitMQ)
}
func (rq RabbitMQ) PublishMessage(ctx context.Context, ename, rkey string, msg amqp.Publishing) {
	err := rq.Channel.PublishWithContext(ctx, ename, rkey, true, true, msg)
	if err != nil {
		return
	}
}

// 创建死信队列
func (rq RabbitMQ) CreateDeadQueue(dqname, dqexchange, dqrouting string, route bool) (*amqp.Queue, error) {
	//fmt.Println(rq)
	var deadq amqp.Queue
	var err error
	if route {
		deadq, err = rq.Channel.QueueDeclare(dqname, true, true, false, false, amqp.Table{
			Publishdeadexchange: dqexchange,  //"dlx_exchange",
			Publishdeadrouter:   dqrouting,   //"dlx_route",
			"x-message-ttl":     Xmessagettl, // 过期时间
			"x-max-length":      Xmaxlength}) // 死信队列配置参数)
	} else {
		deadq, err = rq.Channel.QueueDeclare(dqname, true, true, false, false, amqp.Table{
			Publishdeadexchange: dqexchange,  //"dlx_exchange",
			"x-message-ttl":     Xmessagettl, //过期时间
			"x-max-length":      Xmaxlength}) // 死信队列配置参数)
	}
	err = rq.Channel.QueueBind(dqname, dqrouting, dqexchange, false, nil)
	FailOnError(err, "绑定死信队列失败")
	return &deadq, err
}
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
