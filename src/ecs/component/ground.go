package component

type Ground struct {
	Obstacle bool
}

func NewGround(isObstacle bool) Ground {
	return Ground{Obstacle: isObstacle}
}

func (g Ground) ID() ID {
	return GroundID
}
