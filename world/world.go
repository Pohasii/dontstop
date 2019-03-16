package world

import (
	m "dontstop/map"
	//s "../statistics/statistics"
	//"../being/being"
)

/*
var maps maps.Maps = maps.Maps {
	x = 100,
	y = 100,
	border = 150,
}

var statistics s.Statistics = s.Statistics {
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0
}

var sun maps.Sun = maps.Sun{
	x = 0,
	y = 125,
	intensity = 15,
	radius   = 50,
}
*/
type sWorld struct {
	Maps m.Maps
	Sun  m.Sun
}

var World = sWorld{
	Maps: m.Maps{
		X:      100,
		Y:      100,
		Border: 150,
	},
	Sun: m.Sun{
		X:         (0),
		Y:         125,
		Intensity: 15,
		Radius:    50,
	},
}

func Update() {
	World.Sun.MoveSun(World.Maps)
}

// var units []being.Units = make( []being.Units, 1)

/*

if(l<200 && t==0) div.style.left = l+1+'px';
08
    if(l==200 && t<200) div.style.top = t+1+'px';
09
    if(l>0 && t==200) div.style.left = l-1+'px';
10
    if(l==0 && t>0) div.style.top = t-1+'px';


var currentAngle = 0; // текущее значение угла
var radius = 100; // радиус окружности
var baseX = 200; // x координата центра окружности
var baseY = 200; // y координата центра окружности

// считаем косинус текущего значения угла
    // и умножаем на значение радиуса
    var vx = Math.cos(currentAngle) * radius;

    // считаем синус текущего значения угла
    // и умножаем на значение радиуса
    var vy = Math.sin(currentAngle) * radius;

    context.fillStyle = '#fff'; // устанавливаем цвет заливки в белый

    context.fillRect(baseX + vx, baseY + vy, 5, 5);
    // считем новую позицию по x и y относительно центра окружности
    // и заливаем ее квадратиком высотой и шириной в 5 пикселей

    // увеличиваем значение угла
		currentAngle += 0.1;
*/
