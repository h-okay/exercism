# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.28

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/hokay/Code/exercism/cpp/difference-of-squares

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/hokay/Code/exercism/cpp/difference-of-squares

# Include any dependencies generated for this target.
include CMakeFiles/difference-of-squares.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include CMakeFiles/difference-of-squares.dir/compiler_depend.make

# Include the progress variables for this target.
include CMakeFiles/difference-of-squares.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/difference-of-squares.dir/flags.make

CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o: CMakeFiles/difference-of-squares.dir/flags.make
CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o: difference_of_squares_test.cpp
CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o: CMakeFiles/difference-of-squares.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/hokay/Code/exercism/cpp/difference-of-squares/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o -MF CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o.d -o CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o -c /home/hokay/Code/exercism/cpp/difference-of-squares/difference_of_squares_test.cpp

CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing CXX source to CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/hokay/Code/exercism/cpp/difference-of-squares/difference_of_squares_test.cpp > CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.i

CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling CXX source to assembly CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/hokay/Code/exercism/cpp/difference-of-squares/difference_of_squares_test.cpp -o CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.s

CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o: CMakeFiles/difference-of-squares.dir/flags.make
CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o: difference_of_squares.cpp
CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o: CMakeFiles/difference-of-squares.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/hokay/Code/exercism/cpp/difference-of-squares/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o -MF CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o.d -o CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o -c /home/hokay/Code/exercism/cpp/difference-of-squares/difference_of_squares.cpp

CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing CXX source to CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/hokay/Code/exercism/cpp/difference-of-squares/difference_of_squares.cpp > CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.i

CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling CXX source to assembly CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/hokay/Code/exercism/cpp/difference-of-squares/difference_of_squares.cpp -o CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.s

CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o: CMakeFiles/difference-of-squares.dir/flags.make
CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o: test/tests-main.cpp
CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o: CMakeFiles/difference-of-squares.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/hokay/Code/exercism/cpp/difference-of-squares/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o -MF CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o.d -o CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o -c /home/hokay/Code/exercism/cpp/difference-of-squares/test/tests-main.cpp

CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing CXX source to CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/hokay/Code/exercism/cpp/difference-of-squares/test/tests-main.cpp > CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.i

CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling CXX source to assembly CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/hokay/Code/exercism/cpp/difference-of-squares/test/tests-main.cpp -o CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.s

# Object files for target difference-of-squares
difference__of__squares_OBJECTS = \
"CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o" \
"CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o" \
"CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o"

# External object files for target difference-of-squares
difference__of__squares_EXTERNAL_OBJECTS =

difference-of-squares: CMakeFiles/difference-of-squares.dir/difference_of_squares_test.cpp.o
difference-of-squares: CMakeFiles/difference-of-squares.dir/difference_of_squares.cpp.o
difference-of-squares: CMakeFiles/difference-of-squares.dir/test/tests-main.cpp.o
difference-of-squares: CMakeFiles/difference-of-squares.dir/build.make
difference-of-squares: CMakeFiles/difference-of-squares.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --bold --progress-dir=/home/hokay/Code/exercism/cpp/difference-of-squares/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Linking CXX executable difference-of-squares"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/difference-of-squares.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/difference-of-squares.dir/build: difference-of-squares
.PHONY : CMakeFiles/difference-of-squares.dir/build

CMakeFiles/difference-of-squares.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/difference-of-squares.dir/cmake_clean.cmake
.PHONY : CMakeFiles/difference-of-squares.dir/clean

CMakeFiles/difference-of-squares.dir/depend:
	cd /home/hokay/Code/exercism/cpp/difference-of-squares && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/hokay/Code/exercism/cpp/difference-of-squares /home/hokay/Code/exercism/cpp/difference-of-squares /home/hokay/Code/exercism/cpp/difference-of-squares /home/hokay/Code/exercism/cpp/difference-of-squares /home/hokay/Code/exercism/cpp/difference-of-squares/CMakeFiles/difference-of-squares.dir/DependInfo.cmake "--color=$(COLOR)"
.PHONY : CMakeFiles/difference-of-squares.dir/depend

