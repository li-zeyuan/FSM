package main

import (
	"fmt"
	"fsm"
)

var (
	PowerOff = fsm.FSMState("关闭")
	FirstGear = fsm.FSMState("1档")
	SecondGear = fsm.FSMState("2档")
	ThirdGear = fsm.FSMState("3档")

	PowerOffEvent = fsm.FSMEvent("按下关闭按钮")
	FirstGearEvent = fsm.FSMEvent("按下1挡按钮")
	SecondGearEvent = fsm.FSMEvent("按下2挡按钮")
	ThirdGearEvent = fsm.FSMEvent("按下3挡按钮")

	PowerOffHandler = fsm.FSMHandler(func() fsm.FSMState {
		fmt.Println("关闭")
		return PowerOff
	})
	FirstGearHandler = fsm.FSMHandler(func() fsm.FSMState {
		fmt.Println("开启1挡")
		return FirstGear
	})
	SecondGearHandler = fsm.FSMHandler(func() fsm.FSMState {
		fmt.Println("开启2挡")
		return SecondGear
	})
	ThirdGearHandler = fsm.FSMHandler(func() fsm.FSMState {
		fmt.Println("开启3挡")
		return ThirdGear
	})
)

func main()  {
	efan := fsm.NewFSM(PowerOff) // 初始状态是关闭

	// 关闭状态
	efan.AddHandler(PowerOff, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(PowerOff, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(PowerOff, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(PowerOff, ThirdGearEvent, ThirdGearHandler)
	// 1档状态
	efan.AddHandler(FirstGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(FirstGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(FirstGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(FirstGear, ThirdGearEvent, ThirdGearHandler)
	// 2档状态
	efan.AddHandler(SecondGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(SecondGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(SecondGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(SecondGear, ThirdGearEvent, ThirdGearHandler)
	// 3档状态
	efan.AddHandler(ThirdGear, PowerOffEvent, PowerOffHandler)
	efan.AddHandler(ThirdGear, FirstGearEvent, FirstGearHandler)
	efan.AddHandler(ThirdGear, SecondGearEvent, SecondGearHandler)
	efan.AddHandler(ThirdGear, ThirdGearEvent, ThirdGearHandler)

	// 开始测试状态变化
	efan.Call(ThirdGearEvent)  // 按下3档按钮
	efan.Call(FirstGearEvent)  // 按下1档按钮
	efan.Call(PowerOffEvent)   // 按下关闭按钮
	efan.Call(SecondGearEvent) // 按下2档按钮
	efan.Call(PowerOffEvent)   // 按下关闭按钮
}
