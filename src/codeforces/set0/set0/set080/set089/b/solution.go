package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, s := range res {
		buf.WriteString(s)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func drive(reader *bufio.Reader) []string {
	n := readNum(reader)
	instructions := make([]string, n)
	for i := 0; i < n; i++ {
		instructions[i] = readString(reader)
	}
	return solve(instructions)
}

func solve(instructions []string) []string {
	parent := make(map[string]string)
	widgets := make(map[string]HasDimension)

	doPack := func(s string) {
		// a.pack(b)
		i := strings.Index(s, ".pack")
		name1 := s[:i]
		box := widgets[name1]
		name2 := s[i+len(".pack(") : len(s)-1]
		if hbox, ok := box.(*HBox); ok {
			hbox.pack(widgets[name2])
		} else if vbox, ok := box.(*VBox); ok {
			vbox.pack(widgets[name2])
		}

		parent[name2] = name1
	}

	doSetBorder := func(s string) {
		i := strings.Index(s, ".set_border")
		name := s[:i]
		b := s[i+len(".set_border")+1 : len(s)-1]
		border, _ := strconv.Atoi(b)

		box := widgets[name]
		if hbox, ok := box.(*HBox); ok {
			hbox.setBorder(border)
		} else if vbox, ok := box.(*VBox); ok {
			vbox.setBorder(border)
		}
	}

	doSetSpacing := func(s string) {
		i := strings.Index(s, ".set_spacing")
		name := s[:i]
		spacing, _ := strconv.Atoi(s[i+len(".set_spacing")+1 : len(s)-1])

		box := widgets[name]
		if hbox, ok := box.(*HBox); ok {
			hbox.setSpacing(spacing)
		} else if vbox, ok := box.(*VBox); ok {
			vbox.setSpacing(spacing)
		}
	}

	for _, cur := range instructions {
		if strings.HasPrefix(cur, "Widget ") {
			w := createWidget(cur[len("Widget "):])
			widgets[w.name] = w
		} else if strings.HasPrefix(cur, "HBox ") {
			b := createHBox(cur[len("HBox "):])
			widgets[b.name] = b
		} else if strings.HasPrefix(cur, "VBox ") {
			b := createVBox(cur[len("VBox "):])
			widgets[b.name] = b
		} else if strings.Contains(cur, ".pack(") {
			doPack(cur)
		} else if strings.Contains(cur, ".set_border(") {
			doSetBorder(cur)
		} else if strings.Contains(cur, ".set_spacing(") {
			doSetSpacing(cur)
		}
	}

	for name, w := range widgets {
		if _, ok := parent[name]; !ok {
			// 只去处理root节点
			if w, isHbox := w.(*HBox); isHbox {
				w.invalidate()
			}
			if w, isVbox := w.(*VBox); isVbox {
				w.invalidate()
			}
		}
	}

	var ans []string

	for name, w := range widgets {
		ans = append(ans, fmt.Sprintf("%s %d %d", name, w.getWidth(), w.getHeight()))
	}

	slices.Sort(ans)

	return ans
}

type HasDimension interface {
	getWidth() int
	getHeight() int
}

type Widget struct {
	width  int
	height int
	name   string
}

func createWidget(s string) *Widget {
	i := strings.Index(s, "(")
	name := s[:i]
	s = s[i+1 : len(s)-1]
	parts := strings.Split(s, ",")
	width, _ := strconv.Atoi(parts[0])
	height, _ := strconv.Atoi(parts[1])
	return &Widget{
		width:  width,
		height: height,
		name:   name,
	}
}

func (widget *Widget) getWidth() int {
	return widget.width
}

func (widget *Widget) getHeight() int {
	return widget.height
}

type Box struct {
	Widget
	flag     bool
	border   int
	spacing  int
	children []HasDimension
}

type HBox struct {
	Box
}

type VBox struct {
	Box
}

func createHBox(name string) *HBox {
	return &HBox{
		Box: Box{
			Widget: Widget{
				name: name,
			},
		},
	}
}

func createVBox(name string) *VBox {
	return &VBox{
		Box: Box{
			Widget: Widget{
				name: name,
			},
		},
	}
}

func (box *Box) setBorder(border int) {
	box.border = border
}

func (box *Box) setSpacing(spacing int) {
	box.spacing = spacing
}

func (box *Box) pack(widget HasDimension) {
	box.children = append(box.children, widget)
}

func (box *HBox) invalidate() {
	if len(box.children) == 0 || box.flag {
		return
	}
	box.flag = true
	box.width = 0
	box.height = 0
	for _, child := range box.children {
		if hbox, ok := child.(*HBox); ok {
			hbox.invalidate()
		} else if vbox, ok := child.(*VBox); ok {
			vbox.invalidate()
		}
		box.width += child.getWidth()
		box.height = max(box.height, child.getHeight())
	}

	box.width += (len(box.children) - 1) * box.spacing
	box.height += 2 * box.border
	box.width += 2 * box.border
}

func (box *VBox) invalidate() {
	if len(box.children) == 0 || box.flag {
		return
	}
	box.flag = true
	box.width = 0
	box.height = 0
	for _, child := range box.children {
		if hbox, ok := child.(*HBox); ok {
			hbox.invalidate()
		} else if vbox, ok := child.(*VBox); ok {
			vbox.invalidate()
		}
		box.height += child.getHeight()
		box.width = max(box.width, child.getWidth())
	}

	box.height += (len(box.children) - 1) * box.spacing
	box.width += 2 * box.border
	box.height += 2 * box.border
}
