package world

import (
	b "dontstop/being"
	m "dontstop/map"
	s "dontstop/statistics"
)

// Statistics
var Statistics = s.Statistics{
	Live:          0,
	Dead:          0,
	AllObject:     0,
	Predatory:     0,
	Herbivores:    0,
	Displacements: 0,
	CurrentStep:   0,
	Total:         0,
}

type SWorld struct {
	Status int
	Maps   m.Maps
	Sun    m.Sun
	Units  []b.Units
}

// World полностью мир
var World = SWorld{}

func initWorld(w *SWorld) {

	(*w).Status = 1
	(*w).Maps = m.Maps{
		X:      m.DefaultMap.X,
		Y:      m.DefaultMap.Y,
		Border: m.DefaultMap.Border,
	}

	(*w).Sun = m.Sun{
		X:         m.DefaultSun.X,
		Y:         m.DefaultSun.Y,
		Intensity: m.DefaultSun.Intensity,
		Radius:    m.DefaultSun.Radius,
	}
}

// CreateUnit - функция создания стартовых существ
func CreateUnit(w *SWorld, count int) {

	for i := 0; i < count; i++ {
		t := b.NewUnits()
		//fmt.Print(t)
		if serachCord(t.X, t.Y, w.Units) {
			(*w).Units = append(w.Units, *t)
		}
	}

}

// serachCord - функция проверки на занятость координат
func serachCord(x m.Coordinates, y m.Coordinates, w []b.Units) bool {
	for _, r := range w {
		if r.X == x && r.Y == y {
			return false
		}
	}
	return true
}

// Update - функция обновления мира
func Update() {
	if World.Status == 0 {
		initWorld(&World)
		CreateUnit(&World, 10)
	}
	//fmt.Print(World)
	World.Sun.MoveSun(World.Maps)
}

/*
Status: 0,
	Maps: m.Maps{
		X:      m.DefaultMap.X,
		Y:      m.DefaultMap.Y,
		Border: m.DefaultMap.Border,
	},
	Sun: m.Sun{
		X:         m.DefaultSun.X,
		Y:         m.DefaultSun.Y,
		Intensity: m.DefaultSun.Intensity,
		Radius:    m.DefaultSun.Radius,
	},
}
*/
