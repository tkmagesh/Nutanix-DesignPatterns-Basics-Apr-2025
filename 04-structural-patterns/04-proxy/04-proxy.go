/* Go Example – Virtual Proxy for Image Loading */
package main

import "fmt"

// Subject interface
type Image interface {
	Display()
}

// RealImage - the actual object
type RealImage struct {
	filename string
}

func (r *RealImage) Display() {
	fmt.Println("Displaying", r.filename)
}

func NewRealImage(filename string) *RealImage {
	fmt.Println("Loading", filename)
	return &RealImage{filename}
}

// ProxyImage - the proxy
type ProxyImage struct {
	filename string
	real     *RealImage
}

func (p *ProxyImage) Display() {
	if p.real == nil {
		p.real = NewRealImage(p.filename)
	}
	p.real.Display()
}

func main() {
	image := &ProxyImage{filename: "photo.png"}
	image.Display() // loads + displays
	image.Display() // only displays
}
