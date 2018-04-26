// Copyright © 2018 Andreas Fritzler <andreas.fritzler@gmail.com>
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
	"fmt"
	"os"

	"github.com/afritzler/garden-universe/pkg/gardener"
	stats "github.com/afritzler/garden-universe/pkg/stats"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getStats()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

func getStats() {
	kubeconfig := rootCmd.Flag("kubeconfig").Value.String()
	garden, err := gardener.NewGardener(kubeconfig)
	if err != nil {
		fmt.Printf("failed to get garden client for landscape: %s", err)
		os.Exit(1)
	}
	s := stats.NewStats(garden)
	data, err := s.GetStatsJSON()
	if err != nil {
		fmt.Printf("failed to render landscape stats: %s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)
}
