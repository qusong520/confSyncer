/*
 * 	ConfSyncer - a little sync config files tool in the Linux.
 *     Copyright (C) 2020  Amatist_kurisu<misaki.zhcy@gmail.com>
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package main

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deamonCmd represents the deamon command
var deamonCmd = &cobra.Command{
	Use:   "deamon",
	Short: "deamon",
	Long:  `deamon`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := time.NewTicker(time.Duration(viper.GetInt("gitPullTimeInternal")) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := ConfigPull()
				if err != nil {
					log.Fatalln(err.Error())
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deamonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deamonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deamonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
