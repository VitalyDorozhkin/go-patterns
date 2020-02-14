package datastorage

type Memory interface {
	Read(position int) byte
	Write(position int, data byte)
}
