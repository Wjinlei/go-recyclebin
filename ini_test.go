package recyclebin

import "testing"

func TestConfigStruct_Create(t *testing.T) {
	configStruct := NewIni()
	configObject, err := configStruct.Create("./config.ini")
	if err != nil {
		t.Fatal(err)
	}

	configObject.File.Section("").Key("Name").SetValue("JerryWang")

	if err := configObject.Save(); err != nil {
		t.Fatal(err)
	}
}
