package common

import (
	"golang-api-server/entity"
	"golang-api-server/kafka"
	"time"
)

type TbCoLogSvcImpl struct {
	kafka kafka.Kafka
}

const (
	TOPIC = "TB_CO_LOG_HIST"
)

func (t *TbCoLogSvcImpl) Init(clienId string, bootstrapServers []string) {
	t.kafka.Init(clienId, bootstrapServers, TOPIC)
}

func (t *TbCoLogSvcImpl) InsertJson(data entity.Tb_co_log) {
	t.kafka.Produce(data)
}

func (t *TbCoLogSvcImpl) InsertMsg(msg string) {
	logJson := entity.Tb_co_log{}
	logJson.BasDt = time.Now().Format("20060102")
	logJson.Msg = msg

	t.kafka.Produce(logJson)
}
