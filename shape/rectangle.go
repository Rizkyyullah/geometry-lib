package shape

type Rectangle struct {
  Width, Length float32
}

func (rectangle *Rectangle) Area() float32 {
  return rectangle.Width * rectangle.Length
}

func (rectangle *Rectangle) Perimeter() float32 {
  return 2 * (rectangle.Width + rectangle.Length)
}