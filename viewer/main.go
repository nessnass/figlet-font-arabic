// nessnass-arabic: a terminal viewer for the Nessnass Arabic figlet font.
//
// The rendered font is the edit surface: type Arabic and it re-renders live,
// right to left. It shapes the text itself (shape.go) and reads
// Nessnass-Arabic.flf with its own renderer (flf.go), so figlet isn't needed.
//
// Usage, from this folder:
//
//	go run .                    start with نسناس
//	go run . "دبي"              start with your own text
//	go run . -print "دبي"       print the text in the font and exit
//	go run . -shape "دبي"       print the shaped text, for figlet
//
// -print and -shape read standard input when no text is given.
//
// Keys:
//
//	type / backspace   edit the text      enter          new line
//	tab / shift+tab    colour             ctrl+a         show the alphabet
//	ctrl+l             clear              ctrl+w         save as .txt
//	esc / ctrl+c       quit
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	fontName       = "Nessnass-Arabic"
	startWord      = "نسناس"
	colorAccent    = "#5599cc"
	headerOverhead = 5 // blank, title, info, statusLine (even when empty), help
)

// ============================================================================
// font + colours
// ============================================================================

// findFontFile locates Nessnass-Arabic.flf, checking source-, executable- and
// cwd-relative locations so it works under `go run .`, a built binary, or from
// the parent folder.
func findFontFile() string {
	var dirs []string
	if _, file, _, ok := runtime.Caller(0); ok {
		dirs = append(dirs, filepath.Dir(file), filepath.Dir(filepath.Dir(file)))
	}
	if exe, err := os.Executable(); err == nil {
		d := filepath.Dir(exe)
		dirs = append(dirs, d, filepath.Dir(d))
	}
	dirs = append(dirs, ".", "..")
	for _, d := range dirs {
		p := filepath.Join(d, fontName+".flf")
		if _, err := os.Stat(p); err == nil {
			abs, _ := filepath.Abs(p)
			return abs
		}
	}
	return ""
}

// renderRTL shapes one line and draws it right to left.
func (ff *figFont) renderRTL(line string) string {
	return ff.render(reverse(shape(line)))
}

// renderText draws every line of text, a blank line for an empty one.
func (ff *figFont) renderText(text string) string {
	var blocks []string
	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			blocks = append(blocks, "")
			continue
		}
		blocks = append(blocks, ff.renderRTL(line))
	}
	return strings.Join(blocks, "\n")
}

type colorMode struct {
	name   string
	repHex string // representative hex for UI labels
	apply  func(string) string
}

func hexToRGB(hex string) (int, int, int) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0
	}
	r, _ := strconv.ParseInt(hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(hex[4:6], 16, 0)
	return int(r), int(g), int(b)
}

func solid(name, hex string) colorMode {
	return colorMode{name, hex, func(s string) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(hex)).Bold(true).Render(s)
	}}
}

// The main colours of the first eight Nessnass artworks, in token order
// (https://www.nessnass.com)
func palettes() []colorMode {
	return []colorMode{
		solid("sky blue", "#87ceeb"),
		solid("pink", "#ffc0cb"),
		solid("green", "#00ff00"),
		solid("indigo", "#4b0082"),
		solid("light green", "#90ee90"),
		solid("yellow", "#ffff00"),
		solid("gold", "#ffd700"),
		solid("red", "#ff0000"),
	}
}

// ============================================================================
// model
// ============================================================================

type model struct {
	colors   []colorMode
	colorIdx int
	word     string
	font     *figFont
	preview  string
	status   string
	width    int
	height   int
	quitting bool
}

func (m model) Init() tea.Cmd { return nil }

func (m *model) refresh() {
	if m.word == "" {
		m.preview = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).
			Render("(type Arabic to render…)")
		return
	}
	m.preview = m.colors[m.colorIdx].apply(m.font.renderText(m.word))
}

// export writes the current rendering to a plain .txt file (no colour codes).
func (m *model) export() {
	if m.word == "" {
		return
	}
	name := "art_" + fontName + ".txt"
	if err := os.WriteFile(name, []byte(m.font.renderText(m.word)+"\n"), 0644); err != nil {
		m.status = "export failed: " + err.Error()
		return
	}
	abs, _ := filepath.Abs(name)
	m.status = "saved → " + abs
}

