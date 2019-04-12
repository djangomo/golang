// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a element",
	Long:  `removes an element from the List based on Linenumber`,
	Run: func(cmd *cobra.Command, args []string) {
		line, _ := cmd.Flags().GetString("lines")
		if line == "" {
			line = "missing --line"
		}

		// using slice read from File
		a := readFile()
		d, _ := strconv.Atoi(line)
		d--

		// deleting entered line and keeping order
		copy(a[d:], a[d+1:])
		a[len(a)-1] = ""
		a = a[:len(a)-1]

		deletFile()

		saveFile(a)
	},
}

func readFile() []string {
	// reads file and returns a slice with the content
	notesfile, err := os.Open("/home/djangomo/input.txt")
	if err != nil {
		log.Fatalf("no such file %s", err)
	}

	scanner := bufio.NewScanner(notesfile)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	notesfile.Close()

	return txtlines
}

func saveFile(arrayToSave []string) {
	// saves File to predefined location TODO: change to optional directory ?
	fileHandle, err := os.OpenFile("/home/djangomo/input.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fileHandle.Close()
	sep := "\n"
	for _, line := range arrayToSave {
		if _, err = fileHandle.WriteString(line + sep); err != nil {
			panic(err)
		}
	}
}

func deletFile() {
	// Truncates File to 0 bytes - deletes the content for rewrite
	err := os.Truncate("/home/djangomo/input.txt", 0)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringP("lines", "l", "", "Enter linenumber")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
