# My Custom Notes

Can generate new patterns with REF: https://playgameoflife.com/

0. Edge of Grid Dead Zone.
rows 0 & rows $MaxHeight
rowPos 0 & rowPos $MaxWidth

1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.

...     ..D     LD.
.D.     .L.     L..
...     D..     .D.

2. Any live cell with more than three live neighbours dies, as if by overcrowding.

*.*     *.*
.D.     *D*
*.*     ...

3. Any live cell with two or three live neighbours lives on to the next generation.

.D.     .L.     D..
.L.     .LL     .LD
.D.     .L.     D..

4. Any dead cell with exactly three live neighbours becomes a live cell.

.*.     .*.     ...     *.*
*N.     .N*     *N*     .n.
.*.     .*.     .*.     *..

Conway's Game of Life uses the 8-cell [Moore neighborhood](https://en.wikipedia.org/wiki/Moore_neighborhood), which includes the diagonal neighbors. 
...
.*.
...

We need to calculate the count of neighbours and know the cell state dead/alive.
R1, $Neighbours 0 or 1 = Dies
R3, $Neighbours 2 or 3 = lives
R2, $Neighbours 4+ = Dies
R4, $Neighbours 3 = reanimate

    Y
    |
    |___ X
   /
  Z

# Timelog
Wrote test with 2 examples plus one I generated via link above. 1hr 30mins
Wrote code with test failing due to bad testdata so corrected. 5hrs 20mins
Total 6hrs 50 mins


https://cyber-dojo.org/kata/edit/L7yg3K
