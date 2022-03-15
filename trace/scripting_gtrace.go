// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// Compose returns a new Scripting which has functional fields composed
// both from t and x.
func (t Scripting) Compose(x Scripting) (ret Scripting) {
	switch {
	case t.OnExecute == nil:
		ret.OnExecute = x.OnExecute
	case x.OnExecute == nil:
		ret.OnExecute = t.OnExecute
	default:
		h1 := t.OnExecute
		h2 := x.OnExecute
		ret.OnExecute = func(s ScriptingExecuteStartInfo) func(ScriptingExecuteDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s ScriptingExecuteDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	switch {
	case t.OnStreamExecute == nil:
		ret.OnStreamExecute = x.OnStreamExecute
	case x.OnStreamExecute == nil:
		ret.OnStreamExecute = t.OnStreamExecute
	default:
		h1 := t.OnStreamExecute
		h2 := x.OnStreamExecute
		ret.OnStreamExecute = func(s ScriptingStreamExecuteStartInfo) func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
					r11 := r1(s)
					r21 := r2(s)
					switch {
					case r11 == nil:
						return r21
					case r21 == nil:
						return r11
					default:
						return func(s ScriptingStreamExecuteDoneInfo) {
							r11(s)
							r21(s)
						}
					}
				}
			}
		}
	}
	switch {
	case t.OnExplain == nil:
		ret.OnExplain = x.OnExplain
	case x.OnExplain == nil:
		ret.OnExplain = t.OnExplain
	default:
		h1 := t.OnExplain
		h2 := x.OnExplain
		ret.OnExplain = func(s ScriptingExplainStartInfo) func(ScriptingExplainDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s ScriptingExplainDoneInfo) {
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
		ret.OnClose = func(s ScriptingCloseStartInfo) func(ScriptingCloseDoneInfo) {
			r1 := h1(s)
			r2 := h2(s)
			switch {
			case r1 == nil:
				return r2
			case r2 == nil:
				return r1
			default:
				return func(s ScriptingCloseDoneInfo) {
					r1(s)
					r2(s)
				}
			}
		}
	}
	return ret
}
func (t Scripting) onExecute(s ScriptingExecuteStartInfo) func(ScriptingExecuteDoneInfo) {
	fn := t.OnExecute
	if fn == nil {
		return func(ScriptingExecuteDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(ScriptingExecuteDoneInfo) {
			return
		}
	}
	return res
}
func (t Scripting) onStreamExecute(s ScriptingStreamExecuteStartInfo) func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
	fn := t.OnStreamExecute
	if fn == nil {
		return func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
			return func(ScriptingStreamExecuteDoneInfo) {
				return
			}
		}
	}
	res := fn(s)
	if res == nil {
		return func(ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
			return func(ScriptingStreamExecuteDoneInfo) {
				return
			}
		}
	}
	return func(s ScriptingStreamExecuteIntermediateInfo) func(ScriptingStreamExecuteDoneInfo) {
		res := res(s)
		if res == nil {
			return func(ScriptingStreamExecuteDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t Scripting) onExplain(s ScriptingExplainStartInfo) func(ScriptingExplainDoneInfo) {
	fn := t.OnExplain
	if fn == nil {
		return func(ScriptingExplainDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(ScriptingExplainDoneInfo) {
			return
		}
	}
	return res
}
func (t Scripting) onClose(s ScriptingCloseStartInfo) func(ScriptingCloseDoneInfo) {
	fn := t.OnClose
	if fn == nil {
		return func(ScriptingCloseDoneInfo) {
			return
		}
	}
	res := fn(s)
	if res == nil {
		return func(ScriptingCloseDoneInfo) {
			return
		}
	}
	return res
}
func ScriptingOnExecute(t Scripting, c *context.Context, query string, parameters scriptingQueryParameters) func(result scriptingResult, _ error) {
	var p ScriptingExecuteStartInfo
	p.Context = c
	p.Query = query
	p.Parameters = parameters
	res := t.onExecute(p)
	return func(result scriptingResult, e error) {
		var p ScriptingExecuteDoneInfo
		p.Result = result
		p.Error = e
		res(p)
	}
}
func ScriptingOnStreamExecute(t Scripting, c *context.Context, query string, parameters scriptingQueryParameters) func(error) func(error) {
	var p ScriptingStreamExecuteStartInfo
	p.Context = c
	p.Query = query
	p.Parameters = parameters
	res := t.onStreamExecute(p)
	return func(e error) func(error) {
		var p ScriptingStreamExecuteIntermediateInfo
		p.Error = e
		res := res(p)
		return func(e error) {
			var p ScriptingStreamExecuteDoneInfo
			p.Error = e
			res(p)
		}
	}
}
func ScriptingOnExplain(t Scripting, c *context.Context, query string) func(plan string, _ error) {
	var p ScriptingExplainStartInfo
	p.Context = c
	p.Query = query
	res := t.onExplain(p)
	return func(plan string, e error) {
		var p ScriptingExplainDoneInfo
		p.Plan = plan
		p.Error = e
		res(p)
	}
}
func ScriptingOnClose(t Scripting, c *context.Context) func(error) {
	var p ScriptingCloseStartInfo
	p.Context = c
	res := t.onClose(p)
	return func(e error) {
		var p ScriptingCloseDoneInfo
		p.Error = e
		res(p)
	}
}
