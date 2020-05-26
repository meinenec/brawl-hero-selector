package cmd

import (
	"fmt"
	"meinenec/brawl-hero-selector/heroes"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type heroRequest struct {
	num   int
	hPool string
}

// HandleBrawl is a handler for resolving the brawl bot command
func HandleBrawl(message []string) *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}

	// Check for errors on message input from handleBrawlInput
	input, err := handleBrawlInput(message)
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title:       "Brawl Hero Selector",
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xad0000,
			Description: err.Error(),
			//Fields:      fields,
		}
		return embed
	}

	// Define and populate fields with hero names and roles selected from heroes slice
	for _, h := range heroes.Assign(input.num, input.hPool) {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   h.Role,
			Value:  h.Name,
			Inline: true,
		})
	}

	// Populate embed with fields
	embed := &discordgo.MessageEmbed{
		Title:       "Brawl Hero Selector",
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x36A8DE,
		Description: "Pick your Hero!",
		Fields:      fields,
	}
	return embed
}

func handleBrawlInput(message []string) (heroRequest, error) {

	input := heroRequest{
		num:   3,
		hPool: "brawl",
	}

	if len(message) > 1 {
		err := input.assignInputValues(message[1])
		if err != nil {
			return input, err
		}
	}

	if len(message) > 2 {
		err := input.assignInputValues(message[2])
		if err != nil {
			return input, err
		}
	}

	return input, nil
}

func (input *heroRequest) assignInputValues(message string) error {

	mNum, err := strconv.Atoi(message)
	// If error exists treat as hPool
	if err != nil {
		err := validateBrawlHPool(message)
		if err != nil {
			return err
		}
		input.hPool = message
		// If error is nil treat as num
	} else {
		err := validateBrawlNum(mNum)
		if err != nil {
			return err
		}
		input.num = mNum
	}

	return nil
}

func validateBrawlHPool(hPool string) error {

	okPool := []string{"all", "brawl"}

	if contains(okPool, hPool) {
		return nil
	}
	return fmt.Errorf("Hero pool requested is invalid. got: %s, wanted: %v", hPool, okPool)
}

func validateBrawlNum(mNum int) error {

	max := 25
	min := 1

	if mNum > max {
		return fmt.Errorf("Number of heroes requested is larger than the max. got: %d, max: %d", mNum, max)
	}
	if mNum < min {
		return fmt.Errorf("Number of heroes requested is smaller than the minimum. got: %d, min: %d", mNum, min)
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.EqualFold(a, e) {
			return true
		}
	}
	return false
}
