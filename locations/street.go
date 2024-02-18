package locations

import (
	"fmt"
	"text_game/consts"
)

type street struct {
	locationState
}

func (s street) GetOnEnterMessage() string {
	return fmt.Sprint(consts.StreetOnEnterFormat)
}
