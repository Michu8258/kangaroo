# Nomenclature

This document assumes thet you already know what a Sudoku is. If not, you can read more [at Wikipedia](https://en.wikipedia.org/wiki/Sudoku) or [here](https://www.britannica.com/topic/sudoku). Please find this document as explanationf of types, types names, variables names that appear in the code and are necessary to understand what is going on. So, without a furher ado, basic definitions.

**Example sudoku layout**:

<img src="./images/Nomenclature.svg" alt="Nomenclature picture" width="300"/>

### Sudoku

Sudoku is a placeholder for **Boxes** that is introduced for organizational purposes, easier management and processing. There are 2 important data pieces when it comes to Sudoku:

- **Box size** - this is one dimension of a sudoku box (square). In case of classic sudoku puzzle, it is 3 - it simply lets us define sudoku puzzle with different characters sets - by default this set has 9 characters (numbers from 1 to 9), but nothing prevents us from solving a sudoku, where each box has 25 fields (box size = 5), and then compose layout of boxes same size. **This value mus be in range 2 - 5 (both sides inclusive).**
- **Layout** - a pair of integers that describes layout of a sudoku puzzle. Those do not have to be equal to each other, meaning that sudoku layout must be rectangle (not necessarily a square like in case of classic sudoku puzzle - 3 and 3). Layout indicates how many boxes are included in the sudoku puzzle. In this case (image above) the layout has **width of 4** and **height of 4** as well, so there are 16 boxes of box size = 3. Both lyout width and height must be in range from 2 to 10 inclusively.

**layout important note**: Only sudoku boxes that **are not** disabled (not considered a box at all) are considered part of the whole puzzle (in the image above disabled boxes has no cells and a gray background). So all boxes that are considered part of a puzzle must meet the following criteria:

- every box must be a part of at least one **sub-sudoku** (there are 2 3x3 boxes sub-sudokus in the above image one consists of blue and yellow boxes, the other of red and yellow boxes) - this allows to build wild layouts.
- sub-sudoku will always consist of boxes square in size **n boxes x n boxes**. What does _n_ mean? _n_ is a Box size so **in case of box size 3, we expect to find at least one sub-sudoku with size 3x3 boxes and each of those boxes should be 3x3 cells.** In case of the above image we have 2 sub-sudokus with box size 3.

### Box (SudokuBox)

Box is a part of sudoku puzzle, all of the boxes share same **Box size** and all of them are squares - always. A box can be **disabled** - which means that it should not be considered a part of the whole puzzle (gryed out areas in the image above).

Every sudoku box has 2 coordinates, **IndexRow** and **IndexColumn** - and those are introduced for ease of localization and querying boxes from the collection. The indexing starts in top left corner of the layout, when both indices have a value of 0, and ends in bottom right corner, where **IndexRow** = Sudoku.Layout.Height and **IndexColumn** = Sudoku.Layout.Width. This lets us build sub-sudokus objects for better rows and columns (cells) management.

So rows indexing goes from top to bottom (0 based) and columns indexing goes from left to right (0 based).

The most important part of a Box is a collection of Cells, where each cell represents single sudoku character container (graphically - the smallest square in the image above - blue, red or yellow).

### Cell (SudokuCell)

Sudoku cell represents a placeholder for single number (character) in the smallest cell os the sudoku puzzle. Every cell has its value or does not have a value at all (nil pointer), but also very important properties of a cell are `IndexRowInBox` and `IndexColumnInBox`. Thise indexes allow s to uniqely identify cell in the box. Similar to box's indexes, these are as well counted from left to right (columns) and from top to bottom (rows).

So in case of a box with size = 3, top left cell will have indexex 0 and 0, and top right cell with have row index of 0 and column index of 2.

Potential Values is a collection of possible values that could be placed in the cell - that is if the cell does not have a value assigned yet. The value can be assigned as a data input - part of the input data or it can be assigned during solving process. Potential values slice is always a nil collection at the solving process start and is populated during solution processing - this collectio is entirely managed by a program.

Box holds a reference to box, which spoken cell is part of - this will speed up querying data.

`Member of lines` - every cell is a part of column and row in the Sudoku puzzle - this slice is holding references to containers that hold references to all cells within a row or a column (there is no distinguish for columns and row, because it simply is irrelevant). This slice is again introduced solely for algorith purposes and is not a part of initial data input. Thanks to such collections it will be fast to check if some character (number) in the puzzle already appeared in the row or column - this will eliminate need of looping all the time in search of another cells within the same row/column.

### Line (SudokuLine)

Sudoku line is just a slice that holds references to all cells that are part of a sudoku cell or row. This type is going to help avoid constant looping through boxex and cells. `SudokuLine` slices are constructed upon solution start and are never rebuilt again.