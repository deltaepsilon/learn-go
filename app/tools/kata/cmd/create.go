/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

const packagePath = "github.com/deltaepsilon/learn-go"
const kataRoot = "/app/kata-test"
const unknownPackagePath = "UNKNOWN_PACKAGE_PATH"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new kata",
	Long: `Create a kata with a readable name and a package name:
usage: kata create <kata-name> <package-name>
example: kata create playing-with-digits digits
`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		folderName := args[0]
		packageName := args[1]

		createPackageFile(folderName, packageName)
		initializeGinkgo(folderName, packageName, &wg)

		wg.Wait()

		setTestImport(folderName, packageName)

		fmt.Printf("Created name: %v package: %v\n", folderName, packageName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createPackageFile(folderName, packageName string) {
	folderPath := filepath.Join(kataRoot, folderName)
	filePath := filepath.Join(folderPath, packageName+".go")

	os.RemoveAll(folderPath)
	os.MkdirAll(folderPath, os.ModePerm)
	file, _ := os.Create(filePath)

	defer file.Close()

	file.Write([]byte("package " + packageName))
}

func initializeGinkgo(folderName, packageName string, wg *sync.WaitGroup) {
	binary, lookPathError := exec.LookPath("ls")

	check(lookPathError)

	fmt.Println("binary", binary)

	wg.Add(1)
	go execInFolder(folderName, "ginkgo bootstrap", wg)

	wg.Add(1)
	go execInFolder(folderName, "ginkgo generate "+packageName, wg)
}

func execInFolder(folderName string, command string, wg *sync.WaitGroup) {
	args := strings.Fields(command)
	folderPath := filepath.Join(kataRoot, folderName)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = folderPath
	out, _ := cmd.Output()

	fmt.Println(string(out))

	wg.Done()
}

func setTestImport(folderName, packageName string) {
	folderPath := filepath.Join(kataRoot, folderName)
	filePath := filepath.Join(folderPath, packageName+"_test.go")
	read, _ := ioutil.ReadFile(filePath)
	replaced := strings.Replace(string(read), unknownPackagePath, packagePath+"/"+folderName, 1)
	err := ioutil.WriteFile(filePath, []byte(replaced), os.ModePerm)

	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
