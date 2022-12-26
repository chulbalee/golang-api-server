package common

import "golang-api-server/kafka"

type TbCoLogSvcImpl struct {
	kafka kafka.Kafka
}

const (
	TOPIC = "tb_co_log"
)

func (t *TbCoLogSvcImpl) Init(clienId string, bootstrapServers []string) {
	t.kafka.Init(clienId, bootstrapServers, TOPIC)
}

func (t *TbCoLogSvcImpl) InsertLog(msg string) {
	t.kafka.Produce(msg)
}
