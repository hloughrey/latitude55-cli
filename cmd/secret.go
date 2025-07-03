/*
Copyright © 2024 Hugh Loughrey
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// secretCmd represents the secret command
var secretCmd = &cobra.Command{
	Use:   "secret-json-to-string [json-file-path]",
	Short: "Convert JSON file to escaped string format for AWS Secret Manager",
	Long: `Convert a JSON file to an escaped string format suitable for AWS Secret Manager.
	
Example:
  latitude55-cli secret-json-to-string secrets.json
  
Input JSON:
  {
    "REDIS_URL": "abc",
    "REDIS_PASSWORD": "abc"
  }
  
Output:
  {\"REDIS_URL\":\"abc\",\"REDIS_PASSWORD\":\"abc\"}`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		convertJSONToString(args[0])
	},
}

func convertJSONToString(filePath string) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", filePath)
	}

	// Read the JSON file
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Validate that it's valid JSON by unmarshaling and remarshaling
	var jsonObj interface{}
	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		log.Fatalf("Invalid JSON in file: %v", err)
	}

	// Marshal to compact JSON (removes whitespace)
	compactJSON, err := json.Marshal(jsonObj)
	if err != nil {
		log.Fatalf("Error processing JSON: %v", err)
	}

	// Escape quotes for the output
	escapedJSON := strings.ReplaceAll(string(compactJSON), "\"", "\\\"")

	// Output the result
	fmt.Println(escapedJSON)
}

func init() {
	rootCmd.AddCommand(secretCmd)
}
