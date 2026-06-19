import unittest
from pieces import *

def create_pawn(x, y, has_moved = False):
    return Pawn(Position(x, y), has_moved)

class TestPawn(unittest.TestCase):
    def test_move_off_board1(self):
        p = create_pawn(6, 8)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=6, y_pos=9)
            p.move(target_pos)
    
    def test_move_off_board2(self):
        p = create_pawn(1, 1)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=1, y_pos=0)
            p.move(target_pos)

    def test_invalid_movement_large_forward(self):
        p = create_pawn(1, 2)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=1, y_pos=8)
            p.move(target_pos)
    
    def test_invalid_movement_backwards(self):
        p = create_pawn(3, 6)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=3, y_pos=5)
            p.move(target_pos)
    
    def test_invalid_movement_none(self):
        p = create_pawn(6, 2)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=6, y_pos=2)
            p.move(target_pos)
    
    def test_invalid_movement_diagonal(self):
        # This will change once captures are implemented
        p = create_pawn(4, 6)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=5, y_pos=7)
            p.move(target_pos)
    
    def test_invalid_movement_sideways(self):
        p = create_pawn(6, 8)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=5, y_pos=8)
            p.move(target_pos)

    def test_invalid_movement_forward2(self):
        p = create_pawn(6, 3, has_moved=True)
        with self.assertRaises(MovementException):
            target_pos = Position(x_pos=6, y_pos=5)
            p.move(target_pos)

    def test_valid_movement_forward2(self):
        p = create_pawn(1, 2)
        target_pos = Position(x_pos=1, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)
    
    def test_valid_movement_forward1(self):
        p = create_pawn(1, 3)
        target_pos = Position(x_pos=1, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)

    def test_valid_multiple_moves(self):
        p = create_pawn(1, 2)
        target_pos = Position(x_pos=1, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)

        target_pos2 = Position(x_pos=1, y_pos=5)
        p.move(target_pos2)
        self.assertEqual(p.position.x_pos, target_pos2.x_pos)
        self.assertEqual(p.position.y_pos, target_pos2.y_pos)

    def test_invalid_multiple_moves(self):
        p = create_pawn(1, 2)
        target_pos = Position(x_pos=1, y_pos=4)
        p.move(target_pos)
        self.assertEqual(p.position.x_pos, target_pos.x_pos)
        self.assertEqual(p.position.y_pos, target_pos.y_pos)

        with self.assertRaises(MovementException):
            target_pos2 = Position(x_pos=1, y_pos=6)
            p.move(target_pos2)


if __name__ == "__main__":
    unittest.main()
