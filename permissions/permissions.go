package permissions

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/toxyl/flo/bitmask"
	"github.com/toxyl/flo/config"
	"github.com/toxyl/flo/utils"
)

type Permissions struct {
	raw        uint32
	rawType    *bitmask.Bitmask
	rawMode    *bitmask.Bitmask
	rawUnknown *bitmask.Bitmask
	rawPerm    *bitmask.Bitmask
	rawOwner   *bitmask.Bitmask
	rawGroup   *bitmask.Bitmask
	rawWorld   *bitmask.Bitmask
	mode       struct {
		link   bool
		sticky bool
		uid    bool
		gid    bool
	}
	char        bool
	dir         bool
	blockDevice bool
	charDevice  bool
	fifo        bool
	socket      bool
	owner       *Permission
	group       *Permission
	world       *Permission
}

func (p *Permissions) HasNoSize() bool {
	return p.dir || p.fifo || p.blockDevice || p.charDevice || p.socket
}

func (p *Permissions) HasSize() bool {
	return !p.HasNoSize()
}

func (p *Permissions) Type() string {
	if p.dir {
		return config.TypeDir
	}
	if p.fifo {
		return config.TypeFIFO
	}
	if p.socket {
		return config.TypeSocket
	}
	if p.charDevice {
		return config.TypeCharDevice
	}
	if p.blockDevice {
		return config.TypeBlockDevice
	}
	return config.TypeFile
}

func (p *Permissions) Mode() string {
	return utils.NewString().
		StrAlt(p.mode.sticky, config.ModeSticky, config.ModeNone).
		StrAlt(p.mode.uid, config.ModeUID, config.ModeNone).
		StrAlt(p.mode.gid, config.ModeGID, config.ModeNone).
		String()
}

func (p *Permissions) RiskString() string {
	r := p.Risk()
	return utils.NewString().
		StrAlt(r >= 0.25, config.RiskLow, config.RiskNone).
		StrAlt(r >= 0.5, config.RiskMedium, config.RiskNone).
		StrAlt(r >= 0.8, config.RiskHigh, config.RiskNone).
		String()
}

func (p *Permissions) String() string {
	return utils.NewString().
		StrAlt(p.mode.link, config.ModeLink, config.ModeNoLink).Pad(1).
		Str(p.Type()).Str(p.Mode()).Pad(1).
		Str(p.owner.String()).Pad(1).
		Str(p.group.String()).Pad(1).
		Str(p.world.String()).Pad(1).
		Str(p.Octal()).Pad(1).
		Str(p.RiskString()).
		String()
}

func (p *Permissions) FileMode() fs.FileMode {
	return fs.FileMode(p.Uint())
}

func (p *Permissions) Uint() uint32 {
	return p.rawType.
		ShiftLeft(MASK_TYPE_SHIFT).
		Set(p.rawMode.ShiftLeft(MASK_MODE_SHIFT).Uint()).
		Set(p.rawUnknown.ShiftLeft(MASK_UNKNOWN_SHIFT).Uint()).
		Set(p.owner.mask.ShiftLeft(MASK_PERM_OWNER_SHIFT).Uint()).
		Set(p.group.mask.ShiftLeft(MASK_PERM_GROUP_SHIFT).Uint()).
		Set(p.world.mask.ShiftLeft(MASK_PERM_WORLD_SHIFT).Uint()).
		Uint()
}

func (p *Permissions) Octal() string {
	pn := p.owner.mask.ShiftLeft(MASK_PERM_OWNER_SHIFT).
		Set(p.group.mask.ShiftLeft(MASK_PERM_GROUP_SHIFT).Uint()).
		Set(p.world.mask.ShiftLeft(MASK_PERM_WORLD_SHIFT).Uint())
	return fmt.Sprintf("%04o", pn.Uint())
}

func (p *Permissions) IsDir() bool         { return p.dir }
func (p *Permissions) IsLink() bool        { return p.mode.link }
func (p *Permissions) IsBlockDevice() bool { return p.blockDevice }
func (p *Permissions) IsCharDevice() bool  { return p.charDevice }
func (p *Permissions) IsFIFO() bool        { return p.fifo }
func (p *Permissions) IsSocket() bool      { return p.socket }
func (p *Permissions) IsSticky() bool      { return p.mode.sticky }
func (p *Permissions) IsSetUID() bool      { return p.mode.uid }
func (p *Permissions) IsSetGID() bool      { return p.mode.gid }

