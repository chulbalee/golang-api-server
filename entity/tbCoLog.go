package entity

func (Tb_co_log) TableName() string {
	return "tb_co_log"
}

type Tb_co_log struct {
	Id  int    `json:"id"`
	Tx  string `json:"tx"`
	Msg string `json:"msg"`
}
