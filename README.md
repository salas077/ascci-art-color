ASCII Art Generator with Color

This program converts text into ASCII art with optional color support.
Made for Zone01 Athens by Giorgos Salaounis and Christos Paloglou.

HOW TO USE

Basic:
  go run . "Hello"

With a banner:
  go run . "Hello" shadow

With color (entire text):
  go run . --color=red "Hello"

With color (specific part):
  go run . --color=blue kit "kitten"

AVAILABLE COLORS

red, green, yellow, blue, magenta, cyan, white, orange, black

AVAILABLE BANNERS

standard, shadow, thinkertoy

EXAMPLES

Normal text:
  go run . "Hello World"

Text with newlines:
  go run . "Hello\nWorld"

Color everything:
  go run . --color=green "Hello"

Color one letter:
  go run . --color=red H "Hello"

Color a substring:
  go run . --color=yellow kit "a kitten has a kit"

Different banner with color:
  go run . --color=blue o "Hello" thinkertoy

WHAT IT DOES

Takes your text and turns it into big ASCII art letters. You can choose 
different styles (banners) and add colors to make parts stand out.

Each character is made from 8 lines of smaller characters. The program 
reads template files to know how to draw each letter.

When you add color, it puts special codes around the letters so your 
terminal shows them in color.

ERRORS

If you mess up the format, it shows you how to use it correctly.
If you pick a color that doesn't exist, it lists the real colors.
If the banner file is missing, it tells you.

TESTING

To run tests:
  go test
  go test -v

FILES

main.go - starts the program
banner.go - loads the letter templates
render.go - draws the ASCII art
color.go - handles colors
main_test.go - tests the basic stuff
color_test.go - tests the color stuff
