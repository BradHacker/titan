# RITSEC Duckies

This is a research project that started because of a mentorship group with [RITSEC](https://ritsec.club) to learn about using [Rubber Duckies](https://shop.hak5.org/products/usb-rubber-ducky-deluxe). This tool can be used for anything from productivity to attacking a target.

## My project

I'm also very interested in learning more about pentesting, red teaming, and persistence. So, i chose to make my project based around developing a beacon and simple c2 that is deploying in a small amount of time via a Rubber Ducky.

### Starting Point

I started developing this project specifically to attack Windows hosts, developing my beacon in C++ with Visual Studio. It directly interfaced with the Windows APIs to achieve Remote Code Execution (RCE). I wrote a simple Command and Control (C2) Server in Python to act as the remote instructions server.

## The Present

I plan on porting this project to Go and making it work on different platforms. This is the current goal of this project.

## Progress

I'm tracking my progress using [SCRUM](https://www.scrum.org/) set up in GitHub Projects with Automation. You can track the progress of each of my sprints via the [sprints folder](/sprints).

- 02/10-02/17 | [Sprint 1](/sprints/sprint1.md)