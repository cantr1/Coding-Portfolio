using System;

namespace Learning.Rpg;

public class Alchemist : Character
{
    public Alchemist(string name) : base(name)
    {
        Race = "Human";
        Health = 80;
        Armor = "Robes";
        Intelligence = 10;
        Strength = 4;
        Luck = 7;
        Charisma = 6;
        Weapon = "Staff";
    }

    public override (int damage, bool natTwenty) GetBasicStrengthRoll() => RollStat(Strength, -2);

    public override (int damage, bool natTwenty) GetBasicIntelligenceRoll() => RollStat(Intelligence, 1);

    public override (int damage, bool natTwenty) GetBasicLuckRoll() => RollStat(Luck);

    public override (int damage, bool natTwenty) GetBasicCharismaRoll() => RollStat(Charisma);
}