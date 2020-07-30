package generate

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
)

const (
	defaultPath    string = "./repository/"
	defaultPackage string = "repository"
)

type structData struct {
	Name    string
	VarName string
	Package string
}

func (g *generate) Write(filePath ...string) {
	var path string
	if len(filePath) == 0 {
		path = defaultPath
	} else {
		path = filePath[0]
	}

	checkDir(path)
	var err error
	for _, v := range g.structName {
		str := structData{
			Name:    v,
			VarName: strings.ToLower(v),
			Package: defaultPackage,
		}
		err = createFile(path, str)
		if err != nil {
			break
		}
	}
	if err != nil {
		fmt.Println(err)
	}
}

func checkDir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.Mkdir(path, 0755)
	}
}

func createFile(path string, sd structData) error {
	f, err := os.Create(path + sd.Name + ".go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bodyData, err := tempFile(sd)
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(bodyData)
	if err != nil {
		fmt.Println(err)
	}
	f.Sync()
	return nil
}

func tempFile(data interface{}) (string, error) {
	t := template.New("")
	tp := template.Must(t.Parse(templFile))
	buffer := new(bytes.Buffer)
	if err := tp.Execute(buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

const templFile string = `package {{.Package}}

import (
	"github.com/globalsign/mgo"
)

type {{.Name}}RepositoryInterface interface {
	Create{{.Name}}({{.VarName}} *{{.Name}}) error
	Get{{.Name}}(params map[string]interface{}) ({{.Name}}, error)
	Update{{.Name}}(params map[string]interface{}, {{.VarName}} *{{.Name}}) error
	Delete{{.Name}}(params map[string]interface{}) error
}

type {{.VarName}}Repository struct {
	db             *mgo.Session
	dbName         string
	collectionName string
}

func New{{.Name}}Repository(db *mgo.Session) {{.Name}}RepositoryInterface {
	return {{.VarName}}Repository{db, "db_name", "collection_name"}
}

func (conn {{.VarName}}Repository) Create{{.Name}}({{.VarName}} *{{.Name}}) error {
	err := conn.db.DB(conn.dbName).C(conn.collectionName).Insert({{.VarName}})
	return err
}

func (conn {{.VarName}}Repository) Get{{.Name}}(params map[string]interface{}) ({{.Name}}, error) {
	var {{.VarName}} {{.Name}}
	err := conn.db.DB(conn.dbName).C(conn.collectionName).Find(params).One(&{{.VarName}})
	return {{.VarName}}, err
}

func (conn {{.VarName}}Repository) Update{{.Name}}(params map[string]interface{}, {{.VarName}} *{{.Name}}) error {
	err := conn.db.DB(conn.dbName).C(conn.collectionName).Update(params, &{{.VarName}})
	return err
}

func (conn {{.VarName}}Repository) Delete{{.Name}}(params map[string]interface{}) error {
	err := conn.db.DB(conn.dbName).C(conn.collectionName).Remove(params)
	return err
}
`
