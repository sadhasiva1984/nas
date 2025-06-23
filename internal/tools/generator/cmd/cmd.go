package main

import "github.com/sadhasiva1984/nas/internal/tools/generator"

func main() {
	generator.ParseSpecs()

	generator.GenerateNasMessage()

	generator.GenerateNasEncDec()

	generator.GenerateTestLarge()
}
