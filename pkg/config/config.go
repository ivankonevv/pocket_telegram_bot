package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	AuthServerURL     string
	TelegramBotURL    string `mapstucture:"bot_url"`
	DBPath            string `mapstucture:"db_file"`

	Messages Messages
}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default      string `mapstucture:"default"`
	InvalidURL   string `mapstucture:"invalid_url"`
	Unauthorized string `mapstucture:"unauthorized"`
	UnableToSave string `mapstucture:"unable_to_save"`
}

type Responses struct {
	Start             string `mapstucture:"start"`
	AlreadyAuthorized string `mapstucture:"alredy_authorized"`
	SavedSuccessfully string `mapstucture:"saved_successfully"`
	UnknownCommand    string `mapstucture:"unknown_command"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := ParseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ParseEnv(cfg *Config) error {
	os.Setenv("TOKEN", "1524930274:AAFxKfwe-zHo87cVC_I3WN3h5ooVm4wJKR0")
	os.Setenv("CONSUMER_KEY", "95960-a3af6bbd52ed0e3638ecb8b1")
	os.Setenv("AUTH_SERVER_URL", "http://localhost/")
	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}

	if err := viper.BindEnv("auth_server_url"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("token")
	cfg.PocketConsumerKey = viper.GetString("consumer_key")
	cfg.AuthServerURL = viper.GetString("auth_server_url")

	return nil
}