#pragma once
#include <string>

bool FilterDirtyWord(const std::string &input, std::string *output);
int AddWord(const std::string &w);
void ClearWordLib();
