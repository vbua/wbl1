package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

/*
Адаптер — это структурный паттерн проектирования, который позволяет объектам с несовместимыми интерфейсами
работать вместе. Адаптер выступает прослойкой между двумя объектами, превращая вызовы одного в вызовы понятные другому.
*/

// юзер, который будет втыкать вилку в розетку
type User struct{}

func (u *User) plugEuConnectorIntoSocket(s Socket) {
	fmt.Println("Plugging EU connector into socket")
	s.plugIntoEuSocket()
}

// интерфейс европейской розетки
type Socket interface {
	plugIntoEuSocket()
}

type EuSocket struct{}

func (es *EuSocket) plugIntoEuSocket() {
	fmt.Println("Сonnector is plugged into EU socket")
}

// американская розетка не поддерживает европейскую вилку
type AmericanSocket struct{}

func (as *AmericanSocket) plugIntoAmericanSocket() {
	fmt.Println("Сonnector is plugged into American socket")
}

type AmericanSocketAdapter struct {
	AmericanSocket *AmericanSocket
}

func (aca *AmericanSocketAdapter) plugIntoEuSocket() {
	fmt.Println("Adapter converts EU connector into American connector")
	aca.AmericanSocket.plugIntoAmericanSocket()
}

func newAmericanSocketAdapter(as *AmericanSocket) AmericanSocketAdapter {
	return AmericanSocketAdapter{
		AmericanSocket: as,
	}
}

func main() {
	user := User{}
	euSocket := EuSocket{}
	amSocket := AmericanSocket{}

	user.plugEuConnectorIntoSocket(&euSocket)

	amSocketAdapter := newAmericanSocketAdapter(&amSocket)
	user.plugEuConnectorIntoSocket(&amSocketAdapter)
}
