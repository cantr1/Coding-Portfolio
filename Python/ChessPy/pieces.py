"""
Main file for building game pieces
"""
from abc import ABC, abstractmethod
from position import Position
from movement_strategy import MovementException, PawnMovement, RookMovement

class Piece(ABC):
    @abstractmethod
    def check_valid_move(self, target: Position) -> bool:
        pass

    @abstractmethod
    def move(self) -> None:
        pass

    def captured(self) -> None:
        # TODO: Decide on how to remove a piece from the board
        raise NotImplementedError

class Pawn(Piece):
    def __init__(self, p: Position, has_moved: bool = False):
        self.position = p
        self.has_moved = has_moved
        self.mb = PawnMovement()
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
        return super().captured()

class Rook(Piece):
    def __init__(self, p: Position, has_moved: bool = False):
        self.position = p
        self.has_moved = has_moved
        self.mb = RookMovement()
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
        return super().captured()