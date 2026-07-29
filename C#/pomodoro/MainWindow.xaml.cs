using System.Windows;
using System.Windows.Media;

namespace Learning.Pomodoro;

public partial class MainWindow : Window
{
    private PomodoroTimer? _currentTimer;

    public MainWindow()
    {
        InitializeComponent();
    }

    private async void StartButton_Click(object sender, RoutedEventArgs e)
    {
        StartButton.IsEnabled = false;
        StopButton.IsEnabled = true;

        await RunPomodoroCycle();

        StartButton.IsEnabled = true;
        StopButton.IsEnabled = false;
    }

    private void StopButton_Click(object sender, RoutedEventArgs e)
    {
        _currentTimer?.Stop();
        StatusLabel.Text = "Stopped";
        StartButton.IsEnabled = true;
        StopButton.IsEnabled = false;
    }

    private async Task RunPomodoroCycle()
    {
        while (true)
        {
            // Work cycle (25 minutes)
            StatusLabel.Text = "Work Cycle";
            StatusLabel.Foreground = new SolidColorBrush(Colors.Green);

            _currentTimer = new PomodoroTimer(25 * 60);
            _currentTimer.TimeChanged += OnTimeChanged;

            bool completed = await _currentTimer.RunTimerAsync();

            if (!completed) // User stopped the timer
            {
                return;
            }

            // Rest cycle (5 minutes)
            StatusLabel.Text = "Break Cycle";
            StatusLabel.Foreground = new SolidColorBrush(Colors.Blue);

            _currentTimer = new PomodoroTimer(5 * 60);
            _currentTimer.TimeChanged += OnTimeChanged;

            completed = await _currentTimer.RunTimerAsync();

            if (!completed) // User stopped the timer
            {
                return;
            }

            // Ask if user wants to continue
            var result = MessageBox.Show(
                "Cycle complete! Start another cycle?",
                "Pomodoro Timer",
                MessageBoxButton.YesNo,
                MessageBoxImage.Question);

            if (result == MessageBoxResult.No)
            {
                StatusLabel.Text = "Ready to start";
                StatusLabel.Foreground = new SolidColorBrush(Colors.Black);
                TimerDisplay.Text = "25:00";
                TimerProgress.Value = 0;
                return;
            }
        }
    }

    private void OnTimeChanged(object? sender, TimeChangedEventArgs e)
    {
        Dispatcher.Invoke(() =>
        {
            int minutes = e.SecondsRemaining / 60;
            int seconds = e.SecondsRemaining % 60;
            TimerDisplay.Text = $"{minutes}:{(seconds > 9 ? seconds : "0" + seconds)}";
            TimerProgress.Value = e.ProgressPercentage;
        });
    }
}
