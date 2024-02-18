package locations

import (
	"fmt"
	"text_game/consts"
)

type kitchen struct {
	locationState
}

func (k *kitchen) GetOnEnterMessage() string {
	return fmt.Sprintf(string(consts.KitchenOnEnterFormat), k.NextPathString())
}

func (k *kitchen) LookAround(step int) string {

	return fmt.Sprintf(string(consts.KitchenLookAroundFormat), k.NextPathString())
}
