import unittest
from pieces import Piece
from position import Position
from movement_strategy import MovementException, OffBoardException, KingMovement

def create_king(x, y, has_moved = False):
    return Piece(Position(x, y), KingMovement(), has_moved)

class TestKing(unittest.TestCase):
    def test_move_off_board1(self):
        p = create_king(1, 1)
        with self.assertRaises(OffBoardException):
            target_pos = Position(x_pos=-1, y_pos=1)
            p.move(target_pos)
    
    def test_move_off_board2(self):
        p = create_king(8, 1)
        with self.assertRaises(OffBoardException):
            target_pos = Position(x_pos=8, y_pos=9)
            p.move(target_pos)
    
    def test_invalid_movement_none(self):
        p = create_king(6, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=6, y_pos=4)
            p.move(target_pos)

    def test_valid_movement_forward(self):
        p = create_king(1, 3)
        target_pos = Position(x_pos=1, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
    
    def test_valid_movement_backwards(self):
        p = create_king(5, 4)
        target_pos = Position(x_pos=5, y_pos=3)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
    
    def test_valid_movement_sideways(self):
        p = create_king(5, 4)
        target_pos = Position(x_pos=6, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)

    def test_valid_movement_diagonal(self):
        p = create_king(5, 4)
        target_pos = Position(x_pos=4, y_pos=5)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
