package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jonyhy96/replacer/pkg/replacer"
	"github.com/jonyhy96/replacer/pkg/utils/file"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.Flags().StringP("file", "f", "", "the replace file filepath e.g. ./keys.json")
	rootCmd.Flags().StringP("work", "w", "", "the work path of files you want to replace")
	rootCmd.Flags().StringVar(&fileType, "t", "", "the type of your replace file(default is json)")
	rootCmd.MarkFlagRequired("file")
	rootCmd.MarkFlagRequired("path")
	viper.SetDefault("author", "jonyhy <github/jonyhy96>")
}

var (
	fileType        = ""
	defaultFileType = "json"
	rootCmd         = &cobra.Command{
		Use:     "replacer",
		Short:   "replacer eplace things for you\n",
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			filePath, _ := cmd.Flags().GetString("file")
			workPath, _ := cmd.Flags().GetString("work")
			filepath.Walk(workPath, func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				if fileType == "" {
					replaceMap, err := file.New(defaultFileType).Transform(filePath)
					if err != nil {
						log.Fatalln(err)
					}
					file, _ := os.OpenFile(filePath, os.O_RDWR, 0644) // omite error because of already under walk.
					err = replacer.New(defaultFileType, replaceMap).Replace(file, file)
					if err != nil {
						log.Fatalln(err)
					}
				}
				return nil
			})
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
