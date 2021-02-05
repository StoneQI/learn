/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#pragma once

#include <fmt/compile.h>

#if defined(__GNUC__) && !defined(__clang__) && __GNUC__ <= 8
// Forcefully disable compiled format strings for GCC 8 & below until fmt is
// updated to do this automatically.
#undef FMT_COMPILE
#endif

#ifdef _MSC_VER
// Workaround broken constexpr in MSVC.
#undef FMT_COMPILE
#endif

// Fallback to runtime format string processing for compatibility with fmt 6.x.
#ifndef FMT_COMPILE
#define FMT_COMPILE(format_str) format_str
#endif
