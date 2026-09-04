#pragma once

enum class LogLevel
{
    Debug,
    Info,
    Warning,
    Error,
    Critical
};

char const* to_string(LogLevel log_level);

void log(
    char const* log_description,
    LogLevel log_level = LogLevel::Info
);