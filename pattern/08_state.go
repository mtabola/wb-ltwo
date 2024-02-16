package main

import (
	"fmt"
)

/*
	Состояние - это поведенческий паттерн, который позволяет объектам изменять поведение тогда,
	когда внутренее состояние изменено.

	Преимущества
	 Избавляет от множества больших условных операторов
	машины состояний.
	 Концентрирует в одном месте код, связанный с определённым состоянием.
	 Упрощает код контекста.
	 Недостатки
	 Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type PlayerI interface {
	next(context *State)
	previous(context *State)
	play(context *State)
	lock(context *State)
}

type pauseState struct {
	s State
}

func (p *pauseState) next(context *State) {
	context.idTrack++
	fmt.Println("трек остановлен ++ ")
}
func (p *pauseState) previous(context *State) {
	context.idTrack--
	fmt.Println("трек остановлен -- ")
}
func (p *pauseState) play(context *State) {
	context.setState(&playState{})
	context.off = true
	fmt.Println("Play")
}
func (p *pauseState) lock(context *State) {
	context.setState(&lockState{})
	context.off = false
	fmt.Println("lock")
}

type playState struct {
	s State
}

func (p *playState) next(context *State) {
	context.idTrack++
	fmt.Println("туц туц ")
}
func (p *playState) previous(context *State) {
	context.idTrack--
	fmt.Println("цут цут ")
}

func (p *playState) play(context *State) {
	context.setState(&pauseState{})
	context.off = true
	fmt.Println("stop")
}
func (p *playState) lock(context *State) {
	context.setState(&lockState{})
	context.off = true
	fmt.Println("lock")
}

type lockState struct {
	s State
}

func (p *lockState) next(context *State) {
	fmt.Println("...")
}
func (p *lockState) previous(context *State) {
	fmt.Println("...")
}
func (p *lockState) play(context *State) {
	fmt.Println("...")
}
func (p *lockState) lock(context *State) {
	context.setState(&pauseState{})
	fmt.Println("unlock")
}

type State struct {
	player  PlayerI
	idTrack int
	off     bool
}

func (s *State) setState(p PlayerI) {
	s.player = p
}

func (s *State) next() {
	s.player.next(s)
}
func (s *State) previous() {
	s.player.previous(s)
}
func (s *State) play() {
	s.player.play(s)
}
func (s *State) lock() {
	s.player.lock(s)
}

func main() {
	pp := State{}
	pp.setState(&pauseState{})
	pp.play()
	pp.next()
	pp.next()
	pp.next()
	fmt.Println(pp.idTrack)
	pp.play()
	fmt.Println(pp.off)
	pp.lock()
	pp.lock()
	pp.play()
	pp.previous()
	fmt.Println(pp.idTrack)
}
