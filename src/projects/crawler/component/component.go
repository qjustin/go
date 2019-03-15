package component

import "sync/atomic"

type (
	IComponent interface {
		// 组件ID
		GetId() CID
		// 调用计数器
		GetCallCounter() uint64
		// 接受调用计数器
		GetAcceptCounter() uint64
		// 完成调用计数器
		GetCompleteCounter() uint64
		// 正在处理计数器
		GetProcessingCounter() uint64
		// 以JSON格式输出组件信息
		ToJSON() Component
		// 调用计数器+1
		IncreCallCounter()
		// 接受调用计数器+1
		IncreAcceptCounter()
		// 完成调用计数器+1
		IncreCompleteCounter()
		// 正在处理计数器+1
		IncreProcessingCounter()
		// 正在处理计数器-1
		DecreProcessingCounter()
		// 所有计数器清零
		ResetCounters()
	}
	Component struct {
		cid               CID    `json:"cid"`
		callCounter       uint64 `json:"callcounter"`
		acceptCounter     uint64 `json:"acceptcounter"`
		completeCounter   uint64 `json:"completecounter"`
		processingCounter uint64 `json:"processingcounter"`
	}
)

func (this *Component) GetId() CID {
	return this.cid
}

func (this *Component) GetCallCounter() uint64 {
	return atomic.LoadUint64(&this.callCounter)
}

func (this *Component) GetAcceptCounter() uint64 {
	return atomic.LoadUint64(&this.acceptCounter)
}

func (this *Component) GetCompleteCounter() uint64 {
	return atomic.LoadUint64(&this.completeCounter)
}

func (this *Component) GetProcessingCounter() uint64 {
	return atomic.LoadUint64(&this.processingCounter)
}

func (this *Component) ToJson() Component {
	return Component{
		cid:               this.GetId(),
		callCounter:       this.GetCallCounter(),
		acceptCounter:     this.GetAcceptCounter(),
		completeCounter:   this.GetCompleteCounter(),
		processingCounter: this.GetProcessingCounter(),
	}
}

func (this *Component) IncreCallCounter() {
	atomic.AddUint64(&this.callCounter, 1)
}

func (this *Component) IncreAcceptCounter() {
	atomic.AddUint64(&this.acceptCounter, 1)
}

func (this *Component) IncreCompleteCounter() {
	atomic.AddUint64(&this.completeCounter, 1)
}

func (this *Component) IncreProcessingCounter() {
	atomic.AddUint64(&this.processingCounter, 1)
}

func (this *Component) DecreProcessingCounter() {
	atomic.AddUint64(&this.processingCounter, ^uint64(0))
}

func (this *Component) ResetCounters() {
	atomic.StoreUint64(&this.callCounter, 0)
	atomic.StoreUint64(&this.acceptCounter, 0)
	atomic.StoreUint64(&this.completeCounter, 0)
	atomic.StoreUint64(&this.processingCounter, 0)
}