// alphabetString lays out every letter alone, a space apart, as many to a
// line as fit the width. Lines read right to left like the text.
func alphabetString(width int) string {
	perRow := (width - 4) / 7
	if perRow < 3 {
		perRow = 3
	}
	var lines []string
	for i := 0; i < len(alphabet); i += perRow {
		end := min(i+perRow, len(alphabet))
		lines = append(lines, strings.Join(alphabet[i:end], " "))
	}
	return strings.Join(lines, "\n")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.refresh()
		return m, nil

	case tea.KeyMsg:
		m.status = ""

		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyTab:
			m.colorIdx = (m.colorIdx + 1) % len(m.colors)
		case tea.KeyShiftTab:
			m.colorIdx = (m.colorIdx - 1 + len(m.colors)) % len(m.colors)

		case tea.KeyCtrlW:
			m.export()
			return m, nil
		case tea.KeyCtrlA:
			m.word = alphabetString(m.width)
		case tea.KeyCtrlL:
			m.word = ""

		case tea.KeyBackspace:
			if r := []rune(m.word); len(r) > 0 {
				m.word = string(r[:len(r)-1])
			}
		case tea.KeySpace:
			m.word += " "
		case tea.KeyEnter:
			m.word += "\n"
		case tea.KeyRunes:
			m.word += string(msg.Runes)
		default:
			return m, nil
		}
		m.refresh()
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	title := lipgloss.NewStyle().
		Background(lipgloss.Color(colorAccent)).Foreground(lipgloss.Color("#ffffff")).
		Bold(true).Padding(0, 1).
		Render("Nessnass Arabic — figlet font viewer")

	info := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render("  colour: ") +
		lipgloss.NewStyle().Foreground(lipgloss.Color(m.colors[m.colorIdx].repHex)).Render(m.colors[m.colorIdx].name)

	help := lipgloss.NewStyle().Foreground(lipgloss.Color("#777777")).Render(
		"  type · backspace · enter br · tab/shift+tab colour · ctrl+a alphabet · ctrl+l clear · ctrl+w export · esc quit")

	statusLine := ""
	if m.status != "" {
		statusLine = lipgloss.NewStyle().Foreground(lipgloss.Color("#4caf50")).Bold(true).
			Render("  " + m.status)
	}

	previewRows := m.height - headerOverhead
	if previewRows < 3 {
		previewRows = 3
	}
	preview := padMin(indent(m.preview), previewRows)

	return strings.Join([]string{
		"", title, info, statusLine,
		preview,
		help,
	}, "\n")
}

func indent(s string) string { return "  " + strings.ReplaceAll(s, "\n", "\n  ") }

// padMin pads the block to at least minRows so the layout doesn't jump when the
// text is short — but never truncates, so multi-line output can grow freely.
func padMin(s string, minRows int) string {
	lines := strings.Split(s, "\n")
	for len(lines) < minRows {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}

// ============================================================================
// main
// ============================================================================

// input is the text after a flag, else standard input without its final
// newline.
func input(args []string) string {
	if len(args) > 0 {
		return strings.Join(args, " ")
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not read standard input:", err)
		os.Exit(1)
	}
	return strings.TrimRight(string(b), "\n")
}

func loadFont() *figFont {
	path := findFontFile()
	if path == "" {
		fmt.Fprintln(os.Stderr, "could not find "+fontName+".flf (expected in this folder or its parent)")
		os.Exit(1)
	}
	font, err := loadFigFont(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not load "+fontName+".flf:", err)
		os.Exit(1)
	}
	return font
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "-shape", "--shape":
			lines := strings.Split(input(args[1:]), "\n")
			for i, l := range lines {
				lines[i] = shape(l)
			}
			fmt.Println(strings.Join(lines, "\n"))
			return
		case "-print", "--print":
			fmt.Println(loadFont().renderText(input(args[1:])))
			return
		case "-h", "-help", "--help":
			fmt.Println("usage: nessnass-arabic [text] | -print [text] | -shape [text]")
			return
		}
	}

	word := startWord
	if len(args) > 0 {
		word = strings.Join(args, " ")
	}
	m := model{colors: palettes(), word: word, font: loadFont()}
	if _, err := tea.NewProgram(&m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
