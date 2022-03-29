package httpabstracts

type WebSocketContext struct {
}

type WebSocketManager interface {
	Accept(ctx *WebSocketContext)
}
