package builder

type Director struct {
	builder ComputerBuilder
}

func (d *Director) ConstructPro() *Computer {
	d.builder.InitComputer()
	d.builder.SetType("Pro")
	d.builder.SetMotherBoard("Pro")
	d.builder.SetCPU("Intel Xeon W12")
	d.builder.SetHardDrive("1 TB")
	d.builder.SetRam("96 GB")
	d.builder.SetGPU("AMD Radeon Pro 580X")
	d.builder.SetOS("Pro")
	return d.builder.Build()
}

func (d *Director) ConstructAir() *Computer {
	d.builder.InitComputer()
	d.builder.SetType("Air")
	d.builder.SetMotherBoard("Air")
	d.builder.SetCPU("Intel Core i5")
	d.builder.SetGPU("Nvidia Geforce")
	d.builder.SetHardDrive("500 GB")
	d.builder.SetRam("16 GB")
	d.builder.SetOS("Lite")
	return d.builder.Build()
}
