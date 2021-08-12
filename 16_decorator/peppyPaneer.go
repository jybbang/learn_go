package main

type peppyPaneer struct {
}

// Implements pizza
func (p *peppyPaneer) getPrice() int {
	return 20
}
