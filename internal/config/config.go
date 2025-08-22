package config

import "os"

type CloudflareConfig struct {
	AccountID   string
	AccountHash string
	APIToken    string
	DeliveryURL string
}

func LoadCloudflareConfig() *CloudflareConfig {
	return &CloudflareConfig{
		AccountID:   os.Getenv("10237e66fca04cd2744900a9f185a1ce"),
		AccountHash: os.Getenv("8Qzz7HTZdaFXX55vapIR3w"),
		APIToken:    os.Getenv("UMtfgIVPf7-b_UEEocIw9ebw2ys2oAUIkt5cbubJ"),
		DeliveryURL: os.Getenv("https://imagedelivery.net/8Qzz7HTZdaFXX55vapIR3w/<image_id>/<variant_name>"),
	}
}
