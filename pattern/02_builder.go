package main

import "fmt"

/*
Строитель - это порождающий паттерн, который позволяет создавать сложные обьекты пошагово,
и использовать один и тот же код для получения разных представлений обьекта


 Шаги реализации
1. Убедитесь в том, что создание разных представлений
объекта можно свести к общим шагам.
2. Описать эти шаги в общем интерфейсе строителей.
3. Для каждого из представлений объекта-продукта создать по одному
классу-строителю и реализовать их методы строительства.
4. Подумать о создании класса директора. Его методы будут создавать
различные конфигурации продуктов, вызывая разные шаги одного и того же строителя.
5. Клиентский код должен будет создавать и объекты строителей,
и объект директора. Перед началом строительства, клиент должен связать
определённого строителя с директором. Это можно сделать либо через конструктор,
либо через сеттер, либо подав строителя напрямую в в строительный метод директора.
6. Результат строительства можно вернуть из директора, но только если метод
возврата продукта удалось поместить в общий интерфейс строителей.
Иначе, вы жёстко привяжете директора к конкретным классам строителей.

Примущества
	 Позволяет создавать продукты пошагово.
	 Позволяет использовать один и тот же код для создания различных продуктов.
	 Изолирует сложный код сборки продукта от его основной бизнес-логики.

Недостатки
	 Усложняет код программы за счёт дополнительных классов.
	 Клиент будет привязан к конкретным классам строителей,
	так как в интерфейсе строителя может не быть метода получения результата.
*/

type Builder interface {
	setWall(count int)
	setCeiling(str string)
	setWindow(count int)
	setDoor(door bool)
}

type BuildHouse struct {
	Wall    int
	Celling string
	Window  int
	Door    bool
}

func (b *BuildHouse) setWall(count int) {
	b.Wall = count
}

func (b *BuildHouse) setCeiling(str string) {
	b.Celling = str
}

func (b *BuildHouse) setWindow(count int) {
	b.Window = count
}

func (b *BuildHouse) setDoor(door bool) {
	b.Door = door
}

func (b *BuildHouse) getHouse() BuildHouse {
	return *b
}

type InfoHouse struct {
	Wall    int
	Celling string
	Window  int
	Door    bool
}

func (i *InfoHouse) setWall(count int) {
	i.Wall = count
}

func (i *InfoHouse) setCeiling(str string) {
	i.Celling = str
}

func (i *InfoHouse) setWindow(count int) {
	i.Window = count
}

func (i *InfoHouse) setDoor(door bool) {
	i.Door = door
}

func (i *InfoHouse) getInfoOfHouse() InfoHouse {
	return *i
}

type Director struct {
	b Builder // интерфейс
}

func (director *Director) fillBuilder(count int, str string, count2 int, door bool) {
	director.b.setWall(count)
	director.b.setDoor(door)
	director.b.setWindow(count2)
	director.b.setCeiling(str)
}

func main() {
	bH := &BuildHouse{}
	//bH.getHouse()
	dir := Director{
		b: bH, //структура для стоительства дома
	}
	dir.fillBuilder(4, "flat", 2, true)
	fmt.Println(bH.getHouse())
	dir = Director{
		b: &InfoHouse{}, // передаем стркутуру для заполнения информации
	}
	dir.fillBuilder(4, "flat", 2, true)
	fmt.Println(dir.b)
}
