/*
Copyright © 2024 Hugh Loughrey <hugh.loughrey@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"

	"github.com/hashicorp/go-getter"
	"github.com/spf13/cobra"
)

var (
	templateType string
	applicationName string
)

type TemplateType string

const (
    Node          TemplateType = "node"
    AwsServerless TemplateType = "aws-serverless"
    ExpressAPI    TemplateType = "express-api"
    NestAPI       TemplateType = "nest-api"
)

var permittedTemplateTypes = []TemplateType{Node, AwsServerless, ExpressAPI, NestAPI}


var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Utils for bootstrapping new apps",

	Run: func(cmd *cobra.Command, args []string) {
		createAppFromTemplate()
	},
}

func createDirectory() string {
    dir, err := os.Getwd() 
    if err != nil {
				log.Fatal(err)
    }

	var appDirectory = filepath.Join(dir, "..", applicationName)

	err = os.MkdirAll(appDirectory, 0750)
	if err != nil {
		log.Fatal(err)
	}

	return appDirectory
}

func downloadAndExtract(appDirectory string) {

	client := &getter.Client{
		Mode: getter.ClientModeDir,
		Dst: appDirectory,
		Dir: true,
		Src: "github.com/hloughrey/latitude55-templates/templates/" + templateType,
	}

	err := client.Get(); 
	
	if err != nil {
		log.Fatal(err)
	}

}

func createAppFromTemplate() {
	if !slices.Contains(permittedTemplateTypes, TemplateType(templateType)) {
			log.Fatal(`Unsupported template type`)
	}
	fmt.Printf("Creating application %s using template %s \n", applicationName, templateType)
	
	var appDirectory = createDirectory()
	downloadAndExtract(appDirectory)
}

func init() {
	templateCmd.Flags().StringVarP(&templateType, "template", "t", "", `Template used to bootstrap your application, allowed options: "node", "aws-serverless", "express-api", "nest-api"`)
	templateCmd.MarkFlagRequired("template")

	templateCmd.Flags().StringVarP(&applicationName, "name", "n", "", "Name of your new application")
	templateCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(templateCmd)
}
