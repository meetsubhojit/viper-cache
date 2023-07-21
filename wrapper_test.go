package viper_cache

import (
	"testing"

	"github.com/spf13/viper"
)

func TestViperGet(t *testing.T) {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatalf("err in TestViperGet, viper.ReadInConfig failed, err: %s", err.Error())
	}
	viperFetchesAndChecks(t)
}

func viperFetchesAndChecks(t *testing.T) {
	if GetString("some_struct.first_field") != "some big string" {
		t.Fatalf("some_struct.first_field should be 'some big string', got %v", GetString("some_struct.first_field"))
	}
	if GetFloat64("some_struct.second_field") != 100 {
		t.Fatalf("some_struct.second_field should be 100, got %v", GetInt("some_struct.second_field"))
	}
	if !GetBool("some_struct.third_field") {
		t.Fatalf("some_struct.third_field should be true, got %v", GetBool("some_struct.third_field"))
	}

}
