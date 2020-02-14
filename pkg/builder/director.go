package builder

type Director struct {
	Builder ComputerBuilder
}

func (d *Director) ConstructPro() *Computer {
	d.Builder.InitComputer()
	d.Builder.SetType("Pro")
	d.Builder.SetMotherBoard("Pro")
	d.Builder.SetCPU("Intel Xeon W12")
	d.Builder.SetHardDrive("1 TB")
	d.Builder.SetRam("96 GB")
	d.Builder.SetGPU("AMD Radeon Pro 580X")
	d.Builder.SetOS("Pro")
	return d.Builder.Build()
}

func (d *Director) ConstructAir() *Computer {
	d.Builder.InitComputer()
	d.Builder.SetType("Air")
	d.Builder.SetMotherBoard("Air")
	d.Builder.SetCPU("Intel Core i5")
	d.Builder.SetGPU("Nvidia Geforce")
	d.Builder.SetHardDrive("500 GB")
	d.Builder.SetRam("16 GB")
	d.Builder.SetOS("Lite")
	return d.Builder.Build()
}
