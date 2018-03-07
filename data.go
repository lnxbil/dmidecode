package dmidecode

import (
	"fmt"
	"strings"
)

// BiosInformation represents an excerpt from DMI type 0
type BiosInformation struct {
	FirmwareRevision string
	ReleaseDate      string
	Vendor           string
	Version          string
}

// SystemInformation represents an excerpt from DMI type 1
type SystemInformation struct {
	Manufacturer string
	ProductName  string
	SerialNumber string
	UUID         string
	Family       string

	// combined
	Name string // consists of Manufacturer + ProductName
}

// ProcessorInformation represents an excerpt from DMI type 4
//type ProcessorInformation struct {
//	Family       string
//	Manufacturer string
//	Version      string
//	Voltage      string
//	CurrentSpeed string
//	MaxSpeed     string
//	Status       string // populated or not
//	CoreCount    int
//	ThreadCount  int
//}

// MemoryDevice represents an excerpt from DMI type 17
//type MemoryDevice struct {
//	Size                 string
//	Locator              string
//	Type                 string
//	Speed                string
//	Manufacturer         string
//	SerialNumber         string
//	ConfiguredClockSpeed string
//}

func (d *DMI) parseToStructs() error {

	// retrieve bios information from dmi type 0
	byNameData, byNameErr := d.SearchByType(0)
	if byNameErr == nil {
		d.BIOS.FirmwareRevision = strings.TrimSpace(byNameData["Firmware Revision"])
		if d.BIOS.FirmwareRevision == "" {
			d.BIOS.FirmwareRevision = strings.TrimSpace(byNameData["BIOS Revision"])
		}
		d.BIOS.Version = strings.TrimSpace(byNameData["Version"])
		d.BIOS.ReleaseDate = strings.TrimSpace(byNameData["Release Date"])
		d.BIOS.Vendor = strings.TrimSpace(byNameData["Vendor"])
	}

	// retrieve system information from dmi type 1
	byNameData, byNameErr = d.SearchByType(1)
	if byNameErr == nil {
		d.System.Family = strings.TrimSpace(byNameData["Family"])
		d.System.Manufacturer = strings.TrimSpace(byNameData["Manufacturer"])
		d.System.ProductName = strings.TrimSpace(byNameData["Product Name"])
		d.System.SerialNumber = strings.TrimSpace(byNameData["Serial Number"])
		d.System.UUID = strings.TrimSpace(byNameData["UUID"])

		// consturct string and remove any leading/trailing and double space
		if d.System.ProductName == "To Be Filled By O.E.M." && d.System.Manufacturer == "To Be Filled By O.E.M." {
			d.System.Name = "To Be Filled By O.E.M."
		} else {
			d.System.Name = strings.Join(strings.Fields(strings.TrimSpace(fmt.Sprintf("%s %s", d.System.Manufacturer, d.System.ProductName))), " ")
		}
	}

	// retrieve processors from dmi type 4
	// byNameData, byNameErr = d.SearchByType(4)
	// if byNameErr == nil {
	// 	n := len(byNameData)
	// 	d.Processor = make([]ProcessorInformation, n)
	// 	for index, p := range byNameData {
	// 		d.Processor[index].CoreCount, _ = strconv.Atoi(p["Core Count"])
	// 	}
	// }

	return nil
}
