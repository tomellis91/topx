/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var n string
var path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "topx",
	Short: "Scan a file and locate the top N values",
	Long:  `Scan a large text file of integers and return the top N values.`,
	Run: func(cmd *cobra.Command, args []string) {
		ni, _ := strconv.Atoi(n)
		scan(path, ni)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&n, "number", "n", "", "(required) The top N number of results to return")
	rootCmd.Flags().StringVarP(&path, "path", "p", "./data/numbers.txt", "The path to the file to search")
	rootCmd.MarkFlagRequired("number")
}
