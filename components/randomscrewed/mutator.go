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

var mutations = []Mutator{}

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

func screwAroundFindOut(m string, into chan *string, multiple bool) {
	n := 2
	for range 1000 / THREADS {
		current := m
		for range n + 1 {
			current = screwedIteration(current)
		}
		if len(current) < 2000 && checkScrewed(current) {
			into <- &current
			if !multiple {
				return
			}
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
		go screwAroundFindOut(m, into, false)
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

func screwedMutateMultiple(m string) []string {
	slc := []string{}
	if checkScrewed(m) {
		slc = append(slc, m)
	}
	into := make(chan *string)
	for range THREADS {
		go screwAroundFindOut(m, into, true)
	}

	finishedCount := 0

	for finishedCount < THREADS {
		x := <-into
		if x == nil {
			finishedCount++
		} else {
			slc = append(slc, *x)
		}
	}

	return slc
}

func screwedify(bot *botsession.BotSession, inter *discordgo.Interaction) {
	inputValue := inter.ApplicationCommandData().Options[0].StringValue()
	if strings.Contains(inputValue, "@") {
		bot.RespondWithMessage(inter, common.RandomChoice(common.Conf.Gifs))
		return
	}

	multi := false
	if multiOpt := inter.ApplicationCommandData().GetOption("multiple"); multiOpt != nil {
		multi = multiOpt.BoolValue()
	}
	if multi {
		slc := screwedMutateMultiple(inputValue)
		betterSlice := make([]string, len(slc))
		for _, i := range slc {
			betterSlice = append(betterSlice, "```\n"+i+"```")
		}
		bot.RespondWithMessage(inter, strings.Join(betterSlice, "\n"))
	} else if screwedified, ok := screwedMutate(inputValue); ok {
		bot.RespondWithMessage(inter, screwedified)
	} else {
		bot.RespondWithMessage(inter, "I could not find a message that is SCREWED enough...")
	}
}

func MutatorInit() {
	mutations = make([]Mutator, len(common.Conf.Mutations))
	for _, i := range common.Conf.Mutations {
		mutations = append(mutations, Mutator{
			weight: i.Probability,
			handler: func(s string) string {
				switch i.Filter {
				case "randomUpper":
					s = randomUpper(s)
				case "titleCaser":
					s = titleCaser(s)
				}
				return strings.ReplaceAll(i.Output, "%", s)
			},
		})
	}
}
