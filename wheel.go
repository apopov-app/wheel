package wheel

import (
	"math/rand"
	"time"
	"errors"
)

type callbackType map[float64][]interface{}

type wheel struct {
	callbacks callbackType
	chanceMin float64
	chanceMax float64
	allChance float64
}

func (w *wheel) AddValue(chance float64, value interface{}) error {
	if _, exist := w.callbacks[chance]; !exist {
		w.callbacks[chance] = make([]interface{}, 0)
	}

	if chance < w.chanceMin || w.chanceMin == 0 {
		w.chanceMin = chance
	}

	if chance > w.chanceMax {
		w.chanceMax = chance
	}

	w.allChance += chance
	if w.allChance > 1.0 {
		return errors.New("to many chance value")
	}

	w.callbacks[chance] = append(w.callbacks[chance], value)
	return nil
}

func (w *wheel) Spin() interface{} {
	rand.Seed(time.Now().UnixNano())
	num := w.chanceMin + rand.Float64() * (w.chanceMax - w.chanceMin)
	prevNum := 0.0
	for cnum, valueSlice := range w.callbacks {
		if num >= prevNum && num <= cnum {
			callPos := rand.Intn(len(valueSlice))

			for n, value := range valueSlice {
				if n == callPos {
					return value
				}
			}
		}

		prevNum = cnum
	}
	return nil
}

func New() *wheel {
	return &wheel{
		callbacks: make(callbackType),
	}
}