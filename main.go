package main

import (
	"fmt"
	"os"
	"strings"
)

// ANSI renk kaçış kodları
const (
	Reset           = "\033[0m"  // Reset color
	Bold            = "\033[1m"  // Bold
	Dim             = "\033[2m"  // Dim
	Italic          = "\033[3m"  // Italic
	Underline       = "\033[4m"  // Underline
	Blink           = "\033[5m"  // Blink
	Reverse         = "\033[7m"  // Reverse (swap foreground and background colors)
	Hidden          = "\033[8m"  // Hidden (invisible)
	Strikethrough   = "\033[9m"  // Strikethrough
	Frame           = "\033[51m" // Framed
	Encircle        = "\033[52m" // Encircled
	Overline        = "\033[53m" // Overlined
	Background      = "\033[7m"  // Background color
	Black           = "\033[30m"
	Red             = "\033[31m"
	Green           = "\033[32m"
	Yellow          = "\033[33m"
	Blue            = "\033[34m"
	Magenta         = "\033[35m"
	Cyan            = "\033[36m"
	White           = "\033[37m"
	Gray            = "\033[90m"
	Brightred       = "\033[91m"
	Brightgreen     = "\033[92m"
	Brightyellow    = "\033[93m"
	Brightblue      = "\033[94m"
	Brightmagenta   = "\033[95m"
	Brightcyan      = "\033[96m"
	Brightwhite     = "\033[97m"
	Bgblack         = "\033[40m"  // Background black
	Bgred           = "\033[41m"  // Background red
	Bggreen         = "\033[42m"  // Background green
	Bgyellow        = "\033[43m"  // Background yellow
	Bgblue          = "\033[44m"  // Background blue
	Bgmagenta       = "\033[45m"  // Background magenta
	Bgcyan          = "\033[46m"  // Background cyan
	Bgwhite         = "\033[47m"  // Background white
	Bggray          = "\033[100m" // Background gray
	Bgbrightred     = "\033[101m" // Background bright red
	Bgbrightgreen   = "\033[102m" // Background bright green
	Bgbrightyellow  = "\033[103m" // Background bright yellow
	Bgbrightblue    = "\033[104m" // Background bright blue
	Bgbrightmagenta = "\033[105m" // Background bright magenta
	Bgbrightcyan    = "\033[106m" // Background bright cyan
	Bgbrightwhite   = "\033[107m" // Background bright white
	Blackbg         = "\033[40m"  // Background black
	Redbg           = "\033[41m"  // Background red
	Greenbg         = "\033[42m"  // Background green
	Yellowbg        = "\033[43m"  // Background yellow
	Bluebg          = "\033[44m"  // Background blue
	Magentabg       = "\033[45m"  // Background magenta
	Cyanbg          = "\033[46m"  // Background cyan
	Whitebg         = "\033[47m"  // Background white
	Orange          = "\033[38;5;208m"
)

func main() {
	if len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] || Example: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}

	colorFlag := os.Args[1]
	if len(colorFlag) < 8 || colorFlag[:7] != "--color" {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		return
	}

	var word string
	if len(os.Args) == 4 {
		word = os.Args[3]
	}
	if len(os.Args) == 3 {
		word = os.Args[2]
	}

	color := os.Args[1][8:]
	lettercolor := os.Args[2]

	// \n kaçış dizisini gerçek satır sonu karakterine dönüştür
	word = strings.ReplaceAll(word, "\\n", "\n")
	// Dosyayı oku
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Dosya okunurken bir hata oluştu")
		panic(err)
	}

	// Dosya içeriğini satırlara ayır
	lines := strings.Split(string(file), "\n")

	// Her kelime için ASCII sanatını ve renklendirmeyi yazdır
	for i, line := range strings.Split(word, "\n") {
		if line == "" {
			if i != 0 {
				fmt.Println()
			}
			continue
		}

		// Her bir karakteri işleyerek ASCII sanatını ve renklendirmeyi oluştur
		for h := 1; h < 9; h++ {
			for _, char := range line {
				colorizeAndPrintAsciiArtForCharacter(char, h, lines, color, lettercolor)
			}
			fmt.Println()
		}
	}
}

func colorizeAndPrintAsciiArtForCharacter(char rune, lineIndex int, lines []string, color, letters string) {
	// ASCII karakterinin indeksini hesapla
	index := (int(char) - 32) * 9
	// ASCII sanatını ve renkli çıktıyı yazdır
	if index >= 0 && index+8 <= len(lines) {
		if letters == "" || strings.ContainsRune(letters, char) {
			fmt.Print(getColorCode(color) + lines[index+lineIndex] + Reset)
		} else {
			fmt.Print(lines[index+lineIndex])
		}
	}
}

// Renk adına bağlı olarak ANSI renk kodunu al
func getColorCode(color string) string {
	switch strings.ToLower(color) {
	case "red":
		return Red
	case "green":
		return Green
	case "yellow":
		return Yellow
	case "blue":
		return Blue
	case "cyan":
		return Cyan
	case "white":
		return White
	case "bold":
		return Bold
	case "dim":
		return Dim
	case "italic":
		return Italic
	case "underline":
		return Underline
	case "blink":
		return Blink
	case "reverse":
		return Reverse
	case "hidden":
		return Hidden
	case "strikethrough":
		return Strikethrough
	case "frame":
		return Frame
	case "encircle":
		return Encircle
	case "overline":
		return Overline
	case "background":
		return Background
	case "magenta":
		return Magenta
	case "gray":
		return Gray
	case "brightred":
		return Brightred
	case "brightgreen":
		return Brightgreen
	case "brightyellow":
		return Brightyellow
	case "brightblue":
		return Brightblue
	case "brightmagenta":
		return Brightmagenta
	case "brightcyan":
		return Brightcyan
	case "brightwhite":
		return Brightwhite
	case "bgblack":
		return Bgblack
	case "bgred":
		return Bgred
	case "bggreen":
		return Bggreen
	case "bgyellow":
		return Bgyellow
	case "bgblue":
		return Bgblue
	case "bgmagenta":
		return Bgmagenta
	case "bgcyan":
		return Bgcyan
	case "bgwhite":
		return Bgwhite
	case "bggray":
		return Bggray
	case "bgbrightred":
		return Bgbrightred
	case "bgbrightgreen":
		return Bgbrightgreen
	case "bgbrightyellow":
		return Bgbrightyellow
	case "bgbrightblue":
		return Bgbrightblue
	case "bgbrightmagenta":
		return Bgbrightmagenta
	case "bgbrightcyan":
		return Bgbrightcyan
	case "bgbrightwhite":
		return Bgbrightwhite
	case "blackbg":
		return Blackbg
	case "redbg":
		return Redbg
	case "greenbg":
		return Greenbg
	case "yellowbg":
		return Yellowbg
	case "bluebg":
		return Bluebg
	case "magentabg":
		return Magentabg
	case "cyanbg":
		return Cyanbg
	case "whitebg":
		return Whitebg
	case "orange":
		return Orange

	default:
		// Renk tanınamazsa varsayılan olarak beyaz
		return White
	}
}
