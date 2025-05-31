package randomscrewed

import (
	"main/botsession"
	"main/common"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Mutator struct {
	weight  float64
	handler func(s string) string
}

var mutations = []Mutator{
	{
		weight: 0.5,
		handler: func(s string) string {
			return s + " screwed"
		},
	},
	{
		weight: 1.5,
		handler: func(s string) string {
			return s + " in haxe"
		},
	},
	{
		weight: 1.5,
		handler: func(s string) string {
			return "filled " + s
		},
	},
	{
		weight: 0.5,
		handler: func(s string) string {
			return s + " in one go im praying it works"
		},
	},
	{
		weight: 0.20,
		handler: func(s string) string {
			return "am am says that " + s
		},
	},
	{
		weight: 0.05,
		handler: func(s string) string {
			return "and i said \"**" + s + "**\" but nothing happened"
		},
	},
	{
		weight: 0.025,
		handler: func(s string) string {
			return "and i said \"**" + s + "**\" but everything happened"
		},
	},
	{
		weight:  0.05,
		handler: randomUpper,
	},
	{
		weight:  0.25,
		handler: titleCaser,
	},
}

func randomUpper(x string) string {
	final := ""
	for index := range x {
		switch rand.Intn(2) {
		case 0:
			final += strings.ToUpper(x[index : index+1])
		case 1:
			final += strings.ToLower(x[index : index+1])
		}
	}
	return final
}

const alphabeticalLower = "qwertyuiopasdfghjklzxcvbnm"
const alphabeticalUpper = "QWERTYUIOPASDFGHJKLZXCVBNM"
const alphabetical = alphabeticalLower + alphabeticalUpper

func titleCaser(x string) string {
	final := ""
	nextUpper := true
	for index := range x {
		ch := x[index : index+1]
		if !strings.Contains(alphabetical, ch) {
			final += ch
			nextUpper = true
		} else if nextUpper {
			final += strings.ToUpper(ch)
			nextUpper = false
		} else {
			final += ch
		}
	}
	return final
}

func screwedIteration(x string) string {
	// sum all weights
	totalWeight := float64(0)
	for _, mut := range mutations {
		totalWeight += mut.weight
	}
	point := rand.Float64() * float64(totalWeight)
	for _, mut := range mutations {
		if point < mut.weight {
			return mut.handler(x)
		}
		point -= mut.weight
	}
	return x
}

const THREADS = 12

func screwAroundFindOut(m string, into chan *string) {
	n := 2
	for range 1000 / THREADS {
		current := m
		for range n + 1 {
			current = screwedIteration(current)
		}
		if len(current) < 2000 && checkScrewed(current) {
			into <- &current
			return
		}
		if rand.Intn(250/THREADS/n) == 1 {
			n++
		}
	}
	into <- nil
}

func screwedMutate(m string) (string, bool) {
	if checkScrewed(m) {
		return m, true
	}
	into := make(chan *string)
	for range THREADS {
		go screwAroundFindOut(m, into)
	}

	failures := 0

	for failures < THREADS {
		x := <-into
		if x == nil {
			failures++
		} else {
			return *x, true
		}
	}

	return "", false
}

func screwedify(bot *botsession.BotSession, inter *discordgo.Interaction) {
	inputValue := inter.ApplicationCommandData().Options[0].StringValue()
	if strings.Contains(inputValue, "@") {
		bot.RespondWithMessage(inter, common.RandomChoice(common.Gifs))
		return
	}
	if screwedified, ok := screwedMutate(inputValue); ok {
		bot.RespondWithMessage(inter, screwedified)
	} else {
		bot.RespondWithMessage(inter, "I could not find a message that is SCREWED enough...")
	}
}
