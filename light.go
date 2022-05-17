package main

import "fmt"

type NightLight int

const (
	START     NightLight = iota
	STATE_ON             = iota
	STATE_OFF            = iota
	END                  = iota
)

type Params struct {
	state NightLight
	time  int
	power bool
	light bool
}

func NightlightWorking(params *Params, timet int, power bool) {
	params.time = timet
	params.power = power
	switch params.state {
	case START:
		{
			fmt.Printf("STATE START ")
			if params.time > 3 && params.time <= 190 && params.power {
				params.state = STATE_ON
			}
			if params.time >= 190 && params.time < 270 && params.power {
				params.state = STATE_OFF
			}
			break
		}
	case STATE_ON:
		{
			fmt.Printf("STATE OON ")

			params.light = true
			if params.time > 3 && params.time <= 190 && params.power {
				params.state = STATE_ON
			} else {
				params.state = END
			}
			if params.time >= 190 && params.time < 270 && params.power {
				params.state = STATE_OFF
			}
			break
		}
	case STATE_OFF:
		{
			fmt.Printf("STATE OFF ")
			params.light = false
			if params.time > 3 && params.time <= 190 && params.power {
				params.state = STATE_ON
			} else {
				params.state = END
			}
			if params.time >= 190 && params.time < 270 && params.power {
				params.state = STATE_OFF
			}
			break
		}
	case END:
		{
			fmt.Printf("STATE END ")
			params.state = END
			params.time = 0
			params.power = false
			params.light = false
			break
		}
	}
}

func main() {
	lightPar := Params{START, 0, true, false}
	fmt.Println(lightPar)
	power := true

	for i := 0; i < 269; i++ {
		NightlightWorking(&lightPar, i, power)
		fmt.Println(lightPar)
		if i == 99 {
			for j := 0; j < 22; j++ {

				NightlightWorking(&lightPar, i, power)
				fmt.Println(lightPar)
			}
		}
		if i == 250 {
			power = false

		}
	}
}
