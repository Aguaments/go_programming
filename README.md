# 《Go程序设计语言》课后习题答案
=================================================
timer.h
------------------------------------------------------------------------------------
#ifndef TIMER_H
#define TIMER_H
#include <string>

namespace ducktimer {
    class timer {
    public:
        explicit timer(int dur):duration(dur), filepath("./bgm.mp3"){};
        void run() const;
    private:
        void panel() const;
        void setDuration() const;
        void startDuration() const;
        void startTimer() const;
        int mutable duration;
        std::string filepath;
    };
}
=================================================

=================================================
timer.cpp
------------------------------------------------------------------------------------
//
// Created by g30067333 on 2025/3/6.
//

#include "timer.h"
#include <iostream>
#include <thread>
#include "tools.h"
#include "music.h"

namespace ducktimer {
    void timer::run() const {
        panel();
        startTimer();
    }


    void timer::panel() const{
        std::cout << "==========================================================================================" << std::endl;
        std::cout << "      Welcome to use the timer! Let us to start a regular life. (Default 30 minutes)      " << std::endl;
        std::cout << "==========================================================================================" << std::endl;
        std::cout << "                                                                                          " << std::endl;
        std::cout << "                                     \\(^ ~ ^)/                                           " << std::endl;
        std::cout << "                                                                                          " << std::endl;
        std::cout << "==========================================================================================" << std::endl;
    }

    void timer::setDuration() const {
        do {
            std::cout << "Please tell me the time duration you choose(Don't set the time duration less than 1): ";
            std::cin >> duration;
            if(duration < 1) {
                std::string msg = "OH DAMN! You guys plan to break the system? Set a correct time duration!!!";
                std::cout << msg;
                std::this_thread::sleep_for(std::chrono::seconds(2));
                tools::clearCurrentLine(static_cast<int>(msg.length()));
            }else if(duration == 30) {
                char ans = 'Y';
                std::string msg = "(> _ <)^ ! You don't modify the time duration...(Default time duration 30 minutes)";
                std::cout << msg << std::endl;
                std::cout << "To reset? (Y/N, Default Y) " ;
                std::cin >> ans;
                if(ans == 'N') {
                    std::cout << "Set success." << std::endl;
                    break;
                }
                if(ans == 'Y') continue;
                {
                    std::cout << "/(= _ =)\\ DAMN! You are a bad guy." << std::endl;
                    exit(1);
                }
            }else if(duration >= 1){
                std::cout << "Set success." << std::endl;
                break;
            }else {
                std::cout << "/(= _ =)\\ DAMN! You are a bad guy." << std::endl;
                exit(1);
            }
        }while(true);
    }

    void timer::startDuration() const{
        int t = 0;
        std::cout << "Timer started ... (Target: " << duration << " minutes)" << std::endl;
        do {
            std::this_thread::sleep_for(std::chrono::minutes(1));
            std::cout << "[Min: " << t + 1 << "] "<< t + 1 << " minute elapsed ... ";
            std::cout << "(Set time duration: " << duration ;
            std::cout << " || Total elapsed time : " << t + 1 << ")" << std::endl;
            t ++;
        }while(t < duration);
    }

    void timer::startTimer() const {
        do {
            char ans = 'Y';
            setDuration();
            startDuration();
            std::thread t1(play, filepath);
            t1.detach();
            std::cout << "Continue? (Y/N, Default Y) " ;
            std::cin >> ans;
            if(ans == 'N') {
                std::cout << "OK, Bye. \\(^ - ^)/" << std::endl;
                break;
            }
            if(ans == 'Y') continue;
            {
                std::cout << "/(= _ =)\\ DAMN! You are a bad guy." << std::endl;
                exit(1);
            }
        }while(true);
    }


}
=================================================

=================================================
utils.h
------------------------------------------------------------------------------------
//
// Created by g30067333 on 2025/3/6.
//

#ifndef TOOLS_H
#define TOOLS_H
#define BUFFER_SIZE 4096

namespace ducktimer {
    class tools {
    public:
        static void clearCurrentLine(int len);
        static void setIONewBuffer();
    private:
        static char ibuffer[BUFFER_SIZE];
        static char obuffer[BUFFER_SIZE];
    };
}

#endif //TOOLS_H
=================================================

=================================================
utils.cpp
------------------------------------------------------------------------------------
//
// Created by g30067333 on 2025/3/6.
//

#include "tools.h"
#include <iostream>
#include <thread>
#include <cmath>

namespace ducktimer {

    char tools::ibuffer[BUFFER_SIZE] = {0};
    char tools::obuffer[BUFFER_SIZE] = {0};

    void tools::clearCurrentLine(int len) {
        int t = 1000;
        while(len > 0) {
            std::cout << "\b \b";
            std::this_thread::sleep_for(std::chrono::milliseconds(t));
            t = static_cast<int>(pow(t, 0.5));
            len --;
        }
    }

    void tools::setIONewBuffer(){
        setvbuf(stdout, obuffer, _IONBF, sizeof(obuffer));
        setvbuf(stdin, ibuffer, _IONBF, sizeof(ibuffer));
    }

}
=================================================

=================================================
main.cpp
------------------------------------------------------------------------------------
//
// Created by g30067333 on 2025/3/5.
//
#include "timer.h"


void alertAfterDuration(const int& duration) {

}

int main() {
    ducktimer::timer t(30); // Default setting: 30 minutes
    t.run();
    return 0;
}
=================================================
