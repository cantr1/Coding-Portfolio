"""
Main file for gameplay loop
"""
from board import Board
from position import Position
from enum import Enum
import subprocess


class AlphaPosition(Enum):
    A = 1
    B = 2
    C = 3
    D = 4
    E = 5
    F = 6
    G = 7
    H = 8

def clear_screen() -> None:
    try:
        subprocess.run(["clear"], check=True)
    except subprocess.CalledProcessError:
        print("Error clearing the screen")

def main() -> None:
    board = Board()
    board.setup_game()
    print("Welcome to Chess\n")
    
    game_on = True
    while game_on:
        print(board)
        user_input: str = input("Choose your move in the following format: B1 C3\n").upper()
        starting_user_pos: str = user_input.split(" ")[0]
        starting_x_pos: int = AlphaPosition[starting_user_pos[0]].value
        starting_y_pos: int = int(starting_user_pos[1])
        starting_pos: Position = Position(starting_x_pos, starting_y_pos)

        target_user_pos: str = user_input.split(" ")[1]
        target_x_pos: int = AlphaPosition[target_user_pos[0]].value
        target_y_pos: int = int(target_user_pos[1])
        target_pos: Position = Position(target_x_pos, target_y_pos)

        #DEBUG
        print(f"Attempting to move {starting_pos} -> {target_pos}")

        board.move_piece(starting_pos, target_pos)

        clear_screen()
        
        


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nGoodbye!")