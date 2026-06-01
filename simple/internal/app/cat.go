package app

import (
	"fmt"
	"os"
)

type CatCommand struct {
	Path string `arg:"" name:"path" help:"Path to read contents from"`
}

func (c *CatCommand) Run(ctx CommandContext) error {
	content, err := os.ReadFile(c.Path)
	if err != nil {
		return fmt.Errorf("CatCommand#Run failed to read file %s: %w", c.Path, err)
	}
	fmt.Println(string(content))
	return nil
}
