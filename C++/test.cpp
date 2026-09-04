#include "logger.h"
#include <iostream>
#include <vector>

void display_temperature(float temp)
{
    std::cout << "Temperature: " << temp << "\n";
}

void update_temperature(std::vector<float>& temp_history, float new_reading)
{
    temp_history.push_back(new_reading);
    log("temperature updated", LogLevel::Info);
}

void display_latest_temperature(std::vector<float> const& temp_history)
{
    if (temp_history.empty())
    {
        std::cout << "";
        return;
    }

    std::cout << "Latest Temperature: " << temp_history.back() << '\n';
}

float get_latest_temperature(std::vector<float> const& temp_history)
{
    return temp_history.back();
}

int main()
{
    std::vector<float> temp_history;
    update_temperature(temp_history, 72.2); // Initial reading
    log("Beginning temperature program", LogLevel::Info);
    while (true)
    {
        display_temperature(get_latest_temperature(temp_history));
        int user_choice;

        std::cout << "1: Record temperature\n" << "2. Show latest temperature\n" << "3. Quit\n" << "Choice: ";
        std::cin >> user_choice;

        switch (user_choice)
        {
            case 1:
                float new_reading;
                std::cout << "Enter new temperature: ";
                std::cin >> new_reading;
                update_temperature(temp_history, new_reading);
                break;
            case 2:
                display_latest_temperature(temp_history);
                break;
            case 3:
                exit(0);
            default:
                std::cout << "unrecognized input";
                break;
        }
    }


    log("Program Complete", LogLevel::Info);

    return 0;
}