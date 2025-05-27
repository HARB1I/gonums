package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// EnumDef описывает один enum
type EnumDef struct {
	Name   string
	Values []string
}

// parseEnums парсит файл с enum'ами
func parseEnums(file string) []EnumDef {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var enums []EnumDef
	var current *EnumDef
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "enum ") && strings.HasSuffix(line, ":") {
			if current != nil {
				enums = append(enums, *current)
			}
			name := strings.TrimSuffix(strings.TrimSpace(line[5:]), ":")
			// Делаем первую букву заглавной
			if len(name) > 0 {
				name = strings.ToUpper(name[:1]) + name[1:]
			}
			current = &EnumDef{Name: name}
		} else if current != nil {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				current.Values = append(current.Values, fields[0])
			}
		}
	}

	if current != nil {
		enums = append(enums, *current)
	}

	return enums
}

// Шаблон для генерации enum
const enumTemplate = `package {{.Name | lower}}

// Implements:
//
{{- range .Values}}
//	{{$.Name | lower}}.{{.|title}}
{{- end}}
type Enum interface {
    is{{.Name}}()
	String() string
}
{{range .Values}}
type {{.|lower}} struct{}

func ({{.|lower}}) is{{$.Name}}() {}

func ({{.|lower}}) String() string { return "{{.|lower}}" }
{{end}}
var (
{{range .Values}}    {{.|title}} = {{.|lower}}{}
{{end}})
`

// Функции для преобразования имён
func toLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func toTitle(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// generateEnums создаёт файлы с enum'ами
func generateEnums(enums []EnumDef, baseDir string) {
	tmpl := template.Must(template.New("enum").Funcs(template.FuncMap{
		"lower": toLower,
		"title": toTitle,
	}).Parse(enumTemplate))

	for _, e := range enums {
		dir := filepath.Join(baseDir, strings.ToLower(e.Name))
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}

		filename := filepath.Join(dir, "enum.go")
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = tmpl.Execute(file, e)
		if err != nil {
			panic(err)
		}
		fmt.Printf("✅ Generated %s\n", filename)
	}
}

func main() {
	// Флаги
	inputFile := flag.String("input", "enums.def", "Путь к DSL-файлу с enum'ами")
	outputDir := flag.String("output", "enums", "Папка, куда будут генерироваться enum'ы")

	flag.Parse()

	enums := parseEnums(*inputFile)

	generateEnums(enums, *outputDir)
}
