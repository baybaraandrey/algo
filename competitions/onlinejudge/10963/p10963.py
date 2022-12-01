import sys
from collections import namedtuple


stdin = sys.stdin

Col = namedtuple('Col', 'y1 y2')


def can_gap_closed(cols):
    prev_diff = None
    for col in cols:
        curr_diff = abs(col.y1 - col.y2)
        if prev_diff is None:
            prev_diff = curr_diff
            continue
        if curr_diff != prev_diff:
            return 'no'
    return 'yes'

if __name__ == '__main__':
    cases = int(stdin.readline())
    for _ in range(cases):
        n_cols = int(stdin.readline())
        cols = [None] * n_cols
        for j in range(n_cols):
            y1, y2 = [int(x) for x in stdin.readline().split()]
            cols[j] = Col(y1, y2)
        print(can_gap_closed(cols))
