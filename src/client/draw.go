package client

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/entity"
	"github.com/niakr1s/nrg-go/src/ecs/tag"
	"github.com/niakr1s/nrg-go/src/img"
	"github.com/sirupsen/logrus"
)

func (c *Client) produceBoard() *ebiten.Image {
	board, _ := ebiten.NewImage(config.BoardWidth, config.BoardHeight, ebiten.FilterDefault)
	board.Fill(color.Gray16{0xaaaf})
	c.Reg.RLock()
	for _, e := range c.Reg.Entities {
		e.RLock()
		if cs := e.GetComponents(component.ShapeID, component.PosID); cs != nil {
			shape := cs[0].(component.Shape)
			pos := cs[1].(component.Pos)
			drawShape(board, e, pos, shape)
			drawWeapon(board, e, pos, shape)
			drawHp(board, e, pos, shape)
		}
		e.RUnlock()
	}
	c.Reg.RUnlock()
	return board
}

func drawHp(board *ebiten.Image, e *entity.Entity, pos component.Pos, shape component.Shape) {
	hpC := e.GetComponents(component.HpID)
	if hpC == nil {
		return
	}
	hp := hpC[0].(component.HP)
	center := shape.OuterPointInDirectionDiff(component.TopVec).Sum(pos).Move(component.TopVec, 10)
	w := shape.Bound(pos).Width()
	drawRectWithCenter(board, center, w, 3.0, color.White)
	drawRectWithCenter(board, center, w*hp.Percent(), 3.0, color.RGBA{0, 200, 0, 255})
}

func drawRectWithCenter(dst *ebiten.Image, center component.Pos, w, h float64, clr color.Color) {
	topLeft := center.MoveDist(component.TopVec, h/2).MoveDist(component.LeftVec, w/2)
	ebitenutil.DrawRect(dst, topLeft.X, topLeft.Y, w, h, clr)
}

func drawShape(board *ebiten.Image, e *entity.Entity, pos component.Pos, shape component.Shape) {
	switch shape := shape.(type) {
	case component.Circle:
		image, err := img.Load(getCirclePath(e))
		if err != nil {
			panic(err)
		}
		drawCircle(board, image, pos, shape)
	default:
		logrus.Warningf("Client.produceBoard(): couldn't draw unknowh shape")
	}
}

func drawWeapon(board *ebiten.Image, e *entity.Entity, pos component.Pos, shape component.Shape) {
	weapC := e.GetComponents(component.WeaponID)
	if weapC == nil {
		return
	}
	weap := weapC[0].(component.Weapon)
	dirs := weap.GetGunDirs()
	for _, dir := range dirs {
		endPos := shape.OuterPointInDirectionDiff(dir).Sum(pos)
		ebitenutil.DrawLine(board, pos.X, pos.Y, endPos.X, endPos.Y, color.Black)
	}
}

func getCirclePath(e *entity.Entity) string {
	if e.HasTags(tag.User) {
		return img.BlueCircle
	} else if e.HasTags(tag.Enemy) {
		return img.WhiteCircle
	} else if e.HasTags(tag.Bullet) {
		return img.RedCircle
	}
	return img.RedCircle
}

func drawCircle(board, image *ebiten.Image, pos component.Pos, circle component.Circle) {
	op := &ebiten.DrawImageOptions{}
	w, h := image.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	scale := float64(circle.R*2) / float64(w)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(pos.X, pos.Y)

	board.DrawImage(image, op)
}

func (c *Client) drawBoard(screen *ebiten.Image) {
	board := c.produceBoard()
	op := &ebiten.DrawImageOptions{}

	swI, shI := screen.Size()
	bwI, bhI := board.Size()

	// sizes of screen and board
	sw, sh := float64(swI), float64(shI)
	bw, bh := float64(bwI), float64(bhI)

	scale := 0.9 * math.Min(sw, sh) / math.Max(bw, bh)

	// scaled board size
	bw, bh = bw*scale, bh*scale

	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate((sw-bw)/2, (sh-bh)/2)

	screen.DrawImage(board, op)
}
