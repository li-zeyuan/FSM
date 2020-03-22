package fsm

import (
	"fmt"
	"sync"
)

type FSMState string
type FSMEvent string
type FSMHandler func() FSMState // 处理方法，返回新的状态

// 有限状态机
type FSM struct {
	mu sync.Mutex  // 排他锁
	CurState FSMState // 当前状态
	handlers map[FSMState]map[FSMEvent]FSMHandler // 处理地图集，每个状态都可以触发有限个事件，执行有限个处理
}

// 实例化FSM
func NewFSM(initState FSMState) *FSM {
	return &FSM{
		CurState: initState,
		handlers: make(map[FSMState]map[FSMEvent]FSMHandler),
	}
}

// 获取当前状态
func (f *FSM) getCurState() FSMState {
	return f.CurState
}

// 设置当前状态
func (f *FSM) setCurState(newState FSMState) {
	f.CurState = newState
}

// 某状态添加事件处理方法
func (f *FSM) AddHandler(state FSMState, event FSMEvent, handler FSMHandler) *FSM {
	if _, ok := f.handlers[state]; !ok {
		f.handlers[state] = make(map[FSMEvent]FSMHandler)
	}
	if _, ok := f.handlers[state][event]; ok {
		fmt.Printf("exist state：%s,event：%s\n", state, event)
		return f
	}
	f.handlers[state][event] = handler

	return f
}

// 事件处理
func (f *FSM) Call(event FSMEvent) FSMState {
	f.mu.Lock()
	defer f.mu.Unlock()
	events, ok := f.handlers[f.getCurState()]
	if !ok {
		return f.getCurState()
	}

	if fn, ok := events[event]; ok {
		oldState := f.getCurState()
		f.setCurState(fn())
		newState := f.getCurState()
		fmt.Printf("state from [%s] to [%s]\n", oldState, newState)
	}

	return f.getCurState()
}