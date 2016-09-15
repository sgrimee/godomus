package godomus

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/viper"
)

var (
	domus *Domus
)

// TODO: get this test data from config file
var (
	cfgFile        string
	testSiteKey    SiteKey
	testUserKey    UserKey
	testPassword   string
	testApiUrl     string
	testSocketPort int
	testSocketAddr string
	testRoomKey    RoomKey
)

func TestMain(m *testing.M) {
	initConfig()
	os.Exit(m.Run())
}

// initConfig reads in config file
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".ldclient") // name of config file (without extension)
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	//fmt.Println("Config:")
	//spew.Dump(viper.AllSettings())

	testSiteKey = NewSiteKey(viper.GetInt("site"))
	testUserKey = NewUserKey(viper.GetInt("user"))
	testPassword = viper.GetString("password")
	testApiUrl = viper.GetString("url")
	testSocketPort = viper.GetInt("socket_port")
	testSocketAddr = viper.GetString("test_socket_addr")
	testRoomKey = NewRoomKey(viper.GetInt("test_room"))
}

func TestKeyConversions(t *testing.T) {
	sk := SiteKey("SITE_00000000000000000000000000000000019")
	if sk != NewSiteKey(19) {
		t.Error("NewSiteKey failed.")
	}
	if 19 != sk.Num() {
		t.Error("SiteKey.Num failed.")
	}
	uk := UserKey("USER_00000000000000000000000000000000012")
	if uk != NewUserKey(12) {
		t.Error("NewUserKey failed.")
	}
	if 12 != uk.Num() {
		t.Error("UserKey.Num failed.")
	}
}

func TestNew(t *testing.T) {
	d, err := New(testApiUrl, testSocketPort)
	if err != nil {
		t.Fatalf("Could not create domus object: %s", err)
	}
	if d.socketAddr != testSocketAddr {
		t.Fatalf("Incorrect socket url: %s", d.socketAddr)
	}
	domus = d
}

func TestGetSites(t *testing.T) {
	sites, err := domus.GetSites()
	if err != nil {
		t.Fatal(err)
	}
	if len(sites) < 1 {
		t.Fatal("No sites returned")
	}
}

func TestGetUsers(t *testing.T) {
	users, err := domus.GetUsers(testSiteKey)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) < 1 {
		t.Fatal("No users returned")
	}
}

func TestLoginInfos(t *testing.T) {
	infos, err := domus.LoginInfos(testSiteKey, testUserKey, testPassword)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Login infos: %+v\n", infos)
	if len(infos.Rooms.Room) <= 0 {
		t.Fatal("No rooms received")
	}
	if len(infos.Scenarios.Scenario) <= 0 {
		t.Fatal("No scenarios received")
	}
}

func TestLogin(t *testing.T) {
	session, err := domus.Login(testSiteKey, testUserKey, testPassword)
	if err != nil {
		t.Fatal(err)
	}
	if len(session) != sessionKeyLen {
		t.Fatal("Invalid session key length")
	}
	if domus.sessionKey != session {
		t.Fatalf("Session key was not saved correctly, saved:%s, received:%s", domus.sessionKey, session)
	}
}

func TestDevicesInRoom(t *testing.T) {
	devices, err := domus.DevicesInRoom(testRoomKey, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(devices) <= 0 {
		t.Fatal("No devices received.")
	}
	t.Logf("Devices:\n%+v\n", devices)
}

func TestListenForEvents(t *testing.T) {
	t.Skip("skipping listening test")
	err := domus.ListenForEvents(10)
	if err != nil {
		t.Fatal(err)
	}
}
