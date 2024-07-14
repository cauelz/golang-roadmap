package interfaces

import (
	"testing"
)

func TestTriangleArea(t *testing.T) {

	triangle := triangle{base: 10, height: 10}

	triangleArea := triangle.getArea()

	if triangleArea != 50 {
		t.Errorf("Expected 50, but got %.2f", triangleArea)
	}
		
	triangle.printArea()
}

func TestRectangleArea(t *testing.T) {
	
	rectangle := rectangle{width: 10, height: 10}

	rectangleArea := rectangle.getArea()

	if rectangleArea != 100 {
		t.Errorf("Expected 100, but got %.2f", rectangleArea)
	}

	rectangle.printArea()
}