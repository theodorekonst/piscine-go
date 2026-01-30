#!/bin/bash

# step 1: put the interview number into a variable
INTERVIEW_FILE=$(grep -rl "key interview" mystery/interviews | head -n1)
INTERVIEW_NUMBER=$(echo "$INTERVIEW_FILE" | grep -oE '[0-9]+')

# step 2: print the number
echo "$INTERVIEW_NUMBER"

# step 3: print what the interview contains
cat "$INTERVIEW_FILE"

# step 4: print MAIN_SUSPECT
echo "$MAIN_SUSPECT"
