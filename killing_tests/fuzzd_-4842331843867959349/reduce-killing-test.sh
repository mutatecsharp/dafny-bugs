#!/bin/bash

TARGET_BACKENDS=("go")
DAFNY_FILE="main.dfy"
MUTATED_FILE_ID="MUTATE_CSHARP_ACTIVATED_MUTANT20"
MUTANT_ID=38
MUTATED_DAFNY="/home/ubuntu/volume/workspace/dafny"
REGULAR_COMPILATION_TIMEOUT=30
REGULAR_EXECUTION_TIMEOUT=30
OVERALL_ERROR_STATUS="MUTANT_RUNTIME_STDOUT_DIFFER"

function conduct_interestingness_candidate_test {
  local check_for="$1"
  local backend="${TARGET_BACKENDS[0]}"

  case "$check_for" in
    # TODO: other conditions.
    "MUTANT_RUNTIME_STDOUT_DIFFER")
      # Set the baseline to compare against.
      local differ_exit_code=1 # set to 0 if interesting

      # non-mutated

      # Check if the stdout file does not exist
      if [[ ! -e "${backend}_stdout.txt" ]]; then
          echo "Standard output file '${backend}_stdout.txt' does not exist."
          exit 1
      fi

      # Check if the stderr file does not exist
      if [[ ! -e "${backend}_stderr.txt" ]]; then
          echo "Standard error file '${backend}_stderr.txt' does not exist."
          exit 1
      fi

      # mutated

      # Check if the stdout file does not exist
      if [[ ! -e "mutated_${backend}_stdout.txt" ]]; then
          echo "Standard output file 'mutated_${backend}_stdout.txt' does not exist."
          exit 1
      fi

      # Check if the stderr file does not exist
      if [[ ! -e "mutated_${backend}_stderr.txt" ]]; then
          echo "Standard error file 'mutated_${backend}_stderr.txt' does not exist."
          exit 1
      fi

      # Compare between non-mutated and mutated backend output.
      local DIFF
      DIFF=$(diff "${backend}_stdout.txt" "mutated_${backend}_stdout.txt")
      if [[ -n "$DIFF" && "$DIFF" == *"3"* && "$DIFF" == *"0"* ]]; then
        differ_exit_code=0
      fi

      rm "mutated_${backend}_stdout.txt"
      rm "${backend}_stdout.txt"
      exit $differ_exit_code
      ;;
    *)
      exit 1 # case not supported.
      ;;
  esac
}

function compile_candidate_program {
  local target="$1"
  compile_command=("dotnet" "$MUTATED_DAFNY/Binaries/Dafny.dll" "build" "--no-verify" "--allow-warnings" \
   "--target:$target" "$DAFNY_FILE")
  command_string=$(printf " %q" "${compile_command[@]}")
  echo "Compiling program against $target with command: " "${command_string}"
  timeout $REGULAR_COMPILATION_TIMEOUT bash -c "${command_string}"
  local exit_code="$?"

  # Check if the compilation exit code is non-zero.
  if [[ $exit_code -eq 124 ]]; then
    echo "Compilation for ${target} timed out."
    exit 124 # exit 0 if we want to find timeout interesting
  elif [[ $exit_code -ne 0 ]]; then
    echo "Compilation for ${target} failed with exit code $exit_code."
    exit $exit_code # exit 0 if we want to find crashes interesting
  else
    echo "Compilation succeeded for ${target}."
  fi
}

# shellcheck disable=SC2317
function compile_mutated_candidate_program {
  local target="$1"
  local mutated_file_id="$2"
  local mutant_id="$3"

  # enable mutant
  compile_command=("$mutated_file_id=$mutant_id" "dotnet" \
  "$MUTATED_DAFNY/Binaries/Dafny.dll" "build" "--no-verify" "--allow-warnings" \
   "--target:$target" "--output" "mutated_${DAFNY_FILE%.*}" "$DAFNY_FILE")
  command_string=$(printf " %q" "${compile_command[@]}")
  echo "Compiling mutated program against $target with command: " "${command_string}"
  timeout $REGULAR_COMPILATION_TIMEOUT bash -c "${command_string}"
  local exit_code="$?"

  # Check if the compilation exit code is non-zero.
  if [[ $exit_code -eq 124 ]]; then
    echo "Compilation for ${target} timed out."
    exit 124 # exit 0 if we want to find timeout interesting
  elif [[ $exit_code -ne 0 ]]; then
    echo "Compilation for ${target} failed with exit code $exit_code."
    exit $exit_code # exit 0 if we want to find crashes interesting
  else
    echo "Compilation succeeded for ${target}."
  fi
}


