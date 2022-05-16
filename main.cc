/*
 * main.cc
 *
 *  Created on: 17 ��� 2022 �.
 *      Author: serge
 */

#include <iostream>
#include <SDL2/SDL_main.h>
#include <SDL2/SDL.h>
#include "Window.h"

int main(int, char **)
{
	SDL_Init(SDL_INIT_EVERYTHING);
	Window w {1920, 1080};
	w.main_loop();
	return 0;
}


