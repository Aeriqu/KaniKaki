// Package schema handles combining .graphql files into a string.
//
// The .graphql files should located in the path defined by GRAPHQL_SCHEMA_ROOT
// (env var). If that var is not defined, then it will be one of two locations
// depending on how you're running this app:
//
// 1. If uncompiled then in api-gateway/schema
//
// 2. If compiled into a binary then in the same directory as the binary

package schema

import (
	"bytes"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// findGraphQLFiles returns a list of paths that were found to have a
// .graphql extension in the provided schemaRoot. This is recursive.
func findGraphQLFiles(schemaRoot string) []string {
	var files []string
	extension := ".graphql"

	fs.WalkDir(os.DirFS(schemaRoot), ".", func(path string, fileEntry fs.DirEntry, err error) error {
		if err != nil {
			log.Println("[Schema - findGraphQLFiles] Error finding files")
			log.Fatal(err)
		}
		if filepath.Ext(fileEntry.Name()) == extension {
			files = append(files, path)
		}
		return nil
	})

	return files
}

// GetSchemaString returns a string representation of the graphql files located
// by either GRAPHQL_SCHEMA_ROOT (env var) or adjacent to the go file or binary.
func GetSchemaString() string {
	var schema bytes.Buffer
	var schemaRoot string = "."

	if graphQLSchemaRoot, rootExists := os.LookupEnv("GRAPHQL_SCHEMA_ROOT"); rootExists {
		schemaRoot = graphQLSchemaRoot
	}

	graphFilePaths := findGraphQLFiles(schemaRoot)

	for _, filePath := range graphFilePaths {
		file, fileError := os.ReadFile(filePath)
		if fileError != nil {
			log.Printf("[Schema - GetSchemaString] Error Reading File (%q)", filePath)
			log.Fatal(fileError)
		}

		schema.Write(file)
		schema.WriteByte('\n')
	}

	return schema.String()
}
