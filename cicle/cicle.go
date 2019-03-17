package cicle

import (
	"dontstop/world"
	"fmt"
	"time"
)

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
	var step float64 // шаг

	// количество кадров
	fps := time.Tick(time.Duration(countFps) * time.Millisecond)

	// сам процесс
	for now := range fps {
		// fmt.Printf("%v %s\n", now)

		// считаем шаги
		step++

		// тут делаем обновления сервера
		world.Update()

		// тут отправляем состояние сервера на клиент
		// now передается только ради использования.
		sender(step, now)
	}

}

// update функция где будут проходить обновление состояния мира
// func update(step float64) {
//
// }

// sender функция где будет отправляться состояние мира на клиен
func sender(step float64, time time.Time) {

	fmt.Println(" value:", step)
	fmt.Println(" Sun:", world.World.Sun)
	fmt.Println(" full map:", world.World)
}
