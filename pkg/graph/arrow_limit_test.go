package graph

import "testing"

func TestGetPathFailsForEnclosedTarget(t *testing.T) {
	wall := &node{}
	g := &graph{grid: map[gridCoord]*node{
		{x: 4, y: 5}: wall,
		{x: 6, y: 5}: wall,
		{x: 5, y: 4}: wall,
		{x: 5, y: 6}: wall,
	}}

	if path, err := g.getPath(gridCoord{x: 0, y: 0}, gridCoord{x: 5, y: 5}); err == nil {
		t.Fatalf("getPath() = %v, want an error for a target no path reaches", path)
	}
}
