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

// mainTemplate is a text template for generating a main file to execute
// the generated template function in {testname}.pug.go
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
		//enteredTest                       = false
		wd, _                             = os.Getwd()
		pugFiles, resultFiles, expectFail = readAllTemplateTestFileTuples()
		tpl                               *template.Template
		f                                 *os.File
		err                               error
		cmd                               *exec.Cmd
		cmdOut                            = new(bytes.Buffer)
	)
	//for debugging the tests
	/*t.Logf("Input: %+v\nOutput: %+v\n", pugFiles, resultFiles)
	for key := range pugFiles {
		t.Logf("Test: %s shouldFail: %t Result Length: %d\n", key, expectFail[key], len(resultFiles[key]))
	}*/
	for name := range pugFiles {
		t.Run(name, func(t *testing.T) {
			t.Log("Name: ", name)
			path := pugFiles[name]
			dir := filepath.Dir(path)
			os.Chdir(dir)
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
			cmdOut.Reset()
			cmd.Stdout = cmdOut
			cmd.Run()
			os.Remove(name + ".pug.go")
			os.Remove("pug.go")
			os.Remove("main.go")
			if resultFiles[name] != cmdOut.String() {
				if expectFail[name] == false {
					t.Errorf("Should Succeed but doesn't:\nHTML output mismatch\nExpected: %s\nGot: %s", resultFiles[name], cmdOut.String())
				}
			} else {
				if expectFail[name] == true {
					t.Errorf("Should Fail, but doesn't:\nHTML output matched\nExpected: %s\nGot: %s", resultFiles[name], cmdOut.String())
				}
			}
		})
	}
	os.Chdir(wd)
}
func readAllTemplateTestFileTuples() (map[string]string, map[string]string, map[string]bool) {
	var (
		testFiles   = make(map[string]string)
		resultFiles = make(map[string]string)
		expectFail  = make(map[string]bool)
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
			if strings.HasSuffix(path, ".pug") {
				dir, fname := filepath.Split(path)
				name := strings.TrimSuffix(fname, ".pug")
				testFiles[name] = path
				expectResult, err := os.ReadFile(filepath.Join(dir, name+".html"))
				if err != nil {
					return err
				}
				resultFiles[name] = string(expectResult)
				expectFail[name] = filepath.Base(dir) == "fails"
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	return testFiles, resultFiles, expectFail
}
