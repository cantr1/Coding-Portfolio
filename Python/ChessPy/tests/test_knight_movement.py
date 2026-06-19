import unittest
from pieces import Piece
from position import Position
from movement_strategy import MovementException, KnightMovement

def create_knight(x, y, has_moved = False):
    return Piece(Position(x, y), KnightMovement(), has_moved)

class TestRook(unittest.TestCase):
    def test_move_off_board1(self):
        p = create_knight(2, 1)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=-1, y_pos=-1)
            p.move(target_pos)
    
    def test_move_off_board2(self):
        p = create_knight(8, 1)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=9, y_pos=3)
            p.move(target_pos)
    
    def test_valid_move1(self):
        p = create_knight(2, 1)
        target_pos = Position(x_pos=3, y_pos=3)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)
    
    def test_valid_move2(self):
        p = create_knight(7, 1)
        target_pos = Position(x_pos=6, y_pos=3)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)
    
    def test_valid_move3(self):
        p = create_knight(6, 3)
        target_pos = Position(x_pos=4, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)