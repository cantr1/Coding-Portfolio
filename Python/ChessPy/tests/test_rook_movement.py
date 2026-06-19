import unittest
from pieces import Rook
from position import Position
from movement_strategy import MovementException

def create_rook(x, y, has_moved = False):
    return Rook(Position(x, y), has_moved)

class TestRook(unittest.TestCase):
    def test_move_off_board1(self):
        p = create_rook(1, 1)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=-1, y_pos=1)
            p.move(target_pos)
    
    def test_move_off_board2(self):
        p = create_rook(8, 1)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=8, y_pos=9)
            p.move(target_pos)

    def test_invalid_move_both_axis(self):
        p = create_rook(3, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=4, y_pos=5)
            p.move(target_pos)