using System;
using System.Threading;

namespace Learning.Pomodoro;

public class Pomodoro
{
    static void Main()
    {
        // Track if user wants to start another cycle
        bool userContinue = true;

        while (userContinue)
        {
            // Start work timer for 25
            Timer workCycle = new Timer(25 * 60);
            workCycle.RunTimer();

            // Start rest timer for 5
            Timer restCycle = new Timer(5 * 60);
            restCycle.RunTimer();

            // End - prompt user if they want to continue
            Console.WriteLine("Would you like to start another cycle? (Y or N)");
            userContinue = ReturnUserDesireContinue();
        }

    }

    static bool ReturnUserDesireContinue()
    {
        bool validInput = false;
        while (!validInput)
        {
            string userInput = Console.ReadLine();
            if (userInput.ToLower() == "y")
            {
                return true;
            }
            else if (userInput.ToLower() == "n")
            {
                return false;
            }
            else
            {
                Console.WriteLine("Unrecognized input... Enter Y or N");
            }
        }
        return true;
    }
}

public class Timer
{
    public Timer(int duration)
    {
        Duration = duration;
    }

    // Duration - measuered in secs
    private int Duration { get; set; }
    private bool Complete { get; set; }

    private void DisplayTimeLeft()
    {
        int minsLeft = Duration / 60;
        int secsLeft = Duration % 60;

        Console.WriteLine($"{minsLeft}:{(secsLeft > 9 ? secsLeft : "0" + secsLeft)}");
    }

    private void DecrementTimer(int interval)
    {
        Duration -= interval;
        if (Duration <= 0)
        {
            Complete = true;
        }
    }

    public void RunTimer()
    {
        while (!Complete)
        {
            DisplayTimeLeft();
            Thread.Sleep(1000);
            DecrementTimer(1);
        }
    }

}