package sms_provider

import (
	"log"
	"fmt"
	"github.com/supabase/auth/internal/conf"
)

type FakeProvider struct {
	Config  *conf.FakeProviderConfiguration
}

// Creates a SmsProvider with the Fake Config
func NewFakeProvider(config conf.FakeProviderConfiguration) (SmsProvider, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &FakeProvider{
		Config:  &config,
	}, nil
}

func (t *FakeProvider) SendMessage(phone, message, channel, otp string) (string, error) {
	switch channel {
	case SMSProvider:
		return t.SendSms(phone, message)
	default:
		return "", fmt.Errorf("channel type %q is not supported for Fake", channel)
	}
}

// Pretend to send an sms
func (t *FakeProvider) SendSms(phone string, message string) (string, error) {
	log.Printf("Sending message: %s", message)
	return message, nil
}
