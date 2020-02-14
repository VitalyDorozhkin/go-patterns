package builder

import "fmt"

type Computer struct {
	Builder     string
	Type        string
	MotherBoard string
	CPU         string
	HardDrive   string
	Ram         string
	GPU         string
	OS          string
}

// Show returns product.
func (c *Computer) Show() string {
	return fmt.Sprintf("Builder: %s\n"+
		"Type: %s\n"+
		"MotherBoard: %s\n"+
		"CPU: %s\n"+
		"HardDrive: %s\n"+
		"Ram: %s\n"+
		"GPU: %s\n"+
		"OS: %s\n", c.Builder, c.Type, c.MotherBoard, c.CPU, c.HardDrive, c.Ram, c.GPU, c.OS)
}
