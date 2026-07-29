// Lesson 1: C# basics
// Run from the C# folder with:
// dotnet run learning/basics.cs

using System;
using System.Threading;

class Basics
{
    static void Main()
    {
        string learnerName = "Kelz";
        int currentLevel = 1;
        double hoursPracticed = 0.5;
        bool isConsistent = true;
        string currentTopic = "Basics";
        int[] myArray = [1, 2, 3, 4, 5];
        char[] vowels = ['a', 'e', 'i', 'o', 'u'];

        Console.WriteLine($"Welcome to C#, {learnerName}!");
        if (learnerName == "Kelz")
        {
            Console.WriteLine("Cool nickname...");
        }

        // Iterate a string
        foreach (char c in learnerName.ToCharArray())
        {
            if (vowels.Contains(c))
            {
                Console.WriteLine($"Hey! I found a vowel! ({c})");
            }
            else if (c == 'K')
            {
                Console.WriteLine("K is the perfect letter...");
            }
            else
            {
                continue;
            }
        }

        Console.WriteLine($"Level: {currentLevel}");
        Console.WriteLine($"Hours practiced: {hoursPracticed}");
        Console.WriteLine($"Current Topic: {currentTopic}");

        Console.WriteLine(GetConsistencyLine(isConsistent));

        int nextLevel = IncreaseLevel(currentLevel);
        Console.WriteLine($"After today's practice, your level is {nextLevel}.");


        // intList - as it sounds
        List<int> intList = new List<int>();

        // Iterate a standard array
        foreach (int num in myArray)
        {
            intList.Add(num * 2);
        }

        IterateAndPrintArray(intList.ToArray());


        // Dictionaries
        Dictionary<string, int> programmingLanguageRating = new Dictionary<string, int>();

        programmingLanguageRating.Add("Go", 9);
        programmingLanguageRating.Add("Python", 7);
        programmingLanguageRating.Add("JS", 4);

        // Iterate Dict
        foreach (var kvp in programmingLanguageRating)
        {
            Console.WriteLine($"{kvp.Key} - Rating: {kvp.Value}");
        }

        // Check for value
        if (programmingLanguageRating.TryGetValue("Haskell", out int rating))
        {
            Console.WriteLine($"Found Haskell in dict - Rating {rating}");
        }
        else
        {
            Console.WriteLine("Haskell not found in dict");
        }

        int upperLimit = 3;
        int currentIter = 0;
        while (currentIter < upperLimit)
        {
            Console.WriteLine($"Current Iteration: {currentIter}");
            Thread.Sleep(1000); //ms
            currentIter += 1;
        }
    }

    // --- Function defines ---
    static void IterateAndPrintArray(int[] numArray)
    {
        foreach (int num in numArray)
        {
            PrintNum(num);
        }
    }

    static void PrintNum(int n)
    {
        Console.Write(n + " - ");
    }

    static int IncreaseLevel(int level)
    {
        return level + 1;
    }

    static string GetConsistencyLine(bool isConsistent)
    {
        if (isConsistent)
        {
            return "Small, steady practice compounds.";
        }
        else
        {
            return "Pick one small exercise and restart there.";
        }
    }

    // Method Overloading
    static int PlusOne(int x)
    {
        return x + 1;
    }

    static double PlusOne(double x)
    {
        return x + 1.0;
    }
}