func (p *Permissions) Owner() *Permission { return p.owner }
func (p *Permissions) Group() *Permission { return p.group }
func (p *Permissions) World() *Permission { return p.world }

// Risk will calculate the overall risk level of the permissions where world-access has the most weight,
// followed by group access and finally owner access.
//
// For example:
// a file with full access for the owner but no-one else will yield a risk level of 0.1428571428571429.
// Changing full access from owner to group would increase this to 0.2857142857142857 and
// changing to world would lead to 0.5714285714285714.
func (p *Permissions) Risk() float64 {
	w := p.world.Risk() * 4
	g := p.group.Risk() * 2
	o := p.owner.Risk()
	return (w + g + o) / 7.0
}

func (p *Permissions) Set(perm uint32) *Permissions {
	p.raw = perm

	// there is a section I don't know the purpose of, let's preserve it
	// in case it contains relevant data on systems I haven't tested
	p.rawUnknown = bitmask.NewWithMask(perm, MASK_UNKNOWN, MASK_UNKNOWN_SHIFT)

	// mode
	p.rawMode = bitmask.NewWithMask(perm, MASK_MODE, MASK_MODE_SHIFT)

	p.mode.link = p.rawMode.MatchAny(MODE_LINK)
	p.mode.sticky = p.rawMode.MatchAny(MODE_STICKY)
	p.mode.uid = p.rawMode.MatchAny(MODE_UID)
	p.mode.gid = p.rawMode.MatchAny(MODE_GID)

	// check for known types
	p.rawType = bitmask.NewWithMask(perm, MASK_TYPE, MASK_TYPE_SHIFT)

	p.char = p.rawType.Match(TYPE_CHAR)
	p.dir = p.rawType.Match(TYPE_DIR)
	p.blockDevice = p.rawType.Match(TYPE_BLOCK_DEVICE)
	p.charDevice = p.rawType.Match(TYPE_CHAR_DEVICE)
	p.fifo = p.rawType.Match(TYPE_FIFO)
	p.socket = p.rawType.Match(TYPE_SOCKET)

	// get the permissions
	p.rawPerm = bitmask.NewWithMask(perm, MASK_PERM, MASK_PERM_SHIFT)
	p.rawOwner = bitmask.NewWithMask(perm, MASK_PERM_OWNER, MASK_PERM_OWNER_SHIFT)
	p.rawGroup = bitmask.NewWithMask(perm, MASK_PERM_GROUP, MASK_PERM_GROUP_SHIFT)
	p.rawWorld = bitmask.NewWithMask(perm, MASK_PERM_WORLD, MASK_PERM_WORLD_SHIFT)

	p.owner = NewPermission(p.rawOwner.Uint())
	p.group = NewPermission(p.rawGroup.Uint())
	p.world = NewPermission(p.rawWorld.Uint())

	return p
}

func New(path string) *Permissions {
	p := &Permissions{
		raw:        0,
		rawType:    bitmask.New(0),
		rawMode:    bitmask.New(0),
		rawUnknown: bitmask.New(0),
		rawPerm:    bitmask.New(0),
		rawOwner:   bitmask.New(0),
		rawGroup:   bitmask.New(0),
		rawWorld:   bitmask.New(0),
		mode: struct {
			link   bool
			sticky bool
			uid    bool
			gid    bool
		}{
			link:   false,
			sticky: false,
			uid:    false,
			gid:    false,
		},
		char:        false,
		dir:         false,
		blockDevice: false,
		charDevice:  false,
		fifo:        false,
		socket:      false,
		owner:       NewPermission(0),
		group:       NewPermission(0),
		world:       NewPermission(0),
	}
	path = filepath.Clean(path)
	p.Set(uint32(utils.GetFileModeL(path)))

	// we might have a link, for those we'd like the permissions of the target instead
	if p.IsLink() {
		p.Set(uint32(utils.GetFileMode(path)))
		p.mode.link = true
	}
	return p
}
