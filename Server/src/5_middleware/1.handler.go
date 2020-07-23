package middleware

type middleWareType = func(functionType) functionType
type Handler struct {
	handler functionType
	middleWares []middleWareType
}
func NewHandler(handler functionType,middleWares []middleWareType) *Handler{
	return &Handler{handler,middleWares}
}

// 产生handler
func (h *Handler)Format() functionType {
	finalHandler := functionType(h.handler)
	for i:=0;i<len(h.middleWares);i++{
		finalHandler = h.middleWares[i](finalHandler)
	}
	return finalHandler
}









