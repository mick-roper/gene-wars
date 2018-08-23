package core

import (
	"log"
	"math/rand"
	"time"
)

// World that a population lives in
type World struct {
	r          *rand.Rand
	genes      int
	Population *Population
	Generation int
}

// NewWorld creates a new World
func NewWorld(popSize, genes int) *World {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	p := NewPopulation(popSize, genes)

	return &World{r, genes, p, 0}
}

// Run the simulation
func (w *World) Run() {
	// 1: calculate initial fitness
	w.Population.CalculatePopulationFitness()

	w.printProgress()

	var fittest *Individual // we need this later

	// 2: loop until we converge
	for w.Population.Fittest < w.genes {
		w.Generation++

		// 3: selection
		f1, f2 := w.selectFittest()

		// 4: crossover
		w.crossover(f1, f2)

		// 5: mutate
		if w.r.Int()%7 < 5 {
			w.mutate(f1, f2)
		}

		// 6: add fittest offspring
		f1.CalcFitness()
		f2.CalcFitness()

		// replace weakest offspring with fittest
		if f1.Fitness >= f2.Fitness {
			fittest = f1
		} else {
			fittest = f2
		}

		weakestIx := w.Population.GetLeastFittestIndex()

		w.Population.Individuals[weakestIx] = fittest

		// 7: recalc fitness
		w.Population.CalculatePopulationFitness()

		w.printProgress()
	}

	fittest = w.Population.GetFittest()

	log.Printf("solution found in generation %v\n", w.Generation)
	log.Printf("Fittest: %v\n", fittest.Fitness)
	log.Print("Genes: ")
	for _, g := range fittest.genes {
		log.Print(g)
	}

	log.Print("\n\n")
}

func (w *World) selectFittest() (*Individual, *Individual) {
	return w.Population.GetFittest(), w.Population.GetSecondFittest()
}

func (w *World) crossover(i1, i2 *Individual) {
	crossOverPoint := w.r.Intn(w.genes)

	for i := 0; i < crossOverPoint; i++ {
		x := i1.genes[i]
		i1.genes[i] = i2.genes[i]
		i2.genes[i] = x
	}
}

func (w *World) mutate(i1, i2 *Individual) {
	for _, i := range []*Individual{i1, i2} {
		mutationPoint := w.r.Intn(w.genes)

		// flip values
		if i.genes[mutationPoint] == 1 {
			i.genes[mutationPoint] = 0
		} else {
			i.genes[mutationPoint] = 1
		}
	}
}

func (w *World) printProgress() {
	log.Printf("Generation: %v Fittest: %v", w.Generation, w.Population.Fittest)
}
