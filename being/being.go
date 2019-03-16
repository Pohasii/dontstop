package being

import "../map/maps"

// hp - тип данных для здоровья
// float64
// используется в структуре существа
type hp float64

// color это RGB цвет существа
// one   int первый
// two   int второй
// three int третий
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
type Colors struct {
	a1 color = color{0,255,0}
	a2 color = color{85,212,0}
	a3 color = color{128,170,0}
	a4 color = color{170,128,0}
	a5 color = color{212,85,0}
	a6 color = color{255,0,0}
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
	heals      hp               // здоровье
	y          maps.Coordinates // положение по y
	x          maps.Coordinates // положение по x
	liveStatus int              // статус 1 жив, 2 мертв
	color      color            // структура цвет RGB
	size       int              // размер существа
	energy     int              // количество энергии
	age        int              // возраст
}
