package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {

	appleDirector := Director{&AppleBuilder{}}
	xiaomiDirector := Director{&XiaomiBuilder{}}

	miNotebookPro := xiaomiDirector.ConstructPro()
	miNotebookAir := xiaomiDirector.ConstructAir()

	macbookPro := appleDirector.ConstructPro()
	macbookAir := appleDirector.ConstructAir()

	cases := make(map[*Computer]string, 4)

	cases[miNotebookPro] =
		"Builder: Xiaomi\n" +
			"Type: Pro\n" +
			"MotherBoard: Xiaomi Pro MotherBoard\n" +
			"CPU: Intel Xeon W12\n" +
			"HardDrive: Samsung HardDrive on 1 TB\n" +
			"Ram: Samsung Ram on 96 GB\n" +
			"GPU: AMD Radeon Pro 580X\n" +
			"OS: Windows 10 Pro\n"
	cases[miNotebookAir] =
		"Builder: Xiaomi\n" +
			"Type: Air\n" +
			"MotherBoard: Xiaomi Air MotherBoard\n" +
			"CPU: Intel Core i5\n" +
			"HardDrive: Samsung HardDrive on 500 GB\n" +
			"Ram: Samsung Ram on 16 GB\n" +
			"GPU: Nvidia Geforce\n" +
			"OS: Windows 10 Lite\n"
	cases[macbookPro] =
		"Builder: Apple\n" +
			"Type: Pro\n" +
			"MotherBoard: Apple Pro MotherBoard\n" +
			"CPU: Intel Xeon W12\n" +
			"HardDrive: Apple HardDrive on 1 TB\n" +
			"Ram: Apple Ram on 96 GB\n" +
			"GPU: AMD Radeon Pro 580X\n" +
			"OS: MacOS Catalina Pro\n"
	cases[macbookAir] =
		"Builder: Apple\n" +
			"Type: Air\n" +
			"MotherBoard: Apple Air MotherBoard\n" +
			"CPU: Intel Core i5\n" +
			"HardDrive: Apple HardDrive on 500 GB\n" +
			"Ram: Apple Ram on 16 GB\n" +
			"GPU: Nvidia Geforce\n" +
			"OS: MacOS Catalina Lite\n"

	for k, v := range cases {
		result := k.Show()
		if result != v {
			t.Errorf("Expect:\n%s\nBut result is:\n%s\n", v, result)
		}
	}

}
