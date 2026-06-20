import unittest
from pieces import Piece
from position import Position
from movement_strategy import MovementException, OffBoardException, QueenMovement

def create_queen(x, y, has_moved = False):
    return Piece(Position(x, y), QueenMovement(), has_moved)

class TestQueen(unittest.TestCase):
    def test_move_off_board1(self):
        p = create_queen(1, 1)
        with self.assertRaises(OffBoardException):
            target_pos = Position(x_pos=-1, y_pos=1)
            p.move(target_pos)
    
    def test_move_off_board2(self):
        p = create_queen(8, 1)
        with self.assertRaises(OffBoardException):
            target_pos = Position(x_pos=8, y_pos=9)
            p.move(target_pos)
    
    def test_invalid_movement_none(self):
        p = create_queen(6, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=6, y_pos=4)
            p.move(target_pos)

    def test_invalid_movement1(self):
        p = create_queen(2, 1)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=3, y_pos=3)
            p.move(target_pos)
    
    def test_invalid_movement2(self):
        p = create_queen(3, 6)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=2, y_pos=4)
            p.move(target_pos)
    
    def test_invalid_movement3(self):
        p = create_queen(6, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=4, y_pos=3)
            p.move(target_pos)

    def test_invalid_move_both_axis(self):
        p = create_queen(3, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=8, y_pos=6)
            p.move(target_pos)

    def test_valid_move_y(self):
        p = create_queen(3, 4)
        target_pos = Position(x_pos=3, y_pos=8)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
    
    def test_valid_move_x(self):
        p = create_queen(3, 4)
        target_pos = Position(x_pos=7, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
    
    def test_valid_move1(self):
        p = create_queen(2, 1)
        target_pos = Position(x_pos=5, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
    
    def test_valid_move2(self):
        p = create_queen(7, 1)
        target_pos = Position(x_pos=4, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
    
    def test_valid_move3(self):
        p = create_queen(6, 3)
        target_pos = Position(x_pos=4, y_pos=1)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
