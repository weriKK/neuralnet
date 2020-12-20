package main

type neuralNet struct {
	layer       *neuronLayer
	outputLayer m64
}

func newNeuralNet(layer *neuronLayer) *neuralNet {
	n := neuralNet{
		layer:       layer,
		outputLayer: make(m64, 0),
	}

	return &n
}

func (n *neuralNet) think(inputs m64) {
	n.outputLayer = inputs.MMul(n.layer.weights).ApplyFunc(n.layer.Activate)
}

func (n *neuralNet) train(inputs, outputs m64, numberOfTrainingIterations int) {
	for i := 0; i < numberOfTrainingIterations; i++ {

		n.think(inputs)

		errorLayer := outputs.Sub(n.outputLayer)
		deltaLayer := errorLayer.SMul(n.outputLayer.ApplyFunc(n.layer.Gradient))

		adjustmentLayer := inputs.Transpose().MMul(deltaLayer)

		n.layer.AdjustWeights(adjustmentLayer)
	}
}

func (n *neuralNet) output() m64 {
	return n.outputLayer
}
