package dagjose

import (
	ipld "github.com/ipld/go-ipld-prime"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
)

type jweRecipientListAssembler struct{ d *DagJOSE }

func (l *jweRecipientListAssembler) AssembleValue() ipld.NodeAssembler {
	l.d.recipients = append(l.d.recipients, jweRecipient{})
	nextRef := &l.d.recipients[len(l.d.recipients)-1]
	return &jweRecipientAssembler{
		recipient: nextRef,
		key:       nil,
		state:     maState_initial,
	}
}

func (l *jweRecipientListAssembler) Finish() error {
	return nil
}
func (l *jweRecipientListAssembler) ValuePrototype(idx int) ipld.NodePrototype {
	return basicnode.Prototype.Map
}
