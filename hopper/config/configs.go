package config

type KafkaConfig struct {
	Hosts              []string
	FormatFaceTopic    string
	SendMessageTimeout uint
}

type FTPConfig struct {
	Address   string
	Path      string
	Username  string
	Password  string
	MaxThread int
}

type Config struct {
	KafkaInputConfig  *KafkaConfig
	KafkaOutputConfig *KafkaConfig
	FTPInputConfig    *FTPConfig
	FTPOutputConfig   *FTPConfig
	SpeedLimit        int
	HopperOut         bool
	HopperIn          bool
	Devices           map[string]string
}
