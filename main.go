package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "cryptopals"
	app.Usage = "The CryptoPals Challenges"
	app.Commands = []cli.Command{
		{
			Name:  "one",
			Usage: "Hex to Base64",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input",
					Value: "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
				},
			},
			Action: func(c *cli.Context) error {
				out := ChallengeOne(c.String("input"))
				println(out)
				return nil
			},
		},
		{
			Name:  "two",
			Usage: "Fixed length XOR",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input",
					Value: "1c0111001f010100061a024b53535009181c",
				},
				cli.StringFlag{
					Name:  "key",
					Value: "686974207468652062756c6c277320657965",
				},
			},
			Action: func(c *cli.Context) error {
				println(ChallengeTwo(c.String("input"), c.String("key")))
				return nil
			},
		},
		{
			Name:  "three",
			Usage: "Hex to Base64",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input",
					Value: "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
				},
			},
			Action: func(c *cli.Context) error {
				out := ChallengeThree(c.String("input"))
				fmt.Println(out)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
