package vector2

import "iter"

func (v Vector[T]) EnumRegion(w, h T, step ...T) iter.Seq[Vector[T]] {
	return func(yield func(Vector[T]) bool) {
		sx := T(1)
		sy := T(1)
		if len(step) > 0 {
			sx = step[0]
		}
		if len(step) > 1 {
			sy = step[1]
		}

		for y := v.Y; y <= v.Y+h; y += sy {
			for x := v.X; x <= v.X+w; x += sx {
				if !yield(Vector[T]{X: x, Y: y}) {
					return
				}
			}
		}
	}
}

func (v Vector[T]) EnumRegionAround(w, h T, step ...T) iter.Seq[Vector[T]] {
	return func(yield func(Vector[T]) bool) {
		sx := T(1)
		sy := T(1)
		if len(step) > 0 {
			sx = step[0]
		}
		if len(step) > 1 {
			sy = step[1]
		}

		for y := v.Y - h; y <= v.Y+h; y += sy {
			for x := v.X - w; x <= v.X+w; x += sx {
				if !yield(Vector[T]{X: x, Y: y}) {
					return
				}
			}
		}
	}
}
