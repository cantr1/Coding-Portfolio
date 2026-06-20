"""
File for the board class - holds position of all pieces
"""
from typing import Optional

from position import Position
from pieces import Piece
from movement_strategy import (
    MovementException,
    PawnMovement,
    RookMovement,
    KnightMovement,
    BishopMovement,
    QueenMovement,
    KingMovement
)

class Square:
    def __init__(self, pos: Position, piece: Piece = None):
        self.pos = pos
        self.piece = piece


class Board:
    def __init__(self):
        self.squares = {
            Position(x, y): Square(Position(x,y))
            for x in range(1, 9)
            for y in range(1, 9)
        }

    def get_piece_at(self, pos: Position) -> Optional[Piece]:
        return self.squares[pos].piece

    def is_empty(self, pos: Position) -> bool:
        return self.get_piece_at(pos) is None

    def place_piece(self, piece: Piece, pos: Position) -> None:
        self.squares[pos].piece = piece

    def move_piece(self, current_pos: Position, target: Position) -> None:
        piece = self.get_piece_at(current_pos)
        if piece is None:
            raise MovementException("No piece at current position")

        piece.validate_move(current_pos, target)
        if not self.is_empty(target):
            raise MovementException("Target position is occupied")
        if not self.is_path_clear(current_pos, target):
            raise MovementException("Path is blocked")

        self.squares[target].piece = piece
        self.squares[current_pos].piece = None
        piece.mark_moved()

    def path_between(self, current_pos: Position, target: Position) -> list[Position]:
        x_diff = target.x_pos - current_pos.x_pos
        y_diff = target.y_pos - current_pos.y_pos

        if abs(x_diff) <= 1 and abs(y_diff) <= 1:
            return []
        if x_diff != 0 and y_diff != 0 and abs(x_diff) != abs(y_diff):
            return []

        x_step = 0 if x_diff == 0 else x_diff // abs(x_diff)
        y_step = 0 if y_diff == 0 else y_diff // abs(y_diff)

        positions = []
        next_pos = Position(current_pos.x_pos + x_step, current_pos.y_pos + y_step)
        while next_pos != target:
            positions.append(next_pos)
            next_pos = Position(next_pos.x_pos + x_step, next_pos.y_pos + y_step)
        return positions

    def is_path_clear(self, current_pos: Position, target: Position) -> bool:
        return all(self.is_empty(pos) for pos in self.path_between(current_pos, target))
    
    def setup_game(self) -> None:
        self.setup_pawns()
        self.setup_rooks()
        self.setup_knights()
        self.setup_bishops()
        self.setup_queens()
        self.setup_kings()

    def setup_pawns(self) -> None:
        # Set white pawns
        for x in range(1, 9):
            self.place_piece(Piece(PawnMovement()), Position(x, 2))
        
        # Set black pawns
        for x in range(1, 9):
            self.place_piece(Piece(PawnMovement()), Position(x, 7))

    def setup_rooks(self) -> None:
        self.place_piece(Piece(RookMovement()), Position(1, 1))
        self.place_piece(Piece(RookMovement()), Position(8, 1))
        self.place_piece(Piece(RookMovement()), Position(1, 8))
        self.place_piece(Piece(RookMovement()), Position(8, 8))
    
    def setup_knights(self) -> None:
        self.place_piece(Piece(KnightMovement()), Position(2, 1))
        self.place_piece(Piece(KnightMovement()), Position(7, 1))
        self.place_piece(Piece(KnightMovement()), Position(2, 8))
        self.place_piece(Piece(KnightMovement()), Position(7, 8))

    def setup_bishops(self) -> None:
        self.place_piece(Piece(BishopMovement()), Position(3, 1))
        self.place_piece(Piece(BishopMovement()), Position(6, 1))
        self.place_piece(Piece(BishopMovement()), Position(3, 8))
        self.place_piece(Piece(BishopMovement()), Position(6, 8))
    
    def setup_queens(self) -> None:
        self.place_piece(Piece(QueenMovement()), Position(4, 1))
        self.place_piece(Piece(QueenMovement()), Position(4, 8))
    
    def setup_kings(self) -> None:
        self.place_piece(Piece(KingMovement()), Position(5, 1))
        self.place_piece(Piece(KingMovement()), Position(5, 8))

    def __str__(self):
        """Prints the board to the terminal"""
        pass
