from abc import ABC, abstractmethod

class MovementException(Exception):
    pass

class Position:
    def __init__(self, x_pos: int, y_pos: int):
        self.x_pos = x_pos
        self.y_pos = y_pos

"""
Using the Startegy Design pattern to handle pieces
"""
class Piece(ABC):
    def __init__(self, p: Position, has_moved: bool = False) -> None:
        self.position = p
        self.has_moved = has_moved

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
        super().__init__(p, has_moved)
    
    def check_valid_move(self, target: Position) -> bool:
        # test move is on board (8x8)
        if target.x_pos > 8 or target.y_pos > 8 or target.x_pos < 1 or target.y_pos < 1:
            return False
        if target.x_pos != self.position.x_pos:
            return False  # TODO: implement a board class that can return if a piece is at the position for captures
        
        y_pos_change: int = target.y_pos - self.position.y_pos
        
        if y_pos_change > 2:
            return False
        if y_pos_change <= 0:
            return False
        else:
            if (y_pos_change == 2 and not self.has_moved) or y_pos_change == 1:
                return True
            else:
                return False
            
    def move(self, target: Position) -> None:
        if self.check_valid_move(target):
            self.position = target
            if not self.has_moved:
                self.has_moved = True

        else:
            raise MovementException("Invalid move")
    
    def captured(self):
        return super().captured()
