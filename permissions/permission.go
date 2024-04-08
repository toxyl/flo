package permissions

import (
	"github.com/toxyl/flo/bitmask"
	"github.com/toxyl/flo/config"
	"github.com/toxyl/flo/utils"
)

type Permission struct {
	mask *bitmask.Bitmask
}

func (t *Permission) Uint() uint32 { return t.mask.Uint() }

func (t *Permission) HasRead() bool  { return t.mask.MatchAny(PERM_READ) }
func (t *Permission) HasWrite() bool { return t.mask.MatchAny(PERM_WRITE) }
func (t *Permission) HasExec() bool  { return t.mask.MatchAny(PERM_EXEC) }

func (t *Permission) Set(r, w, x bool) {
	if r {
		t.SetRead()
	}
	if w {
		t.SetWrite()
	}
	if x {
		t.SetExec()
	}
}

func (t *Permission) SetRead()  { t.mask = t.mask.Set(PERM_READ) }
func (t *Permission) SetWrite() { t.mask = t.mask.Set(PERM_WRITE) }
func (t *Permission) SetExec()  { t.mask = t.mask.Set(PERM_EXEC) }

func (t *Permission) Clear(r, w, x bool) {
	if r {
		t.ClearRead()
	}
	if w {
		t.ClearWrite()
	}
	if x {
		t.ClearExec()
	}
}

func (t *Permission) ClearRead()  { t.mask = t.mask.Clear(PERM_READ) }
func (t *Permission) ClearWrite() { t.mask = t.mask.Clear(PERM_WRITE) }
func (t *Permission) ClearExec()  { t.mask = t.mask.Clear(PERM_EXEC) }

func (t *Permission) String() string {
	return utils.NewString().
		StrAlt(t.HasRead(), config.PermRead, config.PermNone).
		StrAlt(t.HasWrite(), config.PermWrite, config.PermNone).
		StrAlt(t.HasExec(), config.PermExec, config.PermNone).
		String()
}

func (t *Permission) Risk() float64 {
	score := 0.0
	if t.HasExec() {
		score += 4
	}
	if t.HasWrite() {
		score += 2
	}
	if t.HasRead() {
		score += 1
	}
	return score / 7.0
}

func NewPermission(v uint32) *Permission {
	p := &Permission{
		mask: bitmask.New(v),
	}
	return p
}
