/*
 * Copyright © 2019 Hedzr Yeh.
 */

package cmdr_test

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/hedzr/cmdr"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

type (
	tomlConfig struct {
		Title   string
		Runmode string
		App     app
		Owner   ownerInfo
		DB      database `toml:"database"`
		Servers map[string]server
		Clients clients
	}

	app struct {
		Debug bool
	}

	ownerInfo struct {
		Name string
		Org  string `toml:"organization"`
		Bio  string
		DOB  time.Time
	}

	database struct {
		Server  string
		Ports   []int
		ConnMax int `toml:"connection_max"`
		Enabled bool
	}

	server struct {
		IP string
		DC string
	}

	clients struct {
		Data  [][]interface{}
		Hosts []string
	}
)

func TestHasParent(t *testing.T) {
	s := cmdr.BaseOpt{
		Name:  "A",
		Short: "A",
		Full:  "Abcuse",
	}
	if s.HasParent() {
		t.Failed()
	}
	if s.GetTitleNames() != "A, Abcuse" {
		t.Failed()
	}
}

func TestTomlLoad(t *testing.T) {
	var (
		err    error
		b      []byte
		mm     map[string]map[string]interface{}
		config tomlConfig
		meta   toml.MetaData
	)

	b = []byte(`

runmode="devel"

title = "TOML Example"

[app]
debug=true

[owner]
name = "Tom Preston-Werner"
organization = "GitHub"
bio = "GitHub Cofounder & CEO\nLikes tater tots and beer."
dob = 1979-05-27T07:32:00Z # First class dates? Why not?

[database]
server = "192.168.1.1"
ports = [ 8001, 8001, 8002 ]
connection_max = 5000
enabled = true

[servers]

  # You can indent as you please. Tabs or spaces. TOML don't care.
  [servers.alpha]
  ip = "10.0.0.1"
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"

[clients]
data = [ ["gamma", "delta"], [1, 2] ] 
# just an update to make sure parsers support it

# Line breaks are OK when inside arrays
hosts = [
  "alpha",
  "omega"
]

`)

	if err = ioutil.WriteFile(".tmp.toml", b, 0644); err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = os.Remove(".tmp.toml")
		_ = os.Remove(".tmp.2.toml")
	}()

	mm = make(map[string]map[string]interface{})
	if err = toml.Unmarshal(b, &mm); err != nil {
		return
	}

	t.Log(mm)

	if meta, err = toml.DecodeFile(".tmp.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	t.Log(config)
	t.Log(meta)

	if err = cmdr.SaveObjAsToml(config, ".tmp.2.toml"); err != nil {
		t.Fatal(err)
	}

	// if err = cmdr.LoadConfigFile(".tmp.toml"); err != nil {
	// 	t.Fatal(err)
	// }

}

func TestConfigFiles(t *testing.T) {
	var err error

	defer func() {
		_ = os.Remove(".tmp.json")
		_ = os.Remove(".tmp.yaml")
		_ = os.Remove(".tmp.toml")
	}()

	// try loading cfg again for gocov
	if err = cmdr.LoadConfigFile(".tmp.yaml"); err != nil {
		t.Fatal(err)
	}
	_ = os.Remove(".tmp.yaml")

	// try loading cfg again for gocov
	if err = cmdr.LoadConfigFile(".tmp.yaml"); err != nil {
		t.Fatal(err)
	}

	_ = ioutil.WriteFile(".tmp.yaml", []byte(`
app'x':"
`), 0644)
	// try loading cfg again for gocov
	if err = cmdr.LoadConfigFile(".tmp.yaml"); err == nil {
		t.Fatal("loading cfg file should be failed (err != nil), but it returns nil as err.")
	}
	_ = os.Remove(".tmp.yaml")

	_ = ioutil.WriteFile(".tmp.json", []byte(`{"app":{"debug":false}}`), 0644)
	// try loading cfg again for gocov
	if err = cmdr.LoadConfigFile(".tmp.json"); err != nil {
		t.Fatal(err)
	}
	_ = os.Remove(".tmp.json")

	_ = ioutil.WriteFile(".tmp.toml", []byte(`
runmode="devel"
[app]
debug=true
`), 0644)
	// try loading cfg again for gocov
	if err = cmdr.LoadConfigFile(".tmp.toml"); err != nil {
		t.Fatal(err)
	}
	_ = os.Remove(".tmp.toml")
}

func TestLaunchEditor2(t *testing.T) {
	if b, err := cmdr.LaunchEditorWith("cat", "/etc/passwd"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(string(b))
	}

	_, _ = cmdr.LaunchEditorWith("cat", "/etc/not-exists")
}

func TestClone1(t *testing.T) {
	var aa = "dsajkld"
	var b int

	// cmdr.Clone(b, aa)

	cmdr.Clone(b, &aa)

	cmdr.Clone(&b, &aa)

	cmdr.Clone(&b, nil)

	var c, d bool
	cmdr.Clone(&c, &d)

	var e, f int
	cmdr.Clone(&e, &f)
	var e1, f1 int8
	f1 = 1
	cmdr.Clone(&e1, &f1)
	if e1 != 1 {
		t.Failed()
	}
	var e2, f2 int16
	cmdr.Clone(&e2, &f2)
	var e3, f3 int32
	cmdr.Clone(&e3, &f3)
	var e4, f4 int64
	f4 = 9
	cmdr.Clone(&e4, &f4)
	if e1 != 9 {
		t.Failed()
	}

	var g, h string
	cmdr.Clone(&g, &h)
}

func TestClone2(t *testing.T) {

}