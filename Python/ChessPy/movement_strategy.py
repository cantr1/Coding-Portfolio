"""
Movement Behavior, follwing strategy design pattern
"""
from abc import ABC, abstractmethod
from position import Position

class MovementException(Exception):
    pass


class MovementBehavior(ABC):
    @abstractmethod
    def check_valid_move(self) -> bool:
        pass


class PawnMovement(MovementBehavior):
    def check_valid_move(self, current_pos: Position, target: Position, has_moved: bool = False) -> bool:
        # test move is on board (8x8)
        if target.x_pos > 8 or target.y_pos > 8 or target.x_pos < 1 or target.y_pos < 1:
            return False
        if target.x_pos != current_pos.x_pos:
            return False  # TODO: implement a board class that can return if a piece is at the position for captures
        
        y_pos_change: int = target.y_pos - current_pos.y_pos
        
        if y_pos_change > 2:
            return False
        if y_pos_change <= 0:
            return False
        else:
            if (y_pos_change == 2 and not has_moved) or y_pos_change == 1:
                return True
            else:
                return False

class RookMovement(MovementBehavior):
    def check_valid_move(self, current_pos: Position, target: Position, has_moved: bool = False) -> bool:
        # test move is on board (8x8)
        if target.x_pos > 8 or target.y_pos > 8 or target.x_pos < 1 or target.y_pos < 1:
            return False
        # test move is not on both x and y axis
        if (current_pos.x_pos != target.x_pos) and (current_pos.y_pos != target.y_pos):
            return False