using System;

namespace Learning.Rpg;

public class Rogue : Character
{
    public Alchemist(string name, string race) : base(name, race)
    {
        Health = 85;
        Armor = "Leather";
        Intelligence = 7;
        Strength = 6;
        Luck = 9;
        Charisma = 8;
        Weapon = "Dagger";
    }

    public override (int damage, bool natTwenty) GetBasicStrengthRoll() => RollStat(Strength, 1);

    public override (int damage, bool natTwenty) GetBasicIntelligenceRoll() => RollStat(Intelligence);

    public override (int damage, bool natTwenty) GetBasicLuckRoll() => RollStat(Luck + 3);

    public override (int damage, bool natTwenty) GetBasicCharismaRoll() => RollStat(Charisma + 1);
}