// Code generated by make generate-models DO NOT EDIT.

package private

type GetAccountSummaryResponse struct {
	AvailableFunds            float64 `json:"available_funds" mapstructure:"available_funds"`
	AvailableWithdrawalFunds  float64 `json:"available_withdrawal_funds" mapstructure:"available_withdrawal_funds"`
	Balance                   float64 `json:"balance" mapstructure:"balance"`
	Currency                  string  `json:"currency" mapstructure:"currency"`
	DeltaTotal                float64 `json:"delta_total" mapstructure:"delta_total"`
	DepositAddress            string  `json:"deposit_address" mapstructure:"deposit_address"`
	Email                     string  `json:"email" mapstructure:"email"`
	Equity                    float64 `json:"equity" mapstructure:"equity"`
	FuturesPl                 int64   `json:"futures_pl" mapstructure:"futures_pl"`
	FuturesSessionRpl         int64   `json:"futures_session_rpl" mapstructure:"futures_session_rpl"`
	FuturesSessionUpl         int64   `json:"futures_session_upl" mapstructure:"futures_session_upl"`
	ID                        int64   `json:"id" mapstructure:"id"`
	InitialMargin             float64 `json:"initial_margin" mapstructure:"initial_margin"`
	MaintenanceMargin         float64 `json:"maintenance_margin" mapstructure:"maintenance_margin"`
	MarginBalance             float64 `json:"margin_balance" mapstructure:"margin_balance"`
	OptionsDelta              int64   `json:"options_delta" mapstructure:"options_delta"`
	OptionsGamma              int64   `json:"options_gamma" mapstructure:"options_gamma"`
	OptionsPl                 int64   `json:"options_pl" mapstructure:"options_pl"`
	OptionsSessionRpl         int64   `json:"options_session_rpl" mapstructure:"options_session_rpl"`
	OptionsSessionUpl         int64   `json:"options_session_upl" mapstructure:"options_session_upl"`
	OptionsTheta              int64   `json:"options_theta" mapstructure:"options_theta"`
	OptionsVega               int64   `json:"options_vega" mapstructure:"options_vega"`
	PortfolioMarginingEnabled bool    `json:"portfolio_margining_enabled" mapstructure:"portfolio_margining_enabled"`
	SessionFunding            int64   `json:"session_funding" mapstructure:"session_funding"`
	SessionRpl                int64   `json:"session_rpl" mapstructure:"session_rpl"`
	SessionUpl                int64   `json:"session_upl" mapstructure:"session_upl"`
	SystemName                string  `json:"system_name" mapstructure:"system_name"`
	TfaEnabled                bool    `json:"tfa_enabled" mapstructure:"tfa_enabled"`
	TotalPl                   int64   `json:"total_pl" mapstructure:"total_pl"`
	Type                      string  `json:"type" mapstructure:"type"`
	Username                  string  `json:"username" mapstructure:"username"`
}
