package being

import (
	f "dontstop/functions"
	maps "dontstop/map"
)

// hp - тип данных для здоровья
// float64
// используется в структуре существа
type hp float64

// color это RGB цвет существа
// one   int первый
// two   int второй
// three int третий

// DefaultHp - стартовый размер здоровья при генерации существа
var DefaultHp hp = 500

// DefaultLiveStatus - стартовый статус существа
var DefaultLiveStatus int = 1

// DefaultColor - стартовый цвет
var DefaultColor color = color{185, 244, 66}

// DefaultSize - стартовый размер существа // пикселей
var DefaultSize int = 15

// DefaultEnergy - стартовый размер энергии
var DefaultEnergy int = 100

// DefaultAge - стартовый возраст :)
var DefaultAge int = 0

// color - структура RGB для цвета существа
type color struct {
	one   int // первый
	two   int // второй
	three int // третий
}

// Colors структура возможных цветов
//  a1 = 0,255,0
//	a2 = 85,212,0
//	a3 = 128,170,0
//	a4 = 170,128,0
//	a5 = 212,85,0
//	a6 = 255,0,0
type colors struct {
	a1 color
	a2 color
	a3 color
	a4 color
	a5 color
	a6 color
}

// ListColor - набор RGB цветов
// Структура - структур
var ListColor = colors{
	a1: color{
		one:   0,
		two:   255,
		three: 0,
	},
	a2: color{
		one:   85,
		two:   212,
		three: 0,
	},
	a3: color{128, 170, 0},
	a4: color{170, 128, 0},
	a5: color{212, 85, 0},
	a6: color{255, 0, 0},
}

// Memory - память, записывает занято/свободно место
// вокруг юнита и шаг на котором проводился анализ пространства
type Memory struct {
	UpLeft    int
	UpMid     int
	UpRight   int
	MidLeft   int
	MidMid    int
	MidRight  int
	DownLeft  int
	DownMid   int
	DownRight int
	Step      int
}

// Units структура существа
// heals      hp               // здоровье
// y          maps.Coordinates // положение по y
// x          maps.Coordinates // положение по x
// liveStatus int              // статус 1 жив, 2 мертв
// color      color            // структура цвет RGB
// size       int              // размер существа
// energy     int              // количество энергии
// age        int              // возраст
type Units struct {
	Heals      hp               // здоровье
	Y          maps.Coordinates // положение по y
	X          maps.Coordinates // положение по x
	LiveStatus int              // статус 1 жив, 2 мертв
	Color      color            // структура цвет RGB
	Size       int              // размер существа
	Energy     int              // количество энергии
	Age        int              // возраст
	Memory     Memory
}

// Start - создает структуру существа с стартовыми статами
func NewUnits() *Units {
	var u = &Units{}
	u.Heals = DefaultHp
	u.Y = maps.Coordinates(f.Rand(int(maps.DefaultMap.X)))
	u.X = maps.Coordinates(f.Rand(int(maps.DefaultMap.X)))
	u.LiveStatus = DefaultLiveStatus
	u.Color = DefaultColor
	u.Size = DefaultSize
	u.Energy = DefaultEnergy
	u.Age = DefaultAge
	return u
}
