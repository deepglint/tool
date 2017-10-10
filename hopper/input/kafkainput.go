package input

import (
	"github.com/deepglint/tool/hopper/config"

	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/golang/glog"
)

type KafkaInput struct {
	config        *config.KafkaConfig
	stopChan      chan bool
	resultChan    chan *[]byte
	isLive        bool
	sendMessageWg sync.WaitGroup
}

func NewKafkaInput(kafkaInputConfig *config.KafkaConfig) (*KafkaInput, error) {
	kafkaInput := new(KafkaInput)
	kafkaInput.config = kafkaInputConfig
	return kafkaInput, nil
}

func (this *KafkaInput) Start() (err error) {

	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true
	kafkaConfig.Consumer.Fetch.Max = 104857600

	kafkaConsumer, err := sarama.NewConsumer(this.config.Hosts, kafkaConfig)
	if err != nil {
		glog.Errorln("[INTPUT_KAFKA] start kafka input err:", err.Error())
		return err
	}
	partitionConsumer, err := kafkaConsumer.ConsumePartition(this.config.FormatFaceTopic, 0, sarama.OffsetNewest)
	if err != nil {
		glog.Errorln("[INTPUT_KAFKA] start kafka input err:", err.Error())
		return err
	}

	//open stop chan
	this.stopChan = make(chan bool)

	//open result chan
	this.resultChan = make(chan *[]byte)

	this.isLive = true
	glog.Infoln("[INTPUT_KAFKA] kafka input start ")
	go this.run(kafkaConsumer, partitionConsumer)

	return nil

}

func (this *KafkaInput) run(kafkaConsumer sarama.Consumer, partitionConsumer sarama.PartitionConsumer) {

	defer kafkaConsumer.Close()
	defer glog.Infoln("[INPUT_KAFKA] Kafka Input Closed ")
	for {
		select {
		case <-this.stopChan:
			// wait all send message thread close
			this.sendMessageWg.Wait()
			//close result chan
			close(this.resultChan)
			this.isLive = false
			return
		case err := <-partitionConsumer.Errors():
			glog.Warningln("[INPUT_KAFKA] kafka input err:", err)
			<-time.After(time.Second)
			continue
		case kafkaMessage := <-partitionConsumer.Messages():
			this.sendMessageWg.Add(1)
			glog.Errorln(string(kafkaMessage.Value))
			go func(recResult *[]byte) {
				defer this.sendMessageWg.Done()
				select {
				case <-time.After(time.Duration(this.config.SendMessageTimeout) * time.Second):
					glog.Infoln("[INPUT_KAFKA] send message to matrix transform timeout")
				case this.resultChan <- recResult:
					glog.Infoln("[INPUT_KAFKA] input send valid message to ftp output, message length is", len(*recResult))
				}
			}(&kafkaMessage.Value)

		}

	}
}

func (this *KafkaInput) Close() (err error) {

	if this.isLive {
		close(this.stopChan)
	}
	return nil
}

func (this *KafkaInput) GetMessageInfoChan() chan *[]byte {
	return this.resultChan
}

func (this *KafkaInput) GetLiveState() bool {
	return this.isLive
}
