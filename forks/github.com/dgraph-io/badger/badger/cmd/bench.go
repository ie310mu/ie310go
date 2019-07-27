/*
 * Copyright 2019 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/ie310mu/ie310go/forks/github.com/spf13/cobra"
)

var benchCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "Benchmark Badger database.",
	Long: `This command will benchmark Badger for different usecases. Currently only read benchmark
	is supported. Useful for testing and performance analysis.`,
}

func init() {
	RootCmd.AddCommand(benchCmd)
}
