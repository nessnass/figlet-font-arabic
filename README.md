# Nessnass Arabic figlet font

The Arabic letters of the [Nessnass figlet font](https://github.com/nessnass/figlet-font),
drawn with the same single-line Unicode box-drawing characters. Three rows
tall and right to left, for terminal banners and headers.
By [Nessnass](https://www.nessnass.com/?utm_source=github.com).

```
         │ ·     ·
 ╷  ┌┴┴┘ └─┴─┴┴┴─┘
 ╰──╯
```

نسناس, Nessnass.

## Use

Arabic letters change shape depending on where they stand in a word, and
figlet can't do that by itself. The font draws each form (Unicode Arabic
Presentation Forms-B), so shape the text first with the viewer's `-shape`,
then pass it to figlet:

```sh
cd viewer && go build && cd ..
figlet -f ./Nessnass-Arabic.flf -C utf8 "$(viewer/nessnass-arabic -shape 'نسناس')"
```

Or let the viewer draw it, with no figlet needed:

```sh
viewer/nessnass-arabic -print 'دبي'
echo 'شارع الخيل' | viewer/nessnass-arabic -print
```

figlet right-aligns right-to-left text to its width (80 by default). Add
`-w 200` for longer lines.

## Words

From right to left: دبي Dubai, أبوظبي Abu Dhabi, شارع الخيل Al Khail Road.

```
    │   ╶·─┐ │ │    ╭╴     │ ···            │╭·╮  ╭╮   ╭╴│             ╶╮
 ╷  ├──┴───┴─┘ ╵   ╭┘    ╷ └─┴┴┘     ╷··┌─┴─┴┴─╯  ╰┼─┘   ╵     ╷··┌─┘ ╶─╯
 ╰──╯ ··           ╰──╴ ╶╯           ╰──╯ ·      ╶─╯ ·         ╰──╯ ·
```

## Characters

The 28 letters, then ة ى, the hamza ء and its seats أ إ آ ؤ ئ, and the four
ligatures lam takes before alef. Each row reads right to left.

```
 ╶·─┐   ╶──┐   ╶──┐    ···    ··           │
 ╭──╯   ╭──╯   ╭──╯   └──┘   └──┘   └──┘   ╵
 ╰──╴   ╰──╴   ╰·─╴                  ·
      ╭──╮       ···              ·        ·╶╮    ╶╮
 ╷  ┌┴┴──╯   ╷  ┌┴┴┘   ╷  ┌┴┴┘    ╷    ╷   ╶─╯   ╶─╯
 ╰──╯        ╰──╯      ╰──╯      ╶╯   ╶╯
     ╭··╮        ╭·╮   ·╭╴     ╭╴     │╭·╮    │╭─╮        ╭·─╮
 ╷  ┌┴──╯   └────┴─╯   ╭┘     ╭┘     ╶┴┴─╯   ╶┴┴─╯   ╷  ┌┴┴──╯
 ╰──╯                  ╰──╴   ╰──╴                   ╰──╯
    ╶╮    ╭╮   ╭─╮     ·     ╭╮      │    ╶╴│
 ╷··╭╯    ╰┤   ╰─╯   ╷  ╷   ┌┴╯   ╷  │   └──┘
 ╰──╯    ╶─╯         ╰──╯   ╵     ╰──╯
  ╭╴    ╭╴╭╮   ╭┼╯    │   ╭╴│    ╭╴      ╶╮   ╭··╮
 ╷  ╷     ╰┤    ╵     ╵     ╵   ╶╯    ╷  ╭╯   ╰──╯
 ╰──╯    ╶─╯         ╭╴               ╰──╯
 ─╲│   ╲ │   ╭╴│   ╲ │
 ╰─╯   ╰─╯    ─╯   ╰─╯
       ╭╴
```

Letters that share a shape share its glyphs and differ only by their dots, as
in Arabic itself. A letter that doesn't join the next one (ا د ذ ر ز و) keeps
a blank column before it. Vowel marks aren't drawn and are left out when
shaping. Latin letters, digits and punctuation come from the Nessnass font.

## Viewer

A small Go app in `viewer/` that shapes and shows the font in the terminal as
you type.

```sh
cd viewer
go build
./nessnass-arabic
```

| Key | Action |
| --- | --- |
| type / backspace | edit the text |
| `enter` | new line |
| `tab` / `shift+tab` | colour |
| `ctrl+a` | show the alphabet |
| `ctrl+l` | clear |
| `ctrl+w` | save as `.txt` |
| `esc` / `ctrl+c` | quit |

## Licence

[MIT](LICENSE)
