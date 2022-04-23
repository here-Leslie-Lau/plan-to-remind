package data

type EmailConfig struct {
	Subject string   `json:"subject"`
	Msg     string   `json:"msg"`
	To      []string `json:"to"`
}

type Config struct {
	EmailConfig EmailConfig `json:"email_config"`
}

type ConfigRepo struct {
	data *Data
}

func (repo *ConfigRepo) GetConfig() {

}