function run_candidate_program_command {
  local target="$1"
  local filename="$2"
#  local filename="${DAFNY_FILE%.*}"

  case "$target" in
    "py") echo "python3 $filename-py/__${filename}__.py" ;;
    "cs") echo "dotnet $filename.dll" ;;
    "go") echo "./$filename" ;;
    "js") echo "node $filename.js" ;;
    *) echo "target $target not supported! check reduction setup."; exit 1 ;;
  esac
}

# Results are stored in mutated_{target}_stdout.txt / mutated_{target_stderr}.txt.
# For now, we error if exit code is non-zero.
function run_candidate_program {
  local target="$1"
  local is_mutated="$2" # put a non-empty string to enable
  local run_command

  local prefix=""
  # non-empty string
  if [ -n "$is_mutated" ]; then
    prefix="mutated_"
  fi

  run_command=$(run_candidate_program_command "$target" "${prefix}${DAFNY_FILE%.*}")
  echo "Executing program against $target with command: " "${run_command}"

  timeout $REGULAR_EXECUTION_TIMEOUT bash -c "$run_command" > "${prefix}${target}_stdout.txt" 2> "${prefix}${target}_stderr.txt"
  local exit_code="$?"

  # Check if the stdout file does not exist
  if [[ ! -e "${prefix}${target}_stdout.txt" ]]; then
      echo "Standard output file '${prefix}${target}_stdout.txt' does not exist."
      exit 1
  fi

  # Check if the stderr file does not exist
  if [[ ! -e "${prefix}${target}_stderr.txt" ]]; then
      echo "Standard error file '${prefix}${target}_stderr.txt' does not exist."
      exit 1
  fi

  # Check if the exit code is non-zero.
  if [[ $exit_code -eq 124 ]]; then
    echo "Execution for ${target} timed out."
    exit 124
  #   exit_code=0  # Reset to 0 since timeout is interesting in itself.
  elif [[ $exit_code -ne 0 ]]; then
    echo "Execution for ${target} failed with exit code $exit_code."
    exit $exit_code
  else
    echo "Execution succeeded for ${target}."
  fi
}

# Sanity checks. (target backends should have at least 1)
if [[ "${#TARGET_BACKENDS[@]}" -le 0 ]]; then
  echo "0 targets specified as backend."
  exit 1 # Unable to conduct differential testing.
fi

# 0) Install bignumber.js (dependency.)
npm install bignumber.js

# 1) Validate program compiles with the mutated Dafny compiler
for target_backend in "${TARGET_BACKENDS[@]}"; do
 compile_candidate_program "$target_backend"
done

# 2) Run the program compiled by the mutated Dafny compiler
for target_backend in "${TARGET_BACKENDS[@]}"; do
 run_candidate_program "$target_backend" ""
done

# 1) Validate program compiles with the mutated Dafny compiler with enabled mutants
for target_backend in "${TARGET_BACKENDS[@]}"; do
 compile_mutated_candidate_program "$target_backend" "$MUTATED_FILE_ID" "$MUTANT_ID"
done

# 2) Run the program compiled by the mutated Dafny compiler with enabled mutants
for target_backend in "${TARGET_BACKENDS[@]}"; do
 run_candidate_program "$target_backend" "mutate"
done

# 3) Check for interestingness condition. (miscompilation)
conduct_interestingness_candidate_test "$OVERALL_ERROR_STATUS"

# Sanity checks.
echo "Execution slipped through interesting-ness check: this should not happen!"
exit 1