/*
Copyright © 2019 Alex Engelberg

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

	"github.com/aengelberg/gh-actions-orb/action"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-actions-orb",
	Short: "Run a GitHub Action in an orb",
	Long:  `Run a GitHub Action with the specified inputs, while simulating the GitHub execution platform which translates certain stdout sequences into outputs and UI appearances.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		actionName := args[0]
		inputsFileName := args[1]
		action.Run(actionName, parseStepInputs(inputsFileName))
	},
}

// Takes the name of a YAML file, then parses it into a map of step inputs
// to pass to the GH Action.
func parseStepInputs(inputsFileName string) action.StepInputs {
	in, err := os.Open(inputsFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer in.Close()
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var yamlData map[string]interface{}
	if err := yaml.Unmarshal(bytes, &yamlData); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	stepInputs := make(map[string]string)
	for k, v := range yamlData {
		if s, ok := v.(string); ok {
			stepInputs[k] = s
		} else {
			fmt.Printf("Found %t in `with` value; expected string", v)
			os.Exit(1)
		}
	}
	return stepInputs
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-actions-orb.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gh-actions-orb" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gh-actions-orb")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
