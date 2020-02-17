package datastorage

type MemoryEditor interface {
	Read(position int) byte
	Write(position int, data byte)
}
