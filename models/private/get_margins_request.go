// Code generated by make generate-models DO NOT EDIT.

package private

type GetMarginsRequest struct {
	Amount         int64  `json:"amount" mapstructure:"amount"`
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
	Price          int64  `json:"price" mapstructure:"price"`
}
