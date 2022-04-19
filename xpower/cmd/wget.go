/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	output string
)

// wgetCmd represents the wget command
var wgetCmd = &cobra.Command{
	Use:     "wget",
	Example: "xpower wget iqsing.github.io/download.tar.gz -o /tmp/download.tar.gz",
	Args:    cobra.ExactArgs(1),
	Short:   "wget is a download cli.",
	Long:    `use wget to download everything you want from net.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---wget running---")
		Download(args[0], output)
	},
}

func init() {
	rootCmd.AddCommand(wgetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wgetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wgetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	wgetCmd.Flags().StringVarP(&output, "output", "o", "", "output file")
	wgetCmd.MarkFlagRequired("output")
}
func Download(url string, path string) {
	out, err := os.Create(path)
	check(err)
	defer out.Close()

	res, err := http.Get(url)
	check(err)
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	check(err)
	fmt.Println("save as" + path)
}
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
