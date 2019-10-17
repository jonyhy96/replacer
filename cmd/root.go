package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jonyhy96/replacer/pkg/replacer"
	"github.com/jonyhy96/replacer/pkg/utils/file"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.Flags().StringP("file", "f", "", "the replace file filepath e.g. ./keys.json")
	rootCmd.Flags().StringP("work", "w", "", "the work path of files you want to replace")
	rootCmd.Flags().StringSliceP("exclude", "e", []string{}, "specify exclude file names e.g. go.mod,go.sum")
	rootCmd.Flags().StringVar(&fileType, "t", "", "the type of your replace file(default is json)")
	rootCmd.MarkFlagRequired("file")
	rootCmd.MarkFlagRequired("work")
	viper.SetDefault("author", "jonyhy <github/jonyhy96>")
}

var (
	fileType        = ""
	defaultFileType = "json"
	_replaceMap     map[string]interface{}
	err             error
	wg              = sync.WaitGroup{}
	rootCmd         = &cobra.Command{
		Use:     "replacer",
		Short:   "replacer replace things for you\n",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			replaceFilePath, _ := cmd.Flags().GetString("file")
			workPath, _ := cmd.Flags().GetString("work")
			excludedFiles, _ := cmd.Flags().GetStringSlice("exclude")
			if fileType == "" {
				_replaceMap, err = file.New(defaultFileType).Transform(replaceFilePath)
				if err != nil {
					log.Fatalln(err)
				}
			}
			filepath.Walk(workPath, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if strings.Contains(path, ".git") {
					return nil
				}
				if strings.Contains(replaceFilePath, info.Name()) {
					return nil
				}
				if len(excludedFiles) > 0 {
					for _, filename := range excludedFiles {
						if info.Name() == filename {
							fmt.Println(filename)
							return nil
						}
					}
				}
				wg.Add(1)
				go func() {
					defer wg.Done()
					in, _ := os.Open(path) // omite error because of already under walk.
					defer in.Close()
					err = replacer.New(defaultFileType, _replaceMap, path).Replace(in)
					if err != nil {
						log.Fatalln(err)
					}
				}()
				return nil
			})
			wg.Wait()
			log.Printf("Success replace all files under %s by %s!\n", workPath, replaceFilePath)
		},
	}
)

// Execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
