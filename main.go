package main

import "fmt"

func main() {
	neuronLayer := newNeuronLayer(1, 3)
	neuralNet := newNeuralNet(neuronLayer)

	inputs := m64{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1},
	}

	outputs := m64{
		{0},
		{0},
		{0},
		{0},
		{1},
		{1},
		{1},
		{1},
	}

	fmt.Println("Training the neural net...")
	neuralNet.train(inputs, outputs, 1)
	fmt.Println("Finished training")

	fmt.Printf("Layer weights: %v\n", neuronLayer.weights)

	predict(m64{{0, 0, 0}}, neuralNet)
	predict(m64{{0, 0, 1}}, neuralNet)
	predict(m64{{0, 1, 0}}, neuralNet)
	predict(m64{{0, 1, 1}}, neuralNet)
	predict(m64{{1, 0, 0}}, neuralNet)
	predict(m64{{1, 0, 1}}, neuralNet)
	predict(m64{{1, 1, 0}}, neuralNet)
	predict(m64{{1, 1, 1}}, neuralNet)

}

func predict(testInput m64, net *neuralNet) {
	net.think(testInput)

	fmt.Printf("Prediction on data %v %v %v -> %v, expected -> %v\n", testInput[0][0], testInput[0][1], testInput[0][2], net.output()[0][0], testInput[0][0])
}
