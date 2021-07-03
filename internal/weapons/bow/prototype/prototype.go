package generic

import (
	"fmt"

	"github.com/genshinsim/gsim/pkg/combat"
	"github.com/genshinsim/gsim/pkg/def"
)

func init() {
	combat.RegisterWeaponFunc("prototype crescent", weapon)
}

func weapon(c def.Character, s def.Sim, log def.Logger, r int, param map[string]int) {

	dur := 0
	key := fmt.Sprintf("prototype-crescent-%v", c.Name())
	//add on hit effect
	s.AddOnAttackWillLand(func(t def.Target, ds *def.Snapshot) {
		if ds.ActorIndex != c.CharIndex() {
			return
		}
		if ds.HitWeakPoint {
			dur = s.Frame() + 600
		}
	}, key)

	m := make([]float64, def.EndStatType)
	m[def.ATKP] = 0.27 + float64(r)*0.09
	c.AddMod(def.CharStatMod{
		Key: "prototype-crescent",
		Amount: func(a def.AttackTag) ([]float64, bool) {
			if dur < s.Frame() {
				return nil, false
			}
			return m, true
		},
		Expiry: -1,
	})
}