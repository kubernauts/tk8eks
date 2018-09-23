// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"os/exec"

	"github.com/CrowdSurge/banner"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing EKS cluster",
	Long: `
	Delete an EKS cluster on AWS, this command does not require any arguments or flags`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			fmt.Println()
			fmt.Println("Invalid, there is no need to use arguments with this command")
			fmt.Println()
			fmt.Println("Simple use : ekscluster delete")
			fmt.Println()
			os.Exit(0)
		}

		banner.Print("kubernauts eks cli")

		fmt.Println()

		fmt.Println()

		// Check if terraform binary is present in the working directory
		if _, err := os.Stat("./terraform"); err != nil {
			log.Fatalln("Terraform binary not found in the installation folder")
		}

		log.Println("Terraform binary exists in the installation folder, terraform version:")

		terr, err := exec.Command("./terraform", "version").Output()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(terr))

		log.Println("Checking AWS Credentials")

		if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
			log.Fatalln("AWS_ACCESS_KEY_ID not exported as environment variable, kindly check")
		}

		if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
			log.Fatalln("AWS_SECRET_ACCESS_KEY not exported as environment variable, kindly check")
		}

		// Check if a terraform state file already exists
		if _, err := os.Stat("./terraform.tfstate"); err != nil {
			log.Fatalln("Terraform.tfstate file not found, seems there is no existing cluster definition in this directory")
		}

		// Terraform destroy the EKS cluster

		log.Println("starting terraform destroy")

		terrDel := exec.Command("terraform", "destroy", "-force")
		terrDel.Dir = "./"
		out, _ := terrDel.StdoutPipe()
		terrDel.Start()
		scanDel := bufio.NewScanner(out)
		for scanDel.Scan() {
			m := scanDel.Text()
			fmt.Println(m)
		}

		terrDel.Wait()

		// Delete terraform state file

		log.Println("Removing the terraform state file")

		err = os.Remove("./terraform.tfstate")
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
