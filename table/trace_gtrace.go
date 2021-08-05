// Code generated by gtrace. DO NOT EDIT.

package table

import (
	"context"
	"time"
)

// Compose returns a new ClientTrace which has functional fields composed
// both from t and x.
func (t ClientTrace) Compose(x ClientTrace) (ret ClientTrace) {
	switch {
	case t.OnCreateSession == nil:
		ret.OnCreateSession = x.OnCreateSession
	case x.OnCreateSession == nil:
		ret.OnCreateSession = t.OnCreateSession
	default:
		h1 := t.OnCreateSession
		h2 := x.OnCreateSession
		ret.OnCreateSession = func(c CreateSessionStartInfo) func(CreateSessionDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CreateSessionDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnKeepAlive == nil:
		ret.OnKeepAlive = x.OnKeepAlive
	case x.OnKeepAlive == nil:
		ret.OnKeepAlive = t.OnKeepAlive
	default:
		h1 := t.OnKeepAlive
		h2 := x.OnKeepAlive
		ret.OnKeepAlive = func(k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
			r1 := h1(k)
			r2 := h2(k)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(k KeepAliveDoneInfo) {
					r1(k)
					r2(k)
				}
			}
		}
	}
	switch {
	case t.OnDeleteSession == nil:
		ret.OnDeleteSession = x.OnDeleteSession
	case x.OnDeleteSession == nil:
		ret.OnDeleteSession = t.OnDeleteSession
	default:
		h1 := t.OnDeleteSession
		h2 := x.OnDeleteSession
		ret.OnDeleteSession = func(d DeleteSessionStartInfo) func(DeleteSessionDoneInfo) {
			r1 := h1(d)
			r2 := h2(d)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(d DeleteSessionDoneInfo) {
					r1(d)
					r2(d)
				}
			}
		}
	}
	switch {
	case t.OnPrepareDataQuery == nil:
		ret.OnPrepareDataQuery = x.OnPrepareDataQuery
	case x.OnPrepareDataQuery == nil:
		ret.OnPrepareDataQuery = t.OnPrepareDataQuery
	default:
		h1 := t.OnPrepareDataQuery
		h2 := x.OnPrepareDataQuery
		ret.OnPrepareDataQuery = func(p PrepareDataQueryStartInfo) func(PrepareDataQueryDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PrepareDataQueryDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	switch {
	case t.OnExecuteDataQuery == nil:
		ret.OnExecuteDataQuery = x.OnExecuteDataQuery
	case x.OnExecuteDataQuery == nil:
		ret.OnExecuteDataQuery = t.OnExecuteDataQuery
	default:
		h1 := t.OnExecuteDataQuery
		h2 := x.OnExecuteDataQuery
		ret.OnExecuteDataQuery = func(e ExecuteDataQueryStartInfo) func(ExecuteDataQueryDoneInfo) {
			r1 := h1(e)
			r2 := h2(e)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(e ExecuteDataQueryDoneInfo) {
					r1(e)
					r2(e)
				}
			}
		}
	}
	switch {
	case t.OnStreamReadTable == nil:
		ret.OnStreamReadTable = x.OnStreamReadTable
	case x.OnStreamReadTable == nil:
		ret.OnStreamReadTable = t.OnStreamReadTable
	default:
		h1 := t.OnStreamReadTable
		h2 := x.OnStreamReadTable
		ret.OnStreamReadTable = func(s StreamReadTableStartInfo) func(StreamReadTableDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamReadTableDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnStreamExecuteScanQuery == nil:
		ret.OnStreamExecuteScanQuery = x.OnStreamExecuteScanQuery
	case x.OnStreamExecuteScanQuery == nil:
		ret.OnStreamExecuteScanQuery = t.OnStreamExecuteScanQuery
	default:
		h1 := t.OnStreamExecuteScanQuery
		h2 := x.OnStreamExecuteScanQuery
		ret.OnStreamExecuteScanQuery = func(s StreamExecuteScanQueryStartInfo) func(StreamExecuteScanQueryDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamExecuteScanQueryDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnBeginTransaction == nil:
		ret.OnBeginTransaction = x.OnBeginTransaction
	case x.OnBeginTransaction == nil:
		ret.OnBeginTransaction = t.OnBeginTransaction
	default:
		h1 := t.OnBeginTransaction
		h2 := x.OnBeginTransaction
		ret.OnBeginTransaction = func(b BeginTransactionStartInfo) func(BeginTransactionDoneInfo) {
			r1 := h1(b)
			r2 := h2(b)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(b BeginTransactionDoneInfo) {
					r1(b)
					r2(b)
				}
			}
		}
	}
	switch {
	case t.OnCommitTransaction == nil:
		ret.OnCommitTransaction = x.OnCommitTransaction
	case x.OnCommitTransaction == nil:
		ret.OnCommitTransaction = t.OnCommitTransaction
	default:
		h1 := t.OnCommitTransaction
		h2 := x.OnCommitTransaction
		ret.OnCommitTransaction = func(c CommitTransactionStartInfo) func(CommitTransactionDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CommitTransactionDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	switch {
	case t.OnRollbackTransaction == nil:
		ret.OnRollbackTransaction = x.OnRollbackTransaction
	case x.OnRollbackTransaction == nil:
		ret.OnRollbackTransaction = t.OnRollbackTransaction
	default:
		h1 := t.OnRollbackTransaction
		h2 := x.OnRollbackTransaction
		ret.OnRollbackTransaction = func(r RollbackTransactionStartInfo) func(RollbackTransactionDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RollbackTransactionDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	return ret
}

type clientTraceContextKey struct{}

// WithClientTrace returns context which has associated ClientTrace with it.
func WithClientTrace(ctx context.Context, t ClientTrace) context.Context {
	return context.WithValue(ctx,
		clientTraceContextKey{},
		ContextClientTrace(ctx).Compose(t),
	)
}

// ContextClientTrace returns ClientTrace associated with ctx.
// If there is no ClientTrace associated with ctx then zero value
// of ClientTrace is returned.
func ContextClientTrace(ctx context.Context) ClientTrace {
	t, _ := ctx.Value(clientTraceContextKey{}).(ClientTrace)
	return t
}

func (t ClientTrace) onCreateSession(ctx context.Context, c1 CreateSessionStartInfo) func(CreateSessionDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(CreateSessionStartInfo) func(CreateSessionDoneInfo)
	switch {
	case t.OnCreateSession == nil:
		fn = c.OnCreateSession
	case c.OnCreateSession == nil:
		fn = t.OnCreateSession
	default:
		h1 := t.OnCreateSession
		h2 := c.OnCreateSession
		fn = func(c CreateSessionStartInfo) func(CreateSessionDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CreateSessionDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	if fn == nil {
		return func(CreateSessionDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(CreateSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onKeepAlive(ctx context.Context, k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(KeepAliveStartInfo) func(KeepAliveDoneInfo)
	switch {
	case t.OnKeepAlive == nil:
		fn = c.OnKeepAlive
	case c.OnKeepAlive == nil:
		fn = t.OnKeepAlive
	default:
		h1 := t.OnKeepAlive
		h2 := c.OnKeepAlive
		fn = func(k KeepAliveStartInfo) func(KeepAliveDoneInfo) {
			r1 := h1(k)
			r2 := h2(k)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(k KeepAliveDoneInfo) {
					r1(k)
					r2(k)
				}
			}
		}
	}
	if fn == nil {
		return func(KeepAliveDoneInfo) {
			return
		}
	}
	res := fn(k)
	if res == nil {
		return func(KeepAliveDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onDeleteSession(ctx context.Context, d DeleteSessionStartInfo) func(DeleteSessionDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(DeleteSessionStartInfo) func(DeleteSessionDoneInfo)
	switch {
	case t.OnDeleteSession == nil:
		fn = c.OnDeleteSession
	case c.OnDeleteSession == nil:
		fn = t.OnDeleteSession
	default:
		h1 := t.OnDeleteSession
		h2 := c.OnDeleteSession
		fn = func(d DeleteSessionStartInfo) func(DeleteSessionDoneInfo) {
			r1 := h1(d)
			r2 := h2(d)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(d DeleteSessionDoneInfo) {
					r1(d)
					r2(d)
				}
			}
		}
	}
	if fn == nil {
		return func(DeleteSessionDoneInfo) {
			return
		}
	}
	res := fn(d)
	if res == nil {
		return func(DeleteSessionDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onPrepareDataQuery(ctx context.Context, p PrepareDataQueryStartInfo) func(PrepareDataQueryDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(PrepareDataQueryStartInfo) func(PrepareDataQueryDoneInfo)
	switch {
	case t.OnPrepareDataQuery == nil:
		fn = c.OnPrepareDataQuery
	case c.OnPrepareDataQuery == nil:
		fn = t.OnPrepareDataQuery
	default:
		h1 := t.OnPrepareDataQuery
		h2 := c.OnPrepareDataQuery
		fn = func(p PrepareDataQueryStartInfo) func(PrepareDataQueryDoneInfo) {
			r1 := h1(p)
			r2 := h2(p)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(p PrepareDataQueryDoneInfo) {
					r1(p)
					r2(p)
				}
			}
		}
	}
	if fn == nil {
		return func(PrepareDataQueryDoneInfo) {
			return
		}
	}
	res := fn(p)
	if res == nil {
		return func(PrepareDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onExecuteDataQuery(ctx context.Context, e ExecuteDataQueryStartInfo) func(ExecuteDataQueryDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(ExecuteDataQueryStartInfo) func(ExecuteDataQueryDoneInfo)
	switch {
	case t.OnExecuteDataQuery == nil:
		fn = c.OnExecuteDataQuery
	case c.OnExecuteDataQuery == nil:
		fn = t.OnExecuteDataQuery
	default:
		h1 := t.OnExecuteDataQuery
		h2 := c.OnExecuteDataQuery
		fn = func(e ExecuteDataQueryStartInfo) func(ExecuteDataQueryDoneInfo) {
			r1 := h1(e)
			r2 := h2(e)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(e ExecuteDataQueryDoneInfo) {
					r1(e)
					r2(e)
				}
			}
		}
	}
	if fn == nil {
		return func(ExecuteDataQueryDoneInfo) {
			return
		}
	}
	res := fn(e)
	if res == nil {
		return func(ExecuteDataQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onStreamReadTable(ctx context.Context, s StreamReadTableStartInfo) func(StreamReadTableDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(StreamReadTableStartInfo) func(StreamReadTableDoneInfo)
	switch {
	case t.OnStreamReadTable == nil:
		fn = c.OnStreamReadTable
	case c.OnStreamReadTable == nil:
		fn = t.OnStreamReadTable
	default:
		h1 := t.OnStreamReadTable
		h2 := c.OnStreamReadTable
		fn = func(s StreamReadTableStartInfo) func(StreamReadTableDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamReadTableDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	if fn == nil {
		return func(StreamReadTableDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(StreamReadTableDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onStreamExecuteScanQuery(ctx context.Context, s StreamExecuteScanQueryStartInfo) func(StreamExecuteScanQueryDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(StreamExecuteScanQueryStartInfo) func(StreamExecuteScanQueryDoneInfo)
	switch {
	case t.OnStreamExecuteScanQuery == nil:
		fn = c.OnStreamExecuteScanQuery
	case c.OnStreamExecuteScanQuery == nil:
		fn = t.OnStreamExecuteScanQuery
	default:
		h1 := t.OnStreamExecuteScanQuery
		h2 := c.OnStreamExecuteScanQuery
		fn = func(s StreamExecuteScanQueryStartInfo) func(StreamExecuteScanQueryDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s StreamExecuteScanQueryDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	if fn == nil {
		return func(StreamExecuteScanQueryDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(StreamExecuteScanQueryDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onBeginTransaction(ctx context.Context, b BeginTransactionStartInfo) func(BeginTransactionDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(BeginTransactionStartInfo) func(BeginTransactionDoneInfo)
	switch {
	case t.OnBeginTransaction == nil:
		fn = c.OnBeginTransaction
	case c.OnBeginTransaction == nil:
		fn = t.OnBeginTransaction
	default:
		h1 := t.OnBeginTransaction
		h2 := c.OnBeginTransaction
		fn = func(b BeginTransactionStartInfo) func(BeginTransactionDoneInfo) {
			r1 := h1(b)
			r2 := h2(b)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(b BeginTransactionDoneInfo) {
					r1(b)
					r2(b)
				}
			}
		}
	}
	if fn == nil {
		return func(BeginTransactionDoneInfo) {
			return
		}
	}
	res := fn(b)
	if res == nil {
		return func(BeginTransactionDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onCommitTransaction(ctx context.Context, c1 CommitTransactionStartInfo) func(CommitTransactionDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(CommitTransactionStartInfo) func(CommitTransactionDoneInfo)
	switch {
	case t.OnCommitTransaction == nil:
		fn = c.OnCommitTransaction
	case c.OnCommitTransaction == nil:
		fn = t.OnCommitTransaction
	default:
		h1 := t.OnCommitTransaction
		h2 := c.OnCommitTransaction
		fn = func(c CommitTransactionStartInfo) func(CommitTransactionDoneInfo) {
			r1 := h1(c)
			r2 := h2(c)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(c CommitTransactionDoneInfo) {
					r1(c)
					r2(c)
				}
			}
		}
	}
	if fn == nil {
		return func(CommitTransactionDoneInfo) {
			return
		}
	}
	res := fn(c1)
	if res == nil {
		return func(CommitTransactionDoneInfo) {
			return
		}
	}
	return res
}
func (t ClientTrace) onRollbackTransaction(ctx context.Context, r RollbackTransactionStartInfo) func(RollbackTransactionDoneInfo) {
	c := ContextClientTrace(ctx)
	var fn func(RollbackTransactionStartInfo) func(RollbackTransactionDoneInfo)
	switch {
	case t.OnRollbackTransaction == nil:
		fn = c.OnRollbackTransaction
	case c.OnRollbackTransaction == nil:
		fn = t.OnRollbackTransaction
	default:
		h1 := t.OnRollbackTransaction
		h2 := c.OnRollbackTransaction
		fn = func(r RollbackTransactionStartInfo) func(RollbackTransactionDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RollbackTransactionDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	if fn == nil {
		return func(RollbackTransactionDoneInfo) {
			return
		}
	}
	res := fn(r)
	if res == nil {
		return func(RollbackTransactionDoneInfo) {
			return
		}
	}
	return res
}

// Compose returns a new RetryTrace which has functional fields composed
// both from t and x.
func (t RetryTrace) Compose(x RetryTrace) (ret RetryTrace) {
	switch {
	case t.OnLoop == nil:
		ret.OnLoop = x.OnLoop
	case x.OnLoop == nil:
		ret.OnLoop = t.OnLoop
	default:
		h1 := t.OnLoop
		h2 := x.OnLoop
		ret.OnLoop = func(r RetryLoopStartInfo) func(RetryLoopDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RetryLoopDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	return ret
}

type retryTraceContextKey struct{}

// WithRetryTrace returns context which has associated RetryTrace with it.
func WithRetryTrace(ctx context.Context, t RetryTrace) context.Context {
	return context.WithValue(ctx,
		retryTraceContextKey{},
		ContextRetryTrace(ctx).Compose(t),
	)
}

// ContextRetryTrace returns RetryTrace associated with ctx.
// If there is no RetryTrace associated with ctx then zero value
// of RetryTrace is returned.
func ContextRetryTrace(ctx context.Context) RetryTrace {
	t, _ := ctx.Value(retryTraceContextKey{}).(RetryTrace)
	return t
}

func (t RetryTrace) onLoop(ctx context.Context, r RetryLoopStartInfo) func(RetryLoopDoneInfo) {
	c := ContextRetryTrace(ctx)
	var fn func(RetryLoopStartInfo) func(RetryLoopDoneInfo)
	switch {
	case t.OnLoop == nil:
		fn = c.OnLoop
	case c.OnLoop == nil:
		fn = t.OnLoop
	default:
		h1 := t.OnLoop
		h2 := c.OnLoop
		fn = func(r RetryLoopStartInfo) func(RetryLoopDoneInfo) {
			r1 := h1(r)
			r2 := h2(r)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r RetryLoopDoneInfo) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	if fn == nil {
		return func(RetryLoopDoneInfo) {
			return
		}
	}
	res := fn(r)
	if res == nil {
		return func(RetryLoopDoneInfo) {
			return
		}
	}
	return res
}

// Compose returns a new SessionPoolTrace which has functional fields composed
// both from t and x.
func (t SessionPoolTrace) Compose(x SessionPoolTrace) (ret SessionPoolTrace) {
	switch {
	case t.OnGet == nil:
		ret.OnGet = x.OnGet
	case x.OnGet == nil:
		ret.OnGet = t.OnGet
	default:
		h1 := t.OnGet
		h2 := x.OnGet
		ret.OnGet = func(s SessionPoolGetStartInfo) func(SessionPoolGetDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolGetDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnWait == nil:
		ret.OnWait = x.OnWait
	case x.OnWait == nil:
		ret.OnWait = t.OnWait
	default:
		h1 := t.OnWait
		h2 := x.OnWait
		ret.OnWait = func(s SessionPoolWaitStartInfo) func(SessionPoolWaitDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolWaitDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnBusyCheck == nil:
		ret.OnBusyCheck = x.OnBusyCheck
	case x.OnBusyCheck == nil:
		ret.OnBusyCheck = t.OnBusyCheck
	default:
		h1 := t.OnBusyCheck
		h2 := x.OnBusyCheck
		ret.OnBusyCheck = func(s SessionPoolBusyCheckStartInfo) func(SessionPoolBusyCheckDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolBusyCheckDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnTake == nil:
		ret.OnTake = x.OnTake
	case x.OnTake == nil:
		ret.OnTake = t.OnTake
	default:
		h1 := t.OnTake
		h2 := x.OnTake
		ret.OnTake = func(s SessionPoolTakeStartInfo) func(SessionPoolTakeDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolTakeDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnTakeWait == nil:
		ret.OnTakeWait = x.OnTakeWait
	case x.OnTakeWait == nil:
		ret.OnTakeWait = t.OnTakeWait
	default:
		h1 := t.OnTakeWait
		h2 := x.OnTakeWait
		ret.OnTakeWait = func(s SessionPoolTakeWaitInfo) {
			h1(s)
			h2(s)
		}
	}
	switch {
	case t.OnPut == nil:
		ret.OnPut = x.OnPut
	case x.OnPut == nil:
		ret.OnPut = t.OnPut
	default:
		h1 := t.OnPut
		h2 := x.OnPut
		ret.OnPut = func(s SessionPoolPutStartInfo) func(SessionPoolPutDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolPutDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnClose == nil:
		ret.OnClose = x.OnClose
	case x.OnClose == nil:
		ret.OnClose = t.OnClose
	default:
		h1 := t.OnClose
		h2 := x.OnClose
		ret.OnClose = func(s SessionPoolCloseStartInfo) func(SessionPoolCloseDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolCloseDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	return ret
}

type sessionPoolTraceContextKey struct{}

// WithSessionPoolTrace returns context which has associated SessionPoolTrace with it.
func WithSessionPoolTrace(ctx context.Context, t SessionPoolTrace) context.Context {
	return context.WithValue(ctx,
		sessionPoolTraceContextKey{},
		ContextSessionPoolTrace(ctx).Compose(t),
	)
}

// ContextSessionPoolTrace returns SessionPoolTrace associated with ctx.
// If there is no SessionPoolTrace associated with ctx then zero value
// of SessionPoolTrace is returned.
func ContextSessionPoolTrace(ctx context.Context) SessionPoolTrace {
	t, _ := ctx.Value(sessionPoolTraceContextKey{}).(SessionPoolTrace)
	return t
}

func (t SessionPoolTrace) onGet(ctx context.Context, s SessionPoolGetStartInfo) func(SessionPoolGetDoneInfo) {
	c := ContextSessionPoolTrace(ctx)
	var fn func(SessionPoolGetStartInfo) func(SessionPoolGetDoneInfo)
	switch {
	case t.OnGet == nil:
		fn = c.OnGet
	case c.OnGet == nil:
		fn = t.OnGet
	default:
		h1 := t.OnGet
		h2 := c.OnGet
		fn = func(s SessionPoolGetStartInfo) func(SessionPoolGetDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolGetDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	if fn == nil {
		return func(SessionPoolGetDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolGetDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onWait(ctx context.Context, s SessionPoolWaitStartInfo) func(SessionPoolWaitDoneInfo) {
	c := ContextSessionPoolTrace(ctx)
	var fn func(SessionPoolWaitStartInfo) func(SessionPoolWaitDoneInfo)
	switch {
	case t.OnWait == nil:
		fn = c.OnWait
	case c.OnWait == nil:
		fn = t.OnWait
	default:
		h1 := t.OnWait
		h2 := c.OnWait
		fn = func(s SessionPoolWaitStartInfo) func(SessionPoolWaitDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolWaitDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	if fn == nil {
		return func(SessionPoolWaitDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolWaitDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onBusyCheck(ctx context.Context, s SessionPoolBusyCheckStartInfo) func(SessionPoolBusyCheckDoneInfo) {
	c := ContextSessionPoolTrace(ctx)
	var fn func(SessionPoolBusyCheckStartInfo) func(SessionPoolBusyCheckDoneInfo)
	switch {
	case t.OnBusyCheck == nil:
		fn = c.OnBusyCheck
	case c.OnBusyCheck == nil:
		fn = t.OnBusyCheck
	default:
		h1 := t.OnBusyCheck
		h2 := c.OnBusyCheck
		fn = func(s SessionPoolBusyCheckStartInfo) func(SessionPoolBusyCheckDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolBusyCheckDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	if fn == nil {
		return func(SessionPoolBusyCheckDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolBusyCheckDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onTake(ctx context.Context, s SessionPoolTakeStartInfo) func(SessionPoolTakeDoneInfo) {
	c := ContextSessionPoolTrace(ctx)
	var fn func(SessionPoolTakeStartInfo) func(SessionPoolTakeDoneInfo)
	switch {
	case t.OnTake == nil:
		fn = c.OnTake
	case c.OnTake == nil:
		fn = t.OnTake
	default:
		h1 := t.OnTake
		h2 := c.OnTake
		fn = func(s SessionPoolTakeStartInfo) func(SessionPoolTakeDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolTakeDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	if fn == nil {
		return func(SessionPoolTakeDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolTakeDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onTakeWait(ctx context.Context, s SessionPoolTakeWaitInfo) {
	c := ContextSessionPoolTrace(ctx)
	var fn func(SessionPoolTakeWaitInfo)
	switch {
	case t.OnTakeWait == nil:
		fn = c.OnTakeWait
	case c.OnTakeWait == nil:
		fn = t.OnTakeWait
	default:
		h1 := t.OnTakeWait
		h2 := c.OnTakeWait
		fn = func(s SessionPoolTakeWaitInfo) {
			h1(s)
			h2(s)
		}
	}
	if fn == nil {
		return
	}
	fn(s)
}
func (t SessionPoolTrace) onPut(ctx context.Context, s SessionPoolPutStartInfo) func(SessionPoolPutDoneInfo) {
	c := ContextSessionPoolTrace(ctx)
	var fn func(SessionPoolPutStartInfo) func(SessionPoolPutDoneInfo)
	switch {
	case t.OnPut == nil:
		fn = c.OnPut
	case c.OnPut == nil:
		fn = t.OnPut
	default:
		h1 := t.OnPut
		h2 := c.OnPut
		fn = func(s SessionPoolPutStartInfo) func(SessionPoolPutDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolPutDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	if fn == nil {
		return func(SessionPoolPutDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolPutDoneInfo) {
			return
		}
	}
	return res
}
func (t SessionPoolTrace) onClose(ctx context.Context, s SessionPoolCloseStartInfo) func(SessionPoolCloseDoneInfo) {
	c := ContextSessionPoolTrace(ctx)
	var fn func(SessionPoolCloseStartInfo) func(SessionPoolCloseDoneInfo)
	switch {
	case t.OnClose == nil:
		fn = c.OnClose
	case c.OnClose == nil:
		fn = t.OnClose
	default:
		h1 := t.OnClose
		h2 := c.OnClose
		fn = func(s SessionPoolCloseStartInfo) func(SessionPoolCloseDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s SessionPoolCloseDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	if fn == nil {
		return func(SessionPoolCloseDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(SessionPoolCloseDoneInfo) {
			return
		}
	}
	return res
}

// Compose returns a new createSessionTrace which has functional fields composed
// both from t and x.
func (t createSessionTrace) Compose(x createSessionTrace) (ret createSessionTrace) {
	switch {
	case t.OnCheckEnoughSpace == nil:
		ret.OnCheckEnoughSpace = x.OnCheckEnoughSpace
	case x.OnCheckEnoughSpace == nil:
		ret.OnCheckEnoughSpace = t.OnCheckEnoughSpace
	default:
		h1 := t.OnCheckEnoughSpace
		h2 := x.OnCheckEnoughSpace
		ret.OnCheckEnoughSpace = func(enoughSpace bool) {
			h1(enoughSpace)
			h2(enoughSpace)
		}
	}
	switch {
	case t.OnCreateSessionGoroutineStart == nil:
		ret.OnCreateSessionGoroutineStart = x.OnCreateSessionGoroutineStart
	case x.OnCreateSessionGoroutineStart == nil:
		ret.OnCreateSessionGoroutineStart = t.OnCreateSessionGoroutineStart
	default:
		h1 := t.OnCreateSessionGoroutineStart
		h2 := x.OnCreateSessionGoroutineStart
		ret.OnCreateSessionGoroutineStart = func() func(createSessionResult) {
			r1 := h1()
			r2 := h2()
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(r createSessionResult) {
					r1(r)
					r2(r)
				}
			}
		}
	}
	switch {
	case t.OnStartSelect == nil:
		ret.OnStartSelect = x.OnStartSelect
	case x.OnStartSelect == nil:
		ret.OnStartSelect = t.OnStartSelect
	default:
		h1 := t.OnStartSelect
		h2 := x.OnStartSelect
		ret.OnStartSelect = func() {
			h1()
			h2()
		}
	}
	switch {
	case t.OnReadResult == nil:
		ret.OnReadResult = x.OnReadResult
	case x.OnReadResult == nil:
		ret.OnReadResult = t.OnReadResult
	default:
		h1 := t.OnReadResult
		h2 := x.OnReadResult
		ret.OnReadResult = func(r createSessionResult) {
			h1(r)
			h2(r)
		}
	}
	switch {
	case t.OnContextDone == nil:
		ret.OnContextDone = x.OnContextDone
	case x.OnContextDone == nil:
		ret.OnContextDone = t.OnContextDone
	default:
		h1 := t.OnContextDone
		h2 := x.OnContextDone
		ret.OnContextDone = func() {
			h1()
			h2()
		}
	}
	switch {
	case t.OnPutSession == nil:
		ret.OnPutSession = x.OnPutSession
	case x.OnPutSession == nil:
		ret.OnPutSession = t.OnPutSession
	default:
		h1 := t.OnPutSession
		h2 := x.OnPutSession
		ret.OnPutSession = func(session *Session, err error) {
			h1(session, err)
			h2(session, err)
		}
	}
	return ret
}
func (t createSessionTrace) onCheckEnoughSpace(enoughSpace bool) {
	fn := t.OnCheckEnoughSpace
	if fn == nil {
		return
	}
	fn(enoughSpace)
}
func (t createSessionTrace) onCreateSessionGoroutineStart() func(createSessionResult) {
	fn := t.OnCreateSessionGoroutineStart
	if fn == nil {
		return func(createSessionResult) {
			return
		}
	}
	res := fn()
	if res == nil {
		return func(createSessionResult) {
			return
		}
	}
	return res
}
func (t createSessionTrace) onStartSelect() {
	fn := t.OnStartSelect
	if fn == nil {
		return
	}
	fn()
}
func (t createSessionTrace) onReadResult(r createSessionResult) {
	fn := t.OnReadResult
	if fn == nil {
		return
	}
	fn(r)
}
func (t createSessionTrace) onContextDone() {
	fn := t.OnContextDone
	if fn == nil {
		return
	}
	fn()
}
func (t createSessionTrace) onPutSession(session *Session, err error) {
	fn := t.OnPutSession
	if fn == nil {
		return
	}
	fn(session, err)
}
func clientTraceOnCreateSession(ctx context.Context, t ClientTrace, c context.Context) func(_ context.Context, _ *Session, endpoint string, latency time.Duration, _ error) {
	var p CreateSessionStartInfo
	p.Context = c
	res := t.onCreateSession(ctx, p)
	return func(c context.Context, s *Session, endpoint string, latency time.Duration, e error) {
		var p CreateSessionDoneInfo
		p.Context = c
		p.Session = s
		p.Endpoint = endpoint
		p.Latency = latency
		p.Error = e
		res(p)
	}
}
func clientTraceOnKeepAlive(ctx context.Context, t ClientTrace, c context.Context, s *Session) func(context.Context, *Session, SessionInfo, error) {
	var p KeepAliveStartInfo
	p.Context = c
	p.Session = s
	res := t.onKeepAlive(ctx, p)
	return func(c context.Context, s *Session, s1 SessionInfo, e error) {
		var p KeepAliveDoneInfo
		p.Context = c
		p.Session = s
		p.SessionInfo = s1
		p.Error = e
		res(p)
	}
}
func clientTraceOnDeleteSession(ctx context.Context, t ClientTrace, c context.Context, s *Session) func(_ context.Context, _ *Session, latency time.Duration, _ error) {
	var p DeleteSessionStartInfo
	p.Context = c
	p.Session = s
	res := t.onDeleteSession(ctx, p)
	return func(c context.Context, s *Session, latency time.Duration, e error) {
		var p DeleteSessionDoneInfo
		p.Context = c
		p.Session = s
		p.Latency = latency
		p.Error = e
		res(p)
	}
}
func clientTraceOnPrepareDataQuery(ctx context.Context, t ClientTrace, c context.Context, s *Session, query string) func(_ context.Context, _ *Session, query string, result *DataQuery, cached bool, _ error) {
	var p PrepareDataQueryStartInfo
	p.Context = c
	p.Session = s
	p.Query = query
	res := t.onPrepareDataQuery(ctx, p)
	return func(c context.Context, s *Session, query string, result *DataQuery, cached bool, e error) {
		var p PrepareDataQueryDoneInfo
		p.Context = c
		p.Session = s
		p.Query = query
		p.Result = result
		p.Cached = cached
		p.Error = e
		res(p)
	}
}
func clientTraceOnExecuteDataQuery(ctx context.Context, t ClientTrace, c context.Context, s *Session, txID string, query *DataQuery, parameters *QueryParameters) func(_ context.Context, _ *Session, txID string, query *DataQuery, parameters *QueryParameters, prepared bool, _ *Result, _ error) {
	var p ExecuteDataQueryStartInfo
	p.Context = c
	p.Session = s
	p.TxID = txID
	p.Query = query
	p.Parameters = parameters
	res := t.onExecuteDataQuery(ctx, p)
	return func(c context.Context, s *Session, txID string, query *DataQuery, parameters *QueryParameters, prepared bool, r *Result, e error) {
		var p ExecuteDataQueryDoneInfo
		p.Context = c
		p.Session = s
		p.TxID = txID
		p.Query = query
		p.Parameters = parameters
		p.Prepared = prepared
		p.Result = r
		p.Error = e
		res(p)
	}
}
func clientTraceOnStreamReadTable(ctx context.Context, t ClientTrace, c context.Context, s *Session) func(context.Context, *Session, *Result, error) {
	var p StreamReadTableStartInfo
	p.Context = c
	p.Session = s
	res := t.onStreamReadTable(ctx, p)
	return func(c context.Context, s *Session, r *Result, e error) {
		var p StreamReadTableDoneInfo
		p.Context = c
		p.Session = s
		p.Result = r
		p.Error = e
		res(p)
	}
}
func clientTraceOnStreamExecuteScanQuery(ctx context.Context, t ClientTrace, c context.Context, s *Session, query *DataQuery, parameters *QueryParameters) func(_ context.Context, _ *Session, query *DataQuery, parameters *QueryParameters, _ *Result, _ error) {
	var p StreamExecuteScanQueryStartInfo
	p.Context = c
	p.Session = s
	p.Query = query
	p.Parameters = parameters
	res := t.onStreamExecuteScanQuery(ctx, p)
	return func(c context.Context, s *Session, query *DataQuery, parameters *QueryParameters, r *Result, e error) {
		var p StreamExecuteScanQueryDoneInfo
		p.Context = c
		p.Session = s
		p.Query = query
		p.Parameters = parameters
		p.Result = r
		p.Error = e
		res(p)
	}
}
func clientTraceOnBeginTransaction(ctx context.Context, t ClientTrace, c context.Context, s *Session) func(_ context.Context, _ *Session, txID string, _ error) {
	var p BeginTransactionStartInfo
	p.Context = c
	p.Session = s
	res := t.onBeginTransaction(ctx, p)
	return func(c context.Context, s *Session, txID string, e error) {
		var p BeginTransactionDoneInfo
		p.Context = c
		p.Session = s
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func clientTraceOnCommitTransaction(ctx context.Context, t ClientTrace, c context.Context, s *Session, txID string) func(_ context.Context, _ *Session, txID string, _ error) {
	var p CommitTransactionStartInfo
	p.Context = c
	p.Session = s
	p.TxID = txID
	res := t.onCommitTransaction(ctx, p)
	return func(c context.Context, s *Session, txID string, e error) {
		var p CommitTransactionDoneInfo
		p.Context = c
		p.Session = s
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func clientTraceOnRollbackTransaction(ctx context.Context, t ClientTrace, c context.Context, s *Session, txID string) func(_ context.Context, _ *Session, txID string, _ error) {
	var p RollbackTransactionStartInfo
	p.Context = c
	p.Session = s
	p.TxID = txID
	res := t.onRollbackTransaction(ctx, p)
	return func(c context.Context, s *Session, txID string, e error) {
		var p RollbackTransactionDoneInfo
		p.Context = c
		p.Session = s
		p.TxID = txID
		p.Error = e
		res(p)
	}
}
func retryTraceOnLoop(ctx context.Context, t RetryTrace, c context.Context) func(_ context.Context, latency time.Duration, attempts int) {
	var p RetryLoopStartInfo
	p.Context = c
	res := t.onLoop(ctx, p)
	return func(c context.Context, latency time.Duration, attempts int) {
		var p RetryLoopDoneInfo
		p.Context = c
		p.Latency = latency
		p.Attempts = attempts
		res(p)
	}
}
func sessionPoolTraceOnGet(ctx context.Context, t SessionPoolTrace, c context.Context) func(_ context.Context, _ *Session, latency time.Duration, retryAttempts int, _ error) {
	var p SessionPoolGetStartInfo
	p.Context = c
	res := t.onGet(ctx, p)
	return func(c context.Context, s *Session, latency time.Duration, retryAttempts int, e error) {
		var p SessionPoolGetDoneInfo
		p.Context = c
		p.Session = s
		p.Latency = latency
		p.RetryAttempts = retryAttempts
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnWait(ctx context.Context, t SessionPoolTrace, c context.Context) func(context.Context, *Session, error) {
	var p SessionPoolWaitStartInfo
	p.Context = c
	res := t.onWait(ctx, p)
	return func(c context.Context, s *Session, e error) {
		var p SessionPoolWaitDoneInfo
		p.Context = c
		p.Session = s
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnBusyCheck(ctx context.Context, t SessionPoolTrace, c context.Context, s *Session) func(_ context.Context, _ *Session, reused bool, _ error) {
	var p SessionPoolBusyCheckStartInfo
	p.Context = c
	p.Session = s
	res := t.onBusyCheck(ctx, p)
	return func(c context.Context, s *Session, reused bool, e error) {
		var p SessionPoolBusyCheckDoneInfo
		p.Context = c
		p.Session = s
		p.Reused = reused
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnTake(ctx context.Context, t SessionPoolTrace, c context.Context, s *Session) func(_ context.Context, _ *Session, took bool, _ error) {
	var p SessionPoolTakeStartInfo
	p.Context = c
	p.Session = s
	res := t.onTake(ctx, p)
	return func(c context.Context, s *Session, took bool, e error) {
		var p SessionPoolTakeDoneInfo
		p.Context = c
		p.Session = s
		p.Took = took
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnTakeWait(ctx context.Context, t SessionPoolTrace, c context.Context, s *Session) {
	var p SessionPoolTakeWaitInfo
	p.Context = c
	p.Session = s
	t.onTakeWait(ctx, p)
}
func sessionPoolTraceOnPut(ctx context.Context, t SessionPoolTrace, c context.Context, s *Session) func(context.Context, *Session, error) {
	var p SessionPoolPutStartInfo
	p.Context = c
	p.Session = s
	res := t.onPut(ctx, p)
	return func(c context.Context, s *Session, e error) {
		var p SessionPoolPutDoneInfo
		p.Context = c
		p.Session = s
		p.Error = e
		res(p)
	}
}
func sessionPoolTraceOnClose(ctx context.Context, t SessionPoolTrace, c context.Context) func(context.Context, error) {
	var p SessionPoolCloseStartInfo
	p.Context = c
	res := t.onClose(ctx, p)
	return func(c context.Context, e error) {
		var p SessionPoolCloseDoneInfo
		p.Context = c
		p.Error = e
		res(p)
	}
}
func createSessionTraceOnCheckEnoughSpace(t createSessionTrace, enoughSpace bool) {
	t.onCheckEnoughSpace(enoughSpace)
}
func createSessionTraceOnCreateSessionGoroutineStart(t createSessionTrace) func() {
	res := t.onCreateSessionGoroutineStart()
	return func() {
		var p createSessionResult
		res(p)
	}
}
func createSessionTraceOnStartSelect(t createSessionTrace) {
	t.onStartSelect()
}
func createSessionTraceOnReadResult(t createSessionTrace) {
	var p createSessionResult
	t.onReadResult(p)
}
func createSessionTraceOnContextDone(t createSessionTrace) {
	t.onContextDone()
}
func createSessionTraceOnPutSession(t createSessionTrace, session *Session, err error) {
	t.onPutSession(session, err)
}
