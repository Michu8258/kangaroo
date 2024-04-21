# Sudoku Binary Format

Since this application is a CLI it is required to be able to call commands without the need of creating a sudoku configuration file and then reading result from another file. There fore, the binary representation of sudoku data is introduced and detailed in this documentation.

**What data about sudoku we need to represent the sudoku?**

- box size
- layout width and height
- collection of boxes data
- if box is disabled
- collection of values - value per cell in box

## Binary representation design

| Chunk number | Data | Bytes count | Description |
|--------------|------|-------------|-------------|
| 1 | Version | 2 | For now, we only have first version but it is asways a good idea to include version information anytime you deal with binary representation of any data.
| 2 | Box size | 1 | This is sudoku configuration related information - required. There is no need for 2 or more bytes, as the CLI supports box size of 5 max.
| 3 | Layout Width & Height | 2 | Two bytes for layout data - **first for width, second for height**. One byte per dimension is sufficient as the CLI support maz layout size of 5.
| 4 | Box disable data | `Math.ceil((layout.width * layout.height) / 8)` | This is a mask for amount of boxes. In case of layout width = 3 and layout height = 3 (classic sudoku) we have 9 boxes, so we need 9 bits to represent enabled/disabled state of a box -> we need 2 bytes to hold that information. Index of bit in the value indicates index of a box the bit reffers to. **So amount of bytes to hold this information varies.** Example: if all boxes are enabled: `11111111 10000000`. If box with index 2 is disabled, then we expect the value: `11011111 10000000`;
| 5 | Boxes data | `(enabled boxes count) * (box size) * (box size)` | If box is disabled (see **4**) then no data should be inclided in this data chunk for that box. Otherwise we only need values for the sudoku cells. One byte per sudoku value is sufficient, for this CLI supports box size up to 5, which gives max value of 25. 255 is plenty more. So in case of classic sudoku, amount of bytes per single box of sudoku will be 9. **If cell has no value, we expect 0 as a byte value.** In case of classic sudoku, and all boxes enabled, this chunk of data should be 81 bytes long.

**Important note: Indexes both, of boxes and cells in boxes, are incrementing columns first, then rows. In case of classic sudoku:**

```
0 1 2
3 4 5
6 7 8
```


## Concrete Example

Consider the following sudoku (a classic one):

```
╔═══════════╦═══════════╦═══════════╗
║ 6 │   │   ║   │   │ 3 ║ 4 │   │   ║
║───────────║───────────║───────────║
║   │ 1 │   ║ 2 │   │ 5 ║   │ 7 │   ║
║───────────║───────────║───────────║
║   │   │ 7 ║   │   │   ║   │   │ 1 ║
║═══════════╬═══════════╬═══════════║
║   │   │   ║ 1 │ 8 │   ║ 5 │   │   ║
║───────────║───────────║───────────║
║ 9 │   │   ║   │   │   ║   │ 8 │   ║
║───────────║───────────║───────────║
║   │ 4 │   ║ 6 │   │ 7 ║   │   │   ║
║═══════════╬═══════════╬═══════════║
║   │   │ 6 ║   │   │   ║   │   │ 3 ║
║───────────║───────────║───────────║
║   │ 8 │   ║   │ 3 │   ║   │ 2 │   ║
║───────────║───────────║───────────║
║ 2 │   │   ║ 5 │ 6 │   ║ 7 │   │   ║
╚═══════════╩═══════════╩═══════════╝
```

| Chunk number | Data | Bytes | Comment |
|--------------|------|-------|---------|
| 1 | Version | 0 1 | Only supported version for now |
| 2 | Box size | 3 | - |
| 3 | Layout Width & Height | 3 3 | width height |
| 4 | Box disable data | 255 128 | `11111111 10000000` in binary (all boxes enabled)
| 5 | Boxes data | [] | `Look below` |

```
6 0 0 0 1 0 0 0 7
0 0 3 2 0 5 0 0 0
4 0 0 0 7 0 0 0 1
0 0 0 9 0 0 0 4 0
1 8 0 0 0 0 6 0 7
5 0 0 0 8 0 0 0 0
0 0 6 0 8 0 2 0 0
0 0 0 0 3 0 5 6 0
0 0 3 0 2 0 7 0 0
```

To sum up, in this particular case, bytes representation is (88 bytes):

```
0 1 3 3 3 255 128 6 0 0 0 1 0 0 0 7 0 0 3 2 0 5 0 0 0 4 0 0 0 7 0 0 0 1 0 0 0 9 0 0 0 4 0 1 8 0 0 0 0 6 0 7 5 0 0 0 8 0 0 0 0 0 0 6 0 8 0 2 0 0 0 0 0 0 3 0 5 6 0 0 0 3 0 2 0 7 0 0
```

### And finally, presented as base64 string:

To make easy prowiding the data to the CLI, conversion to base64 string should be made to easily pass sudoku data as a string to the CLI. CLI should also respond with base64 encoded solved sudoku data.

```
AAEDAwP/gAYAAAABAAAABwAAAwIABQAAAAQAAAAHAAAAAQAAAAkAAAAEAAEIAAAAAAYABwUAAAAIAAAAAAAABgAIAAIAAAAAAAADAAUGAAAAAwACAAcAAA==
```

### Go conversion example:

```go
    sudokuData := []byte{
		0, 1,
		3,
		3, 3,
		255, 128,
		6, 0, 0, 0, 1, 0, 0, 0, 7,
		0, 0, 3, 2, 0, 5, 0, 0, 0,
		4, 0, 0, 0, 7, 0, 0, 0, 1,
		0, 0, 0, 9, 0, 0, 0, 4, 0,
		1, 8, 0, 0, 0, 0, 6, 0, 7,
		5, 0, 0, 0, 8, 0, 0, 0, 0,
		0, 0, 6, 0, 8, 0, 2, 0, 0,
		0, 0, 0, 0, 3, 0, 5, 6, 0,
		0, 0, 3, 0, 2, 0, 7, 0, 0,
	}

	base64Representation := base64.StdEncoding.EncodeToString(sudokuData)
	sudokuBytes, _ := base64.StdEncoding.DecodeString(base64Representation)
	fmt.Println(base64Representation)
	fmt.Println(sudokuBytes)

    // Output:
    // AAEDAwP/gAYAAAABAAAABwAAAwIABQAAAAQAAAAHAAAAAQAAAAkAAAAEAAEIAAAAAAYABwUAAAAIAAAAAAAABgAIAAIAAAAAAAADAAUGAAAAAwACAAcAAA==
    // [0 1 3 3 3 255 128 6 0 0 0 1 0 0 0 7 0 0 3 2 0 5 0 0 0 4 0 0 0 7 0 0 0 1 0 0 0 9 0 0 0 4 0 1 8 0 0 0 0 6 0 7 5 0 0 0 8 0 0 0 0 0 0 6 0 8 0 2 0 0 0 0 0 0 3 0 5 6 0 0 0 3 0 2 0 7 0 0]
```