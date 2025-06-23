package main

import (
	"bufio"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/term"
)

//go:embed db.json
var database string

// Estructura para cada comando
type CommandInfo struct {
	Comando     string `json:"comando"`
	Descripcion string `json:"descripcion"`
	Info        string `json:"info"`
}

var commandsDB []CommandInfo

func loadCommandsDB() error {

	return json.Unmarshal([]byte(database), &commandsDB)
}

func findCommandInfo(cmd string) *CommandInfo {
	for _, c := range commandsDB {
		if c.Comando == cmd {
			return &c
		}
	}
	return nil
}

func main() {
	// Cargar base de datos de comandos
	if err := loadCommandsDB(); err != nil {
		fmt.Println("Error cargando db.json:", err)
		os.Exit(1)
	}

	// Almacenar contenido de los paneles
	leftPanel := []string{}
	rightPanel := []string{"Escribe un comando (ej: ls) → Verás su explicación aquí"}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearScreen()
		renderPanels(leftPanel, rightPanel)

		fmt.Print("SHELL> ")
		if !scanner.Scan() {
			break
		}
		cmdInput := scanner.Text()

		if cmdInput == "exit" {
			break
		}

		// Ejecutar comando
		parts := strings.Fields(cmdInput)
		if len(parts) == 0 {
			continue
		}

		cmd := exec.Command(parts[0], parts[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			output = []byte(fmt.Sprintf("Error: %v", err))
		}

		// Actualizar panel izquierdo
		leftPanel = append(leftPanel, fmt.Sprintf("$ %s", cmdInput))
		leftPanel = append(leftPanel, string(output))

		// Actualizar panel derecho
		cinfo := findCommandInfo(parts[0])
		if cinfo != nil {
			rightPanel = []string{
				fmt.Sprintf("Comando: %s", cinfo.Comando),
				"",
				cinfo.Descripcion,
				"",
				cinfo.Info,
			}
		} else {
			rightPanel = []string{"Comando desconocido. Usa comandos comunes como ls, pwd, etc."}
		}
	}
}

// Limpiar pantalla
func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

// Renderizar los dos paneles
func renderPanels(left, right []string) {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	rightPanelWidth := int(float64(width) * 0.25)
	if rightPanelWidth < 10 {
		rightPanelWidth = 10 // ancho mínimo para legibilidad
	}
	leftPanelWidth := width - rightPanelWidth - 3 // -3 por el separador y márgenes
	panelHeight := height - 2

	// Procesar panel izquierdo (últimas líneas que caben)
	start := 0
	if len(left) > panelHeight {
		start = len(left) - panelHeight
	}
	leftContent := strings.Join(left[start:], "\n")

	// Ajustar contenido a paneles
	leftLines := wrapLines(leftContent, leftPanelWidth, panelHeight)
	rightLines := wrapLines(strings.Join(right, "\n"), rightPanelWidth, panelHeight)

	maxLines := max(len(leftLines), len(rightLines))

	// Combinar paneles línea por línea
	for i := 0; i < maxLines; i++ {
		leftLine := safeGetLine(leftLines, i, leftPanelWidth)
		rightLine := safeGetLine(rightLines, i, rightPanelWidth)
		fmt.Printf("%s | %s\n", leftLine, rightLine)
	}
}

func safeGetLine(lines []string, idx, width int) string {
	if idx < len(lines) {
		return text.AlignLeft.Apply(lines[idx], width)
	}
	return strings.Repeat(" ", width)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// wrapLines corta el texto en líneas y ajusta la altura
func wrapLines(text string, width, height int) []string {
	lines := []string{}
	for _, line := range strings.Split(text, "\n") {
		for len(line) > width {
			lines = append(lines, line[:width])
			line = line[width:]
		}
		lines = append(lines, line)
	}
	// Ajustar altura
	if len(lines) > height {
		lines = lines[len(lines)-height:]
	}
	for len(lines) < height {
		lines = append(lines, "")
	}
	return lines
}
