package being

import (
	// c "dontstop/cycle"
	f "dontstop/functions"
	maps "dontstop/map"
	"math"
)

// MoveUnit - функция движения на равная на свой размер
func (u *Units) MoveUnit() {

	vectorX := u.X - maps.Coordinates(f.Rand(int(maps.DefaultMap.X)))
	vectorY := u.Y - maps.Coordinates(f.Rand(int(maps.DefaultMap.X)))
	dist := math.Sqrt(float64(vectorX*vectorX + vectorY*vectorY))

	nX := maps.Coordinates(int(math.Abs(float64(vectorX) / float64(dist) * float64(u.Size))))
	nY := maps.Coordinates(int(math.Abs(float64(vectorY) / float64(dist) * float64(u.Size))))

	if nX < maps.DefaultMap.X && nX > 0 {
		(*u).X = nX
	}

	if nY < maps.DefaultMap.X || nY > 0 {
		(*u).Y = nY
	}
}

func (u *Units) LookAround(list []Units) {

	// spaceAtU - space around the unit
	// два размера существа
	var spaceAtU maps.Coordinates = maps.Coordinates(u.Size + u.Size)

	for i := range list {
		if (list[i].X > u.X-spaceAtU || list[i].X < u.X+spaceAtU) && (list[i].Y > u.Y-spaceAtU || list[i].Y < u.Y+spaceAtU) {

			// результат функции MathSign
			x := f.MathSign(int(list[i].X - u.X))
			y := f.MathSign(int(list[i].Y - u.Y))

			LookAroundRS(x, y, u)
		}
	}

	/*
		// каждый из углов вокруг точки
		var leftUpAngle, leftDownAngle, rightUpAngle, rightDownAngle string
		// maps.GridCoordinates
		// первый левый угол
		leftUpAngle = fmt.Sprint(u.X-spaceAtU) + fmt.Sprint(u.Y-spaceAtU)
		leftDownAngle = fmt.Sprint(u.X-spaceAtU) + fmt.Sprint(u.Y+spaceAtU)
		rightUpAngle = fmt.Sprint(u.X+spaceAtU) + fmt.Sprint(u.Y-spaceAtU)
		rightDownAngle = fmt.Sprint(u.X+spaceAtU) + fmt.Sprint(u.Y+spaceAtU)
	*/
}

// LookAroundRS - LookAroundRegionSelection
func LookAroundRS(x int, y int, u *Units) {
	if x < 0 && y < 0 {
		u.Memory.UpLeft = 1
	}

	if x < 0 && y == 0 {
		u.Memory.UpMid = 1
	}

	if x < 0 && y > 0 {
		u.Memory.UpRight = 1
	}

	if x == 0 && y > 0 {
		u.Memory.MidRight = 1
	}

	if x > 0 && y > 0 {
		u.Memory.DownRight = 1
	}

	if x > 0 && y == 0 {
		u.Memory.DownMid = 1
	}

	if x > 0 && y < 0 {
		u.Memory.DownLeft = 1
	}

	if x == 0 && y < 0 {
		u.Memory.MidLeft = 1
	}

	//u.Memory.Step = c.Step
}
