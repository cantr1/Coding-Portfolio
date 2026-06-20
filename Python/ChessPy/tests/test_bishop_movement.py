import unittest
from pieces import Piece
from position import Position
from movement_strategy import MovementException, OffBoardException, BishopMovement

def create_bishop(x, y, has_moved = False):
    return Piece(Position(x, y), BishopMovement(), has_moved)

class TestBishop(unittest.TestCase):
    def test_move_off_board1(self):
        p = create_bishop(3, 1)
        with self.assertRaises(OffBoardException):
            target_pos = Position(x_pos=2, y_pos=0)
            p.move(target_pos)

    def test_move_off_board2(self):
        p = create_bishop(3, 1)
        with self.assertRaises(OffBoardException):
            target_pos = Position(x_pos=4, y_pos=0)
            p.move(target_pos)
    
    def test_invalid_movement1(self):
        p = create_bishop(2, 1)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=3, y_pos=1)
            p.move(target_pos)
    
    def test_invalid_movement2(self):
        p = create_bishop(3, 6)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=2, y_pos=4)
            p.move(target_pos)
    
    def test_invalid_movement3(self):
        p = create_bishop(6, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=4, y_pos=3)
            p.move(target_pos)
    
    def test_invalid_movement_none(self):
        p = create_bishop(6, 4)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=6, y_pos=4)
            p.move(target_pos)
    
    def test_valid_move1(self):
        p = create_bishop(2, 1)
        target_pos = Position(x_pos=5, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
    
    def test_valid_move2(self):
        p = create_bishop(7, 1)
        target_pos = Position(x_pos=4, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
    
    def test_valid_move3(self):
        p = create_bishop(6, 3)
        target_pos = Position(x_pos=4, y_pos=1)
        p.move(target_pos)
        self.assertEqual(p.position, target_pos)
