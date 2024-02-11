package interfaces

import "time"

type ResponseCreateTransaction struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}

type RequestCreateTransaction struct {
	ID        int64  `json:"id,omitempty"`
	Valor     int64  `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type UltimasTransacoes struct {
	Valor       int64     `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type ResponseGetExtrato struct {
	Saldo struct {
		Total       int64  `json:"total"`
		DataExtrato string `json:"data_extrato"`
		Limite      int64  `json:"limite"`
	} `json:"saldo"`
	UltimasTransacoes []UltimasTransacoes `json:"ultimas_transacoes"`
}

type ExtratoFromDB struct {
	IDCliente     int64     `bson:"id_cliente"`
	Saldo         int64     `bson:"saldo"`
	Limite        int64     `bson:"limit"`
	Valor         *int64    `bson:"valor,omitempty"`
	TipoTransacao *string   `bson:"tipo_transacao,omitempty"`
	Descricao     *string   `bson:"description,omitempty"`
	CreatedAt     time.Time `bson:"created_at,omitempty"`
}
