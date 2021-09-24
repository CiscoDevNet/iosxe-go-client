package client

import (
	"log"
	"os"
	"testing"
)

var (
	host           = os.Getenv("HOST_IOSXE")
	deviceUsername = os.Getenv("DEVICE_USERNAME_IOSXE")
	devicePassword = os.Getenv("DEVICE_PASSWORD_IOSXE")
)

func init() {
	if host == "" {
		log.Fatal("Cisco IOS XE Host for  tests is missing, please set in TEST_API_HOST env var")
	}
	if deviceUsername == "" || devicePassword == "" {
		log.Fatal("Credentials for client tests is missing, " +
			"please set in [DEVICE_USERNAME_IOSXE, DEVICE_PASSWORD_IOSXE] env var")
	}
}

func GetTestClient() (*V2, error) {
	return NewV2(
		host,
		deviceUsername,
		devicePassword,
		30,
		true,
		"",
		"",
		"",
	)
}

func TestClient(t *testing.T) {
	_, err := GetTestClient()
	if err != nil {
		t.Errorf("Error in creating client, %v", err)
	}
}

func TestGet(t *testing.T) {
	testClient, err := GetTestClient()
	if err != nil {
		t.Error("Error in creating client")
		return
	}
	_, _, err = testClient.Get("/data/Cisco-IOS-XE-native:native", nil)
	if err != nil {
		t.Errorf("Error in fetching data, %v", err)
	}
}
