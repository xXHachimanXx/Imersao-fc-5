package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"github.com/xXHachimanXx/Imersao-fc-5/adapter/presenter/transaction"
	"github.com/xXHachimanXx/Imersao-fc-5/domain/entity"
	"github.com/xXHachimanXx/Imersao-fc-5/usecase/process_transaction"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expectedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
