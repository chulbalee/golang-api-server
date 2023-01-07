package entity

func (Tb_co_log) TableName() string {
	return "tb_co_log"
}

type Tb_co_log struct {
	BasDt string
	Id    string
	Msg   string
}
