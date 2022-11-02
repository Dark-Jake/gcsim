package dori

import (
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/gadget"
)

// A1:
// After a character connected to the Jinni triggers an Electro-Charged, Superconduct, Overloaded, Quicken, Aggravate, Hyperbloom,
//
//	or an Electro Swirl or Crystallize reaction, the CD of Spirit-Warding Lamp: Troubleshooter Cannon is decreased by 1s.
//
// This effect can be triggered once every 3s.
func (c *char) a1() {
	const icdKey = "dori-a1"
	icd := 180 // 3s * 60

	reduce := func(args ...interface{}) bool {
		if _, ok := args[0].(*gadget.Gadget); ok {
			return false
		}

		atk := args[1].(*combat.AttackEvent)

		if c.Core.Player.Active() != atk.Info.ActorIndex { // only for on field character
			return false
		}
		if c.StatusIsActive(icdKey) {
			return false
		}
		c.AddStatus(icdKey, icd, true)
		c.ReduceActionCooldown(action.ActionSkill, 60)
		c.Core.Log.NewEvent("dori a1 proc", glog.LogCharacterEvent, c.Index).
			Write("reaction", atk.Info.Abil).
			Write("new cd", c.Cooldown(action.ActionSkill))
		return false
	}

	reduceNoGadget := func(args ...interface{}) bool {
		if _, ok := args[0].(*gadget.Gadget); ok {
			return false
		}

		return reduce(args...)
	}
	c.Core.Events.Subscribe(event.OnOverload, reduceNoGadget, "dori-a1")
	c.Core.Events.Subscribe(event.OnElectroCharged, reduceNoGadget, "dori-a1")
	c.Core.Events.Subscribe(event.OnSuperconduct, reduceNoGadget, "dori-a1")
	c.Core.Events.Subscribe(event.OnQuicken, reduceNoGadget, "dori-a1")
	c.Core.Events.Subscribe(event.OnAggravate, reduceNoGadget, "dori-a1")
	c.Core.Events.Subscribe(event.OnHyperbloom, reduce, "dori-a1")
	c.Core.Events.Subscribe(event.OnCrystallizeElectro, reduceNoGadget, "dori-a1")
	c.Core.Events.Subscribe(event.OnSwirlElectro, reduceNoGadget, "dori-a1")
}
