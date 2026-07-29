namespace Learning.Pomodoro;

public class TimeChangedEventArgs : EventArgs
{
    public int SecondsRemaining { get; set; }
    public double ProgressPercentage { get; set; }
}

public class PomodoroTimer
{
    public PomodoroTimer(int duration)
    {
        Duration = duration;
        _totalDuration = duration;
    }

    private int Duration { get; set; }
    private readonly int _totalDuration;
    private bool _isStopped;

    public event EventHandler<TimeChangedEventArgs>? TimeChanged;

    public void Stop()
    {
        _isStopped = true;
    }

    public async Task<bool> RunTimerAsync()
    {
        _isStopped = false;

        while (Duration > 0 && !_isStopped)
        {
            double progressPercentage = ((double)(_totalDuration - Duration) / _totalDuration) * 100;

            TimeChanged?.Invoke(this, new TimeChangedEventArgs
            {
                SecondsRemaining = Duration,
                ProgressPercentage = progressPercentage
            });

            await Task.Delay(1000);
            Duration--;
        }

        if (!_isStopped && Duration <= 0)
        {
            TimeChanged?.Invoke(this, new TimeChangedEventArgs
            {
                SecondsRemaining = 0,
                ProgressPercentage = 100
            });
            return true; // Completed naturally
        }

        return false; // Stopped by user
    }
}