# Hades Build Start Time Reporter

A lightweight Go program that reports the start time of a build job to a specified HTTP endpoint.

## Overview

This tool reads job metadata from environment variables and captures the current timestamp at runtime. It then sends a JSON payload via HTTP `POST` to an external API.

Use cases include:

- Performance benchmarking systems (e.g., Hades Benchmarking)

---

### Required Configuration

| Variable     | Required | Description                  |
|--------------|----------|------------------------------|
| `ENDPOINT`   | ✅       | The target API endpoint URL  |

### 📝 Job Metadata for Artemis (`ResultMetadata`)

| Variable                        | Type   | Default | Description                         |
|----------------------------------|--------|---------|-------------------------------------|
| `JOB_NAME`                      | string | —       | The name of the job                |
| `UUID`                          | string | —       | UUID of the job                    |
| `ASSIGNMENT_REPO_BRANCH_NAME`  | string | `main`  | Branch name of the assignment repo |
| `IS_BUILD_SUCCESSFUL`          | bool   | —       | Whether the build was successful   |
| `ASSIGNMENT_REPO_COMMIT_HASH`  | string | —       | Commit hash of the assignment repo |
| `TESTS_REPO_COMMIT_HASH`       | string | —       | Commit hash of the tests repo      |
| `BUILD_COMPLETION_TIME`        | string | —       | Completion time in RFC3339 format  |

---
