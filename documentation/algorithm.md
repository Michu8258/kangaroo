# Solving algorithm

## introduction

This application uses Crook's method which core feature is to use preemptive sets as a way to quickly and reliably exclude some values from list of **potential values** that can be assigned to a given cell. First of all - what is a potential value? This is avalue that could be assigned to a cell and would not break any sudoku rule (unique values in column, row and box; all values within permitted values range). In Cook's method, potential values are named a cell's markup, but this codebase uses name _potential values_ as it is more accurate.

**Preemptive set** is a concept that it is necessary for you to know if you want to understand the codebase. This rule applies to every box, column and row, but for simplicity we will discus one case, lets say a column (other cases are similar.)

Let's say you have a column with the following values inside (and you already found out about potential values):

| cell number | cell value | cell potential values |
|-------------|------------|-----------------------|
| 1 | **5** |
| 2 | | [2, 4]
| 3 | | [6, 7]
| 4 | **8** |
| 5 | | [2, 4]
| 6 | | [2, 6, 7]
| 7 | **3** |
| 8 | | [2, 4, 7]
| 9 | **1** |

As you can see, we have two cells: **2** and **5** with exacty the same potential values set: `[2, 4]`. It means that at this point you can be certain that those two cells will for sure hold values 2 and 4, therefore all other cells in this column will have values that are different than 2 and different than 4. So after aplying this elimination logic (applying preemptive set), you can exclude potential values of other cells in same column to the state presented below:

| cell number | cell value | cell potential values |
|-------------|------------|-----------------------|
| 1 | **5** |
| 2 | | [2, 4]
| 3 | | [6, 7]
| 4 | **8** |
| 5 | | [2, 4]
| 6 | | [6, 7]
| 7 | **3** |
| 8 | | [7]
| 9 | **1** |

Now, after applying preemptive set logic, you found out, that only possible value in cell **8** is 7, and it is certain.

**Exaclty same logic can be applied to rows and boxes!**

## Preemptive sets rules

Please bare in mind that if any of the following rules is met, the pair (or collection) of cells considered a preemptive set **are not** trully a preemptive set:

- there is no more cell with potential values (with no value)
- there is at least one other cell in collection that has amount of potential values less than length of potential values of preemptive set. Let's say you found a preemptive set with potential values [1, 2, 3], if there is a cell with 2 potential values in the collection, the set cannot be considered a preemptive one
- preemptive set must consist of amount of cells that is equal to length of potential values of preemptive set. If you found a preemptive set with 2 values, you need 2 cells, if the set has 3 values, you need 3 cells, and so on...
- if you have set with 3 values already found, but you also found set with 2 values, the set with 2 values wins
- if any cell in cells collection has 1 potential value, you are lucky, because you found a certain value

You can find more theory about preemptive sets and Crook's method [here](https://pi.math.cornell.edu/~mec/Summer2009/meerkamp/Site/Solving_any_Sudoku_II.html) and [here](https://www.ams.org/notices/200904/rtx090400460p.pdf).

## The algorithm

1. **Eliminations logic** - assign all potential values, if there are cells, with only one potential values (certain values at this point), assign those values to cells. Then loop as long as you have no cell left without values (sudoku solved, finish) or there is no cell with only one potential value -> all cells without value have at least two potential values. Go to step 2.

2. **Preemptive sets** - find preemptive sets within boxes, rows and cells. If you found any, apply the sets just like described in the **introduction** section - it will probably help you reduce potential values in the cells. After the reduction, if there is at least one cell with only one potential value (certain at this point), assign it. And go back to step one. But if you still has no cell with only one potential value, and no more preemptive set exists, go to step 3.

**UP TO THIS POINT A CELL WITH NO POTENTIAL VALUES MEANS NO SOLUTION FOR A SUDOKU**

3. **Guessing** - now the neat part, we have to make a guess. Select a cell with least amount of potential values (most commonly - with 2 potential values), and select one of potential values, assign it as cell value, clear potential values of that cell and **remember what cell and how you modified**. Go again to step 1. _**There is a difference though**_ - now you may have a cell with no potential values - but from now on it will mean you guessed the wrong value. So if you'll encounter such case. you have to restore sudoku state to the point where you made a quess (hence the remembering part), clean value and restore potential values of the cell where you made a guess. Now you can delete a potential value from a collection of of potentoal values of the cell, because now you know it is wrong. And go back to step 1. If all combination (recursive as well) seem to fail, the sudoku is unsolvable.