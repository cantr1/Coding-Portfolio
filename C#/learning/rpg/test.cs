using System;

namespace Learning.Rpg;

class Test
{
    static void Main()
    {
        Alchemist myAlchemist = new Alchemist("Kelz", "Human");
        Console.WriteLine($"Here is {myAlchemist.Name} current health = {myAlchemist.Health}");
    }
}

