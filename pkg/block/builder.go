pacakge block

type BuilderV1 struct{
	w io.Writer
	entryCount int
	size int64
}


func (b *BuilderV1)Add(e *Entry){
}