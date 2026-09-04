#include "logger.h"

#include <chrono>
#include <ctime>
#include <iomanip>
#include <iostream>

char const* to_string(LogLevel log_level)
{
    switch (log_level)
    {
        case LogLevel::Debug:    return "DEBUG";
        case LogLevel::Info:     return "INFO";
        case LogLevel::Warning:  return "WARNING";
        case LogLevel::Error:    return "ERROR";
        case LogLevel::Critical: return "CRITICAL";
    }

    return "UNKNOWN";
}

void log(char const* log_description, LogLevel log_level)
{
    const auto now = std::chrono::system_clock::now();
    const std::time_t current_time =
        std::chrono::system_clock::to_time_t(now);

    const auto milliseconds =
        std::chrono::duration_cast<std::chrono::milliseconds>(
            now.time_since_epoch()
        ) % 1000;

    std::cout << std::put_time(
                     std::localtime(&current_time),
                     "%Y-%m-%d %H:%M:%S"
                 )
              << '.' << std::setfill('0') << std::setw(3)
              << milliseconds.count()
              << " - " << to_string(log_level)
              << " - " << log_description
              << '\n';
}