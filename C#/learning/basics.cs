// Lesson 1: C# basics
// Run from the C# folder with:
// dotnet run learning/basics.cs

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

// Iterate a standard array
foreach (int num in myArray)
{
    PrintNum(num);
    if (num == 5)
    {
        Console.WriteLine();
    }
}

static void PrintNum(int n)
{
    Console.Write(n);
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
