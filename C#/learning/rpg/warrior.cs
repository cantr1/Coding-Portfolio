using System;

namespace Learning.Rpg;

public class Warrior : Character
{
    public Warrior(string name, string race) : base(name, race)
    {
        Health = 100;
        Armor = "Steel Plate";
        Intelligence = 4;
        Strength = 10;
        Luck = 5;
        Charisma = 3;
        Weapon = "Sword";
    }

    public override (int damage, bool natTwenty) GetBasicStrengthRoll() => RollStat(Strength, 3);

    public override (int damage, bool natTwenty) GetBasicIntelligenceRoll() => RollStat(Intelligence, -4);

    public override (int damage, bool natTwenty) GetBasicLuckRoll() => RollStat(Luck);

    public override (int damage, bool natTwenty) GetBasicCharismaRoll() => RollStat(Charisma);
}