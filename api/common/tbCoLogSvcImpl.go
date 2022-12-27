package common

import (
	"golang-api-server/entity"
	"golang-api-server/kafka"
)

type TbCoLogSvcImpl struct {
	kafka kafka.Kafka
}

const (
	TOPIC = "tb_co_log"
)

func (t *TbCoLogSvcImpl) Init(clienId string, bootstrapServers []string) {
	t.kafka.Init(clienId, bootstrapServers, TOPIC)
}

func (t *TbCoLogSvcImpl) InsertJson(data entity.Tb_co_log) {
	t.kafka.Produce(data)
}

func (t *TbCoLogSvcImpl) InsertMsg(msg string) {
	logJson := entity.Tb_co_log{}
	logJson.Id = 1
	logJson.Tx = "aaa001"
	logJson.Msg = msg

	t.kafka.Produce(logJson)
}
