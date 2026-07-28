using System;

namespace Learning.Rpg;

public abstract class Character
{
    private static readonly Random _random = new Random();

    protected Character(string name, string race)
    {
        Name = name;
        Race = race;
    }

    public string Name { get; set; }
    public int Level { get; set; }
    public string Race { get; set; }
    public int Health { get; set; }
    public string Armor { get; set; }
    public int Intelligence { get; set; }
    public int Strength { get; set; }
    public int Luck { get; set; }
    public int Charisma { get; set; }
    public string Weapon { get; set; }

    protected (int damage, bool natTwenty) RollStat(int statValue, int modifier = 0)
    {
        int roll = _random.Next(1, 21);
        int amplifiedRoll = roll + statValue + modifier;

        if (roll == 20)
        {
            return (roll, true);
        }

        if (amplifiedRoll > 20)
        {
            return (20, false);
        }

        return (amplifiedRoll, false);
    }

    public virtual (int damage, bool natTwenty) GetBasicStrengthRoll() => RollStat(Strength);

    public virtual (int damage, bool natTwenty) GetBasicIntelligenceRoll() => RollStat(Intelligence);

    public virtual (int damage, bool natTwenty) GetBasicLuckRoll() => RollStat(Luck);

    public virtual (int damage, bool natTwenty) GetBasicCharismaRoll() => RollStat(Charisma);

    public virtual void Heal()
    {
        int roll = _random.Next(1, 8 + level);
        Strength += roll;
    }
}