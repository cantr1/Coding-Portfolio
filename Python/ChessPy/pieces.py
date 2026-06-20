"""
Main file for building game pieces
"""
from abc import ABC, abstractmethod
from position import Position
from movement_strategy import MovementException, MovementBehavior, PawnMovement, RookMovement

class Piece:
    def __init__(self, p: Position, mb: MovementBehavior, has_moved: bool = False):
        self.position = p
        self.has_moved = has_moved
        self.mb = mb
        # TODO: Implement color

    def check_valid_move(self, target: Position) -> bool:
        return self.mb.check_valid_move(self.position, target, self.has_moved)
    
    def move(self, target: Position) -> None:
        if self.mb.check_valid_move(self.position, target, self.has_moved):
            self.position = target
            if not self.has_moved:
                self.has_moved = True
        else:
            raise MovementException("Invalid Move")
    
    def captured(self):
        raise NotImplementedError
