package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Ja7ad/fluora/internal/types"

	_ "embed"
	"gopkg.in/yaml.v3"
)

//go:embed instruction.tmpl
var _template string

func main() {
	instruction := flag.String("instruction", "", "instruction file (yml)")
	outputPath := flag.String("output", "", "output directory for generated file")
	packageName := flag.String("package", "", "package name")
	flag.Parse()

	// Validate input arguments
	if *instruction == "" || *outputPath == "" || *packageName == "" {
		log.Fatal("Usage: -instruction <file.yml> -output <output_dir> -package <package_name>")
	}

	// Read YAML instruction file
	b, err := os.ReadFile(*instruction)
	if err != nil {
		log.Fatalf("Error reading instruction file: %v", err)
	}

	var ins types.Instruction
	if err := yaml.Unmarshal(b, &ins); err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// Normalize instruction title & package name
	fileName := strings.TrimSuffix(filepath.Base(*instruction), filepath.Ext(*instruction))
	title := strings.ReplaceAll(ins.InstructionInfo.Title, " ", "")

	// Generate output file
	if err := generate(&ins, *outputPath, fileName, title, *packageName); err != nil {
		log.Fatalf("Error generating file: %v", err)
	}
}

func generate(ins *types.Instruction, outputPath, fileName, title, packageName string) error {
	// Ensure output directory exists
	if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Parse the template
	tmpl, err := template.New("instruction").Parse(_template)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, ins.ToMap(title, packageName)); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Fix escaping issue
	outputContent := strings.ReplaceAll(buf.String(), "&#34;", "\"")

	// Write output file efficiently
	outputFilePath := filepath.Join(outputPath, fileName+".gen.go")
	if err := os.WriteFile(outputFilePath, []byte(outputContent), 0o644); err != nil {
		return fmt.Errorf("failed to write generated file: %w", err)
	}

	return nil
}
