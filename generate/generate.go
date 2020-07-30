package generate

type GenerateInterface interface {
	ReadFile() error
	Write(filePath ...string)
}

type generate struct {
	path       string
	structName []string
}

func NewFilePath(filepath string) GenerateInterface {
	return &generate{path: filepath}
}

func (g *generate) SetStructName(structName []string) {
	g.structName = structName
}
