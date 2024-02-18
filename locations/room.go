package locations

import (
	"fmt"
	"text_game/consts"
)

type room struct {
	locationState
}

func (r *room) LookAround(step int) string {

	if len(r.GetItems()) > 0 {
		return fmt.Sprintf(string(consts.RoomLookAroundFormat), r.ItemsString(), r.NextPathString())
	} else {
		return fmt.Sprintf(string(consts.RoomEmptyLookAroundFormat), r.NextPathString())
	}
}

func (r *room) GetOnEnterMessage() string {
	return fmt.Sprintf(string(consts.RoomOnEnterFormat), r.NextPathString())
}
