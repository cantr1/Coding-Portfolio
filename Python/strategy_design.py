"""
Strategy Design Pattern
"""

from abc import ABC, abstractmethod


class WeaponBehavior(ABC):
    @abstractmethod
    def use_weapon(self):
        pass


class SwordBehavior(WeaponBehavior):
    def use_weapon(self):
        print("swing sword")


class AxeBehavior(WeaponBehavior):
    def use_weapon(self):
        print("chop axe")


class BowBehavior(WeaponBehavior):
    def use_weapon(self):
        print("shoot arrow")


class Character:
    def __init__(self, name: str, weapon: WeaponBehavior):
        self.name: str = name
        self.weapon: WeaponBehavior = weapon

    def fight(self):
        self.weapon.use_weapon()


king = Character("Kelz", SwordBehavior())
king.fight()
 