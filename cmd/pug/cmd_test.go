package main

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"text/template"
)

const mainTemplate = `package main

import (
	"fmt"
	"strings"

	"github.com/shaban/pug"
)

func main() {
	sb := new(strings.Builder)
	pug_{{.}}(sb)
	sr := strings.NewReader(sb.String())
	sb.Reset()
	pug.Format(sr, sb)
	out := strings.TrimSpace(sb.String())
	fmt.Print(out)
}`

func TestTemplateLanguage(t *testing.T) {
	var (
		enteredTest           = false
		wd, _                 = os.Getwd()
		pugFiles, resultFiles = readAllTemplateTestFileTuples()
		tpl                   *template.Template
		f                     *os.File
		err                   error
		cmd                   *exec.Cmd
		cmdOut                bytes.Buffer
	)
	//t.Logf("Input: %+v\nOutput: %+v\n", pugFiles, resultFiles)
	for name := range pugFiles {
		t.Run(name, func(t *testing.T) {
			path := pugFiles[name]
			dir := filepath.Dir(path)
			if !enteredTest {
				os.Chdir(dir)
				enteredTest = true
			}
			tpl = template.New("main.tpl")
			f, err = os.Create("main.go")
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			tpl, err = tpl.Parse(mainTemplate)
			if err != nil {
				t.Fatal(err)
			}
			err = tpl.Execute(f, name)
			if err != nil {
				t.Fatal(err)
			}
			cmd = exec.Command("pug", "-pkg", "main", "-writer", name+".pug")
			err = cmd.Run()
			if err != nil {
				t.Fatalf("exec.Command: %s\nErr: %s", name, err.Error())
			}
			cmd = exec.Command("go", "run", ".")
			cmd.Stdout = &cmdOut
			cmd.Run()
			os.Remove(name + ".pug.go")
			os.Remove("pug.go")
			os.Remove("main.go")
			if resultFiles[name] != cmdOut.String() {
				t.Errorf("HTML output mismatch\nExpected: %s\nGot: %s", resultFiles[name], cmdOut.String())
			} else {
				t.Logf("HTML output matched\nExpected: %s\nGot: %s", resultFiles[name], cmdOut.String())
			}
		})
	}
	os.Chdir(wd)
}
func readAllTemplateTestFileTuples() (map[string]string, map[string]string) {
	var (
		testFiles   = make(map[string]string)
		resultFiles = make(map[string]string)
		err         error
	)
	// Get the absolute path of the current package
	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current file path")
	}
	currentDir := filepath.Dir(currentFilePath)

	// Construct the absolute path to the testfiles directory
	testFilesDir := filepath.Join(currentDir, "/test_files")

	// Walk the testfiles directory
	err = filepath.Walk(testFilesDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filename := filepath.Base(path)
			if strings.HasSuffix(path, ".pug") {
				name := strings.TrimSuffix(filename, ".pug")
				testFiles[name] = path
			}
			if strings.HasSuffix(path, ".html") {
				name := strings.TrimSuffix(filename, ".html")
				expectResult, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				resultFiles[name] = string(expectResult)
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	return testFiles, resultFiles
}
