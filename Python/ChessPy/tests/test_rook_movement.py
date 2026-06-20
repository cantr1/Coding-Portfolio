import unittest
from pieces import Piece
from position import Position
from movement_strategy import MovementException, OffBoardException, RookMovement

def create_rook(x, y, has_moved = False):
    return Piece(Position(x, y), RookMovement(), has_moved)

class TestRook(unittest.TestCase):
    def test_move_off_board1(self):
        p = create_rook(1, 1)
        with self.assertRaises(OffBoardException):
            target_pos = Position(x_pos=-1, y_pos=1)
            p.move(target_pos)
    
    def test_move_off_board2(self):
        p = create_rook(8, 1)
        with self.assertRaises(OffBoardException):
            target_pos = Position(x_pos=8, y_pos=9)
            p.move(target_pos)

    def test_invalid_move_both_axis(self):
        p = create_rook(3, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=4, y_pos=5)
            p.move(target_pos)
    
    def test_invalid_movement_none(self):
        p = create_rook(6, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=6, y_pos=4)
            p.move(target_pos)
    
    def test_valid_move_y(self):
        p = create_rook(3, 4)
        target_pos = Position(x_pos=3, y_pos=8)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)
    
    def test_valid_move_x(self):
        p = create_rook(3, 4)
        target_pos = Position(x_pos=7, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)