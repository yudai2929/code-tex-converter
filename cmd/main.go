package main

import (
	"code-tex-converter/pkg/filewalker"
	"code-tex-converter/pkg/latex"
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
			&cli.BoolFlag{
				Name:    "lstinput",
				Value:   false,
				Usage:   "use lstinputlisting",
				Aliases: []string{"i"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			dir := cmd.String("target")
			lstinput := cmd.Bool("lstinput")

			paths, err := filewalker.Walk(dir)
			if err != nil {
				return err
			}

			converted, err := latex.Convert(paths, lstinput)
			if err != nil {
				return err
			}

			for _, c := range converted {
				fmt.Println(c)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
