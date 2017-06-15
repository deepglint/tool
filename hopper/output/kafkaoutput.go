package output

import (
	"github.com/deepglint/hopper/config"

	"errors"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/golang/glog"
)

type KafkaWriter struct {
	producer sarama.SyncProducer
	amount   int
	writes   int64
	addr     []string
	topic    string
	buffer   chan []byte
	last     *sarama.ProducerMessage
}

func NewKafkaWriter(config *config.Config) (*KafkaWriter, error) {

	addrlist := config.KafkaOutputConfig.Hosts
	topic := config.KafkaOutputConfig.FormatFaceTopic

	if len(addrlist) == 0 {
		log.Println("kafka output address not set")
		return nil, errors.New("kafka output address not set")
	}

	if len(topic) == 0 {
		log.Println("kafka topic not set")
		return nil, errors.New("kafka topic not set")
	}

	var amount = config.SpeedLimit
	w := &KafkaWriter{
		producer: nil,
		amount:   int(amount),
		writes:   0,
		addr:     addrlist,
		topic:    topic,
		buffer:   make(chan []byte, amount*2),
		last:     nil,
	}
	glog.Errorln(topic, addrlist)
	return w, nil
}

func (w *KafkaWriter) Writes() int64 { return w.writes }

func (w *KafkaWriter) Write(b []byte, tx time.Duration) error {
	if tx == 0 {
		tx = time.Millisecond * 100
	}

	select {
	case w.buffer <- b:
		return nil
	case <-time.After(tx):
		log.Printf("[INFO] kafka %v write timeout", w.addr)
		return errors.New("write timeout")
	}
}

func (w *KafkaWriter) Setup() error {
	var err error
	var producer sarama.SyncProducer

	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.MaxMessageBytes = 104857600
	config.Consumer.Fetch.Max = 104857600
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	if producer, err = sarama.NewSyncProducer(w.addr, config); err != nil {
		log.Printf("create kafka producer to %v failed: %v", w.addr, err)
	}
	w.producer = producer
	return nil
}

func (w *KafkaWriter) Clean() (err error) {
	if w.producer != nil {
		err = w.producer.Close()
		if err != nil {
			log.Printf("[ERROR] kafka %v close error: %v", w.addr, err)
		}
	}
	return nil
}

func (w *KafkaWriter) Perform() {
	if w.producer == nil {
		var err error
		var producer sarama.SyncProducer

		config := sarama.NewConfig()
		config.Producer.Retry.Max = 5
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Compression = sarama.CompressionSnappy
		config.Producer.MaxMessageBytes = 104857600
		config.Consumer.Fetch.Max = 104857600
		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true

		if producer, err = sarama.NewSyncProducer(w.addr, config); err != nil {
			// log.Printf("[ERROR] kafka %v create producer failed: %v", w.addr, err)
			time.Sleep(time.Second)
			return
		}
		w.producer = producer
	}

	for i := 0; i < w.amount; i++ {
		select {
		case b := <-w.buffer:
			m := &sarama.ProducerMessage{
				Topic: w.topic,
				Key:   nil,
				Value: sarama.ByteEncoder(b),
			}
			_, _, err := w.producer.SendMessage(m)
			if err != nil {
				glog.Errorln(err)
			} else {
				w.writes++
				glog.Errorln("write object success, size: ", len(b))
			}
			// select {
			// case <-time.After(time.Millisecond * 10): //wait 10ms for write perform
			// 	w.last = m
			// 	glog.Errorln("write object pending, size: ", len(b))
			// 	return
			// }
		default:
			return
		}
	}

}
