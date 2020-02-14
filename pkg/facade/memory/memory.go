package memory

type Memory interface {
	Read(position int) byte
	Write(position int, data byte)
	Free(position int)
	CleanAll()
	GetSize() int
}
