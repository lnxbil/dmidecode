package dmidecode

import (
	"io/ioutil"
	"testing"
)

func testBiosInformation(t *testing.T, dmi *DMI, version string, releaseDate string, vendor string, firmwareRevision string) {

	if dmi.BIOS.Version != version {
		t.Errorf("should return bios version '%s', but got '%s'", version, dmi.BIOS.Version)
	}
	if dmi.BIOS.ReleaseDate != releaseDate {
		t.Errorf("should return release date '%s', but got '%s", releaseDate, dmi.BIOS.ReleaseDate)
	}
	if dmi.BIOS.Vendor != vendor {
		t.Errorf("should return vendor '%s', but got '%s'", vendor, dmi.BIOS.Vendor)
	}
	if dmi.BIOS.FirmwareRevision != firmwareRevision {
		t.Errorf("should return firmware revsion '%s', but got '%s'", firmwareRevision, dmi.BIOS.FirmwareRevision)
	}
}

func testSystemInformation(t *testing.T, dmi *DMI, manufacturer string, productName string, serialNumber string, uuid string, name string) {
	if dmi.System.Manufacturer != manufacturer {
		t.Errorf("should return manufacturer '%s', but got '%s'", manufacturer, dmi.System.Manufacturer)
	}
	if dmi.System.ProductName != productName {
		t.Errorf("should return product name '%s', but got '%s'", productName, dmi.System.ProductName)
	}
	if dmi.System.SerialNumber != serialNumber {
		t.Errorf("should return serial number '%s', but got '%s", serialNumber, dmi.System.SerialNumber)
	}
	if dmi.System.UUID != uuid {
		t.Errorf("should return UUID '%s', but got '%s'", uuid, dmi.System.UUID)
	}
	if dmi.System.Name != name {
		t.Errorf("should return name '%s', but got '%s'", name, dmi.System.Name)
	}
}

func testAndReadFile(t *testing.T, file string) *DMI {
	dmi := New()
	data, readErr := ioutil.ReadFile(file)
	if readErr != nil {
		t.Errorf("Should not receive errors while reading contents of '%v'. Error: %v", file, readErr)
	}
	if err := dmi.ParseDmidecode(string(data)); err != nil {
		t.Errorf("Should not get errors while parsing '%v'. Error: %v", file, err)
	}

	return dmi
}

func TestDebian8HP(t *testing.T) {
	dmi := testAndReadFile(t, "test_data/debian_8_64bit_hp.txt")
	testBiosInformation(t, dmi, "P62", "01/22/2015", "HP", "2.29")
	testSystemInformation(t, dmi, "HP", "ProLiant DL380 G6", "CZC0000000", "11111111-2222-3333-4444-555555555555", "HP ProLiant DL380 G6")
}

func TestCentOSDell(t *testing.T) {
	dmi := testAndReadFile(t, "test_data/centos_6.5_64bit_dell.txt")
	testBiosInformation(t, dmi, "1.12.0", "07/26/2013", "Dell Inc.", "1.12")
	testSystemInformation(t, dmi, "Dell Inc.", "PowerEdge R510", "23V2JN1", "4C4C4544-0033-5610-8032-B2C04F4A4E31", "Dell Inc. PowerEdge R510")
}
func TestCentOSOEM(t *testing.T) {
	dmi := testAndReadFile(t, "test_data/centos_6.5_64bit_oem.txt")
	testBiosInformation(t, dmi, "P1.30", "10/27/2006", "American Megatrends Inc.", "8.12")
	testSystemInformation(t, dmi, "To Be Filled By O.E.M.", "To Be Filled By O.E.M.", "To Be Filled By O.E.M.", "00020003-0004-0005-0006-000700080009", "To Be Filled By O.E.M.")
}
func TestUbuntuVMware(t *testing.T) {
	dmi := testAndReadFile(t, "test_data/ubuntu_14.04_LTS_64bit_vmware.txt")
	testBiosInformation(t, dmi, "VirtualBox", "12/01/2006", "innotek GmbH", "")
	testSystemInformation(t, dmi, "innotek GmbH", "VirtualBox", "0", "F548DD5F-057D-4F7F-9465-FC529E045C08", "innotek GmbH VirtualBox")
}

//func TestTemplate(t *testing.T) {
//	dmi := testAndReadFile(t, "test_data/")
//	testBiosInformation(t, dmi, "", "", "", "")
//	testSystemInformation(t, dmi, "", "", "", "", "")
//}
