package statistics

// Statistics структура статистики
//	live          int живых
//	dead          int мертвых
//	allObject     int всего на карте сейчас
//	predatory     int хищных
//	herbivores    int травоядных
//	displacements int перемещений
//	currentStep   int текущий шаг
//	total         int всего количество существ за все время
type Statistics struct {
	Live          int // живых
	Dead          int // мертвых
	AllObject     int // всего на карте сейчас
	Predatory     int // хищных
	Herbivores    int // травоядных
	Displacements int // перемещений
	CurrentStep   int // текущий шаг
	Total         int // всего количество существ за все время
}
