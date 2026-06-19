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

def on_the_board(p: Position) -> bool:
    if p.x_pos > 8 or p.y_pos > 8 or p.x_pos < 1 or p.y_pos < 1:
        return False
    return True

class PawnMovement(MovementBehavior):
    def check_valid_move(self, current_pos: Position, target: Position, has_moved: bool = False) -> bool:
        # test move is on board (8x8)
        if not on_the_board(target):
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
        if not on_the_board(target):
            return False
        # test move is not on both x and y axis
        if (current_pos.x_pos != target.x_pos) and (current_pos.y_pos != target.y_pos):
            return False
        
        return True
    
class KnightMovement(MovementBehavior):
    def check_valid_move(self, current_pos: Position, target: Position, has_moved: bool = False) -> bool:
        # test move is on board (8x8)
        if not on_the_board(target):
            return False
        # Knights are tricky, can move  y + 1 then x +- 2 or y + 2 then x +- 1
        x_diff: int = abs(target.x_pos - current_pos.x_pos)
        y_diff: int = abs(target.y_pos - current_pos.y_pos)

        if y_diff > 2 or x_diff > 2:
            return False
        
        if y_diff == 2:
            return x_diff == 1
        elif y_diff == 1:
            return x_diff == 2
