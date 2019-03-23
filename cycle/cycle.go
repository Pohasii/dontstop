package cycle

import (
	"dontstop/world"
	"time"
)

// step количество шагов сделнных сервером
var Step float64 // шаг

// frame функция жизненного цикла сервера
func Frame(thisWorld world.SWorld, speed float64) {

	// Setting

	// получить текущее время
	//var t float64 = float64(time.Now().Millisecond)

	// countFps количество тиков, которое сделает цикл
	// время в миллисекундах
	// тип float64
	var countFps float64 = 1000 / speed // частота кадров

	// step количество шагов сделнных сервером
	// var Step float64 // шаг

	// количество кадров
	fps := time.Tick(time.Duration(countFps) * time.Millisecond)

	// сам процесс
	for now := range fps {
		// fmt.Printf("%v %s\n", now)

		// считаем шаги
		Step++

		// тут делаем обновления сервера
		world.Update()

		// тут отправляем состояние сервера на клиент
		// now передается только ради использования.
		sender(Step, now)
	}

}

// update функция где будут проходить обновление состояния мира
// func update(step float64) {
//
// }

// sender функция где будет отправляться состояние мира на клиен
func sender(Step float64, time time.Time) {

	//fmt.Println(" value:", Step)
	//fmt.Println(" Sun:", world.World.Sun)
	//fmt.Println(" full map:", world.World)
}
