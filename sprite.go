/*
 Copyright 2021 The GoPlus Authors (goplus.org)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package spx

import (
	"image/color"
	"math"

	"github.com/goplus/spx/internal/gdi"
	"github.com/goplus/spx/internal/gdi/clrutil"
)

type Sprite struct {
	baseObj
	*Game

	name string

	x, y          float64
	scale         float64
	direction     float64
	rotationStyle RotationStyle
	say           *sayOrThinker

	penColor color.RGBA
	penShade float64
	penHue   float64
	penWidth float64

	visible     bool
	isDraggable bool
	isCloned    bool
	isPenDown   bool
}

type Object interface {
	objMark()
}

type specialObj int

func (p specialObj) objMark() {}
func (p *Sprite) objMark()    {}

var (
	Mouse  = specialObj(1)
	Edge   = specialObj(2)
	Random = specialObj(3)
)

const (
	Right = 90
	Left  = -90
	Up    = 0
	Down  = 180
)

// -----------------------------------------------------------------------------

func (p *Sprite) Clone(data interface{}) *Sprite {
	panic("todo")
}

func (p *Sprite) Destroy() { // delete this clone
	panic("todo")
}

// -----------------------------------------------------------------------------

func (p *Sprite) Hide() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.visible = false
}

func (p *Sprite) Show() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.visible = true
}

// -----------------------------------------------------------------------------

func (p *Sprite) CostumeName() string {
	return p.costumeName()
}

func (p *Sprite) CostumeIndex() int {
	return p.costumeIndex()
}

func (p *Sprite) SetCostume(costume interface{}) {
	p.setCostume(costume)
}

func (p *Sprite) NextCostume() {
	p.nextCostume()
}

// -----------------------------------------------------------------------------

func (p *Sprite) Say(msg string, secs ...float64) {
	p.sayOrThink(msg, styleSay)
	if secs != nil {
		p.waitStopSay(secs[0])
	}
}

func (p *Sprite) Think(msg string, secs ...float64) {
	p.sayOrThink(msg, styleThink)
	if secs != nil {
		p.waitStopSay(secs[0])
	}
}

func (p *Sprite) sayOrThink(msg string, style int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if msg == "" {
		p.doStopSay()
		return
	}

	old := p.say
	if old == nil {
		p.say = &sayOrThinker{sp: p, msg: msg, style: style}
		p.addShape(p.say)
	} else {
		old.msg, old.style = msg, style
		p.activateShape(old)
	}
}

func (p *Sprite) waitStopSay(secs float64) {
	p.Wait(secs)

	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.doStopSay()
}

func (p *Sprite) doStopSay() {
	if p.say != nil {
		p.removeShape(p.say)
		p.say = nil
	}
}

// -----------------------------------------------------------------------------

func (p *Sprite) getXY() (x, y float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.x, p.y
}

// DistanceTo func:
//   DistanceTo(sprite)
//   DistanceTo(gox.Mouse)
//   DistanceTo(gox.Edge)
func (p *Sprite) DistanceTo(obj Object) float64 {
	p.mutex.Lock()
	x, y := p.x, p.y
	p.mutex.Unlock()

	_, _ = x, y
	panic("todo")
	// return p.distanceTo(x, y, name)
}

func (p *Sprite) doMoveTo(x, y float64) {
	if p.isPenDown {
		p.Game.movePen(p, x, y)
	}
	p.x, p.y = x, y
}

func (p *Sprite) Step(step float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	sin, cos := math.Sincos(toRadian(p.direction))
	p.doMoveTo(p.x+step*sin, p.y+step*cos)
}

// Goto func:
//   Goto(sprite)
//   Goto(gox.Mouse)
func (p *Sprite) Goto(obj Object) {
	panic("todo")
	// x, y := p.g.mouseOrSpritePos(where)
	// p.setXY(x, y)
}

const (
	glideTick = 1e8
)

func (p *Sprite) Glide(x, y float64, secs float64) {
	inDur := int64(secs * 1e9)
	n := int(inDur / glideTick)
	if n > 0 {
		x0, y0 := p.getXY()
		dx := (x - x0) / float64(n)
		dy := (y - y0) / float64(n)
		for i := 1; i < n; i++ {
			p.sleep(glideTick)
			inDur -= glideTick
			x0 += dx
			y0 += dy
			p.SetXY(x0, y0)
		}
	}
	p.sleep(inDur)
	p.SetXY(x, y)
}

func (p *Sprite) SetXY(x, y float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.doMoveTo(x, y)
}

func (p *Sprite) Xpos() float64 {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.x
}

func (p *Sprite) SetXpos(x float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.doMoveTo(x, p.y)
}

func (p *Sprite) ChangeXpos(dx float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.doMoveTo(p.x+dx, p.y)
}

func (p *Sprite) Ypos() float64 {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.y
}

func (p *Sprite) SetYpos(y float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.doMoveTo(p.x, y)
}

func (p *Sprite) ChangeYpos(dy float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.doMoveTo(p.x, p.y+dy)
}

// -----------------------------------------------------------------------------

type RotationStyle int

func (p *Sprite) SetRotationStyle(style RotationStyle) {
	panic("todo")
}

func (p *Sprite) Heading() float64 {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.direction
}

// Turn func:
//   Turn(degree)
//   Turn(gox.Left)
//   Turn(gox.Right)
func (p *Sprite) Turn(delta float64) {
	p.setDirection(delta, true)
}

// TurnTo func:
//   TurnTo(sprite)
//   TurnTo(spriteName)
//   TurnTo(degree)
//   TurnTo(gox.Mouse)
//   TurnTo(gox.Left)
//   TurnTo(gox.Right)
//   TurnTo(gox.Up)
//   TurnTo(gox.Down)
func (p *Sprite) TurnTo(obj interface{}) {
	switch v := obj.(type) {
	case int:
		p.setDirection(float64(v), false)
	case float64:
		p.setDirection(v, false)
	default:
		x, y := p.objectPos(obj)
		dx := x - p.x
		dy := y - p.y
		angle := 90 - math.Atan2(dy, dx)*180/math.Pi
		p.setDirection(angle, false)
	}
}

func (p *Sprite) setDirection(dir float64, change bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if change {
		dir += p.direction
	}
	p.direction = normalizeDirection(dir)
}

// -----------------------------------------------------------------------------

func (p *Sprite) Size() float64 {
	p.mutex.Lock()
	v := p.scale
	p.mutex.Unlock()

	return v
}

func (p *Sprite) SetSize(size float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.scale = size
}

func (p *Sprite) ChangeSize(delta float64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.scale += delta
}

// -----------------------------------------------------------------------------

type Color = color.RGBA

func (p *Sprite) TouchingColor(color Color) bool {
	panic("todo")
}

// Touching func:
//   Touching(spriteName)
//   Touching(sprite)
//   Touching(spx.Mouse)
//   Touching(spx.Edge)
func (p *Sprite) Touching(obj interface{}) bool {
	switch v := obj.(type) {
	case string:
		return p.touchingSpriteBy(p, v)
	case *Sprite:
		if touchingSprite(p, v) {
			return true
		}
	case specialObj:
		if v == Edge {
			if p.checkTouchingScreen(touchingAllEdges) != 0 {
				return true
			}
		} else if v == Mouse {
			x, y := p.getMousePos()
			return p.touchingPoint(p, float64(x), float64(y))
		}
	default:
		panic("Touching: unexpected input")
	}
	return false
}

func touchingSprite(dst, src *Sprite) bool {
	sp1, pt1 := dst.getGdiSprite()
	sp2, pt2 := src.getGdiSprite()
	return gdi.Touching(sp1, pt1, sp2, pt2)
}

const (
	touchingScreenLeft   = 1
	touchingScreenTop    = 2
	touchingScreenRight  = 4
	touchingScreenBottom = 8
	touchingAllEdges     = 15
)

func (p *Sprite) BounceOffEdge() {
	dir := p.Heading()
	where := checkTouchingDirection(dir)
	touching := p.checkTouchingScreen(where)
	if touching == 0 {
		return
	}
	if (touching & (touchingScreenLeft | touchingScreenRight)) != 0 {
		dir = -dir
	} else {
		dir = 180 - dir
	}

	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.direction = normalizeDirection(dir)
}

func checkTouchingDirection(dir float64) int {
	if dir > 0 {
		if dir < 90 {
			return touchingScreenRight | touchingScreenTop
		}
		if dir > 90 {
			if dir == 180 {
				return touchingScreenBottom
			}
			return touchingScreenRight | touchingScreenBottom
		}
		return touchingScreenRight
	}
	if dir < 0 {
		if dir > -90 {
			return touchingScreenLeft | touchingScreenTop
		}
		if dir < -90 {
			return touchingScreenLeft | touchingScreenBottom
		}
		return touchingScreenLeft
	}
	return touchingScreenTop
}

func (p *Sprite) checkTouchingScreen(where int) (touching int) {
	spr, pt := p.getGdiSprite()
	if spr == nil {
		return
	}

	if (where & touchingScreenLeft) != 0 {
		if gdi.TouchingRect(spr, pt, -1e8, -1e8, 0, 1e8) {
			return touchingScreenLeft
		}
	}
	if (where & touchingScreenTop) != 0 {
		if gdi.TouchingRect(spr, pt, -1e8, -1e8, 1e8, 0) {
			return touchingScreenTop
		}
	}
	w, h := p.size()
	if (where & touchingScreenRight) != 0 {
		if gdi.TouchingRect(spr, pt, w, -1e8, 1e8, 1e8) {
			return touchingScreenRight
		}
	}
	if (where & touchingScreenBottom) != 0 {
		if gdi.TouchingRect(spr, pt, -1e8, h, 1e8, 1e8) {
			return touchingScreenBottom
		}
	}
	return
}

// -----------------------------------------------------------------------------

func (p *Sprite) GoBackLayers(n int) {
	p.goBackByLayers(p, n)
}

func (p *Sprite) GotoFront(n int) {
	p.goBackByLayers(p, -1e8)
}

// -----------------------------------------------------------------------------

func (p *Sprite) Stamp() {
	p.stampCostume(p.getDrawInfo())
}

func (p *Sprite) PenUp() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.isPenDown = false
}

func (p *Sprite) PenDown() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.isPenDown = true
}

func (p *Sprite) SetPenColor(color Color) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	h, _, v := clrutil.RGB2HSV(color.R, color.G, color.B)
	p.penHue = (200 * h) / 360
	p.penShade = 50 * v
	p.penColor = color
}

func (p *Sprite) ChangePenColor(delta float64) {
	panic("todo")
}

func (p *Sprite) SetPenShade(shade float64) {
	p.setPenShade(shade, false)
}

func (p *Sprite) ChangePenShade(delta float64) {
	p.setPenShade(delta, true)
}

func (p *Sprite) SetPenHue(hue float64) {
	p.setPenHue(hue, false)
}

func (p *Sprite) ChangePenHue(delta float64) {
	p.setPenHue(delta, true)
}

func (p *Sprite) setPenHue(v float64, change bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if change {
		v += p.penHue
	}
	v = math.Mod(v, 200)
	if v < 0 {
		v += 200
	}
	p.penHue = v
	p.doUpdatePenColor()
}

func (p *Sprite) setPenShade(v float64, change bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if change {
		v += p.penShade
	}
	v = math.Mod(v, 200)
	if v < 0 {
		v += 200
	}
	p.penShade = v
	p.doUpdatePenColor()
}

func (p *Sprite) doUpdatePenColor() {
	r, g, b := clrutil.HSV2RGB((p.penHue*180)/100, 1, 1)
	shade := p.penShade
	if shade > 100 { // range 0..100
		shade = 200 - shade
	}
	if shade < 50 {
		r, g, b = clrutil.MixRGB(0, 0, 0, r, g, b, (10+shade)/60)
	} else {
		r, g, b = clrutil.MixRGB(r, g, b, 255, 255, 255, (shade-50)/60)
	}
	p.penColor = color.RGBA{R: r, G: g, B: b, A: p.penColor.A}
}

func (p *Sprite) SetPenSize(size float64) {
	p.setPenWidth(size, true)
}

func (p *Sprite) ChangePenSize(delta float64) {
	p.setPenWidth(delta, true)
}

func (p *Sprite) setPenWidth(w float64, change bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if change {
		w += p.penWidth
	}
	p.penWidth = w
}

// -----------------------------------------------------------------------------
