package main

import (
	"math"
	"math/rand"
)

type neuronLayer struct {
	weights m64
}

func newNeuronLayer(numNeurons, numInputsPerNeuron int) *neuronLayer {

	var n neuronLayer
	n.weights = newM64(numInputsPerNeuron, numNeurons)

	for i := 0; i < numInputsPerNeuron; i++ {
		for j := 0; j < numNeurons; j++ {
			n.weights[i][j] = 2*rand.Float64() - 1
		}
	}

	return &n
}

func (n *neuronLayer) Activate(sum float64) (output float64) {
	return 1.0 / (1.0 + math.Exp(-1.0*sum))
}

func (n *neuronLayer) Gradient(output float64) (gradient float64) {
	return output * (1 - output)
}

func (n *neuronLayer) AdjustWeights(adjustment m64) {
	n.weights = n.weights.Add(adjustment)
}
