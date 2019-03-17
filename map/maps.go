package maps

// Coordinates тип данных для осей карты
// тип float64
type Coordinates float64

// Maps структура карты
// содержит:
// x      Coordinates
// y      Coordinates
// border int
type Maps struct {
	X      Coordinates
	Y      Coordinates
	Border Coordinates
}

// DefaultMap - настройки карты
var DefaultMap = Maps{
	X:      100,
	Y:      100,
	Border: 150,
}

// Sun используется существами для питания
// x         Coordinates положение по оси x
// y         Coordinates положение по оси y
// intensity int         интенсивность свечения
// radius    int         радиус свечения
type Sun struct {
	X         Coordinates // положение по оси x
	Y         Coordinates // положение по оси y
	Intensity int         // интенсивность свечения
	Radius    Coordinates // радиус свечения
}

// DefaultSun характеристики солнца
var DefaultSun = Sun{
	X:         0,
	Y:         (DefaultMap.Border / 2),
	Intensity: 15,
	Radius:    0,
}

// MoveSun функция движения солнца по квадрату.
func (s *Sun) MoveSun(m Maps) {

	if s.X < m.Border && s.Y == 0 { // увеличиваем Х пока У = 0
		(*s).X = s.X + 1
	}
	if s.X == m.Border && s.Y < m.Border { // Увеличиваем У пока Х == Border
		(*s).Y = s.Y + 1
	}
	if s.X > 0 && s.Y == m.Border { // уменьшаем Х пока У == Border
		(*s).X = s.X - 1
	}
	if s.X == 0 && s.Y > 0 { // уменьшаем У пока Х == Border
		(*s).Y = s.Y - 1
	}

	// уменьшение радиуса свечения солнца
	if s.X == m.Border && s.Y > 0 {
		if (s.X == m.Border && s.Y > 0) && (s.Y < (m.Border / 2)) {
			(*s).Radius = s.Radius - 1
		} else {
			(*s).Radius = 0
		}
	}

	// увиличение радиуса свечения солнца
	if s.X == 0 && s.Y < (m.Border/2) {
		if (s.X == 0 && s.Y < (m.Border/2)) && (s.Y > 0) {
			(*s).Radius = s.Radius + 1
		} else {
			(*s).Radius = (DefaultMap.Border / 2)
		}
	}

	//if (s.X > m.Border && s.Y > (m.Border/2)) || (s.X < m.Border && s.Y > (m.Border/2)) {
	//	if (s.X == m.Border && s.Y > (m.Border/3)) && (s.Y < (m.Border / 2)) {
	//		(*s).Radius = s.Radius - 1
	//	} else {
	//		(*s).Radius = 0
	//	}
	//} else {
	//	if (s.X == m.Border && s.Y > (m.Border/3)) && (s.Y < (m.Border / 2)) {
	//		(*s).Radius = s.Radius + 1
	//	} else {
	//		(*s).Radius = (DefaultMap.Border / 2)
	//	}
	//}

}

/* формула
	if s.X < m.Border && s.Y == 0 {
		(*s).X = s.X + 1
	}
	if s.X == m.Border && s.Y < m.Border {
		(*s).Y = s.Y + 1
	}
	if s.X > 0 && s.Y == m.Border {
		(*s).X = s.X - 1
	}
	if s.X == 0 && s.Y > 0 {
		(*s).Y = s.Y - 1
	}

	if(l<200 && t==0) div.style.left = l+1;

    if(l==200 && t<200) div.style.top = t+1;

    if(l>0 && t==200) div.style.left = l-1;

    if(l==0 && t>0) div.style.top = t-1;

*/
