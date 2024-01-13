package main

import (
	"code-tex-converter/pkg/filewalker"
	"code-tex-converter/pkg/latex"
	"code-tex-converter/pkg/utils"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "code-tex-converter ",
		Usage: "convert code to latex listing format",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "target",
				Usage:    "laTeX listing format target directory",
				Aliases:  []string{"t"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "base",
				Value:   "",
				Usage:   "base path",
				Aliases: []string{"b"},
			},
			&cli.BoolFlag{
				Name:    "lstinput",
				Value:   false,
				Usage:   "use lstinputlisting",
				Aliases: []string{"i"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			dir := cmd.String("target")
			islLtInput := cmd.Bool("lstinput")
			basePath := cmd.String("base")

			resolvedBasePath := utils.TrimPath(basePath)

			paths, err := filewalker.Walk(dir)
			if err != nil {
				return err
			}

			texListings, err := latex.Convert(paths, islLtInput, dir, resolvedBasePath)

			if err != nil {
				return err
			}

			displayTexListings(texListings)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func displayTexListings(texListings []string) {
	for _, texListing := range texListings {
		fmt.Println(texListing)
	}
}
