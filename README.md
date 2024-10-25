## **GitHub Actions Workflow**

This repository uses a **GitHub Actions workflow** to automate testing and reporting. The workflow runs tests both in parallel (using matrix strategy) and separately for fixed input tests. Here’s a breakdown of what the workflow does:

### **Workflow Overview**
1. **Triggers:**
   - **Scheduled Runs**: The workflow runs every Monday at 12:00 PM.
   - **Push Events**: Triggered when code is pushed to the `main` branch.
   - **Pull Requests**: Runs on pull requests targeting the `main` branch.
   - **Manual Trigger**: Can also be triggered manually through GitHub.

2. **Parallel Fuzz Testing**:
   - The workflow uses **matrix strategy** to divide the fuzz tests into two parts:
     - One job tests inputs from range 0–5.
     - Another job tests inputs from range 5–10.
   - Each job uploads its results as artifacts.

3. **Fixed Input Test**:
   - Runs a separate test with a hardcoded input to verify expected output.
   - The results of this test are also uploaded as an artifact.

4. **Report Aggregation**:
   - After all the tests complete, the workflow downloads the individual reports.
   - It **aggregates all reports** into a single file.
   - The aggregated report is uploaded as a final artifact.

### **Key Features**
- **Automated Setup**:
  The workflow **automatically sets up the Go environment**, ensuring consistency and reliability across test runs.
  
- **Flexible Triggers**:  
  The workflow can be triggered through **scheduled events, push or pull requests**, or manually via the GitHub interface.
  
- **Parallel Execution with Go Routines and Channels**:  
  In this project, **Go routines** are used to run tests concurrently, improving efficiency by allowing multiple tests to execute simultaneously. **Channels** facilitate communication between the routines by collecting reports.

- **Artifact Uploads**:  
  All individual test reports are uploaded as artifacts, ensuring that results are available for inspection separately from the aggregated one.

- **Intermediate Reports Access**:  
  Developers can **access intermediate reports** directly from the **GitHub Actions dashboard** to inspect specific ranges or input test results. This is useful for troubleshooting individual tests or checking results early once one parallel job has finished.

- **Aggregated Report**:  
  Once all tests are complete, the individual reports are combined into a **single summary report** for easy review and visibility.

## **How to Extend**
1. **Add New Tests**:  
   You can create new test cases by defining functions using Go's `testing` package. For example, add new tests in the `tests` directory that test different inputs or edge cases.

2. **Modify the Fuzzer**:  
   You can modify the fuzzer's input generation logic to test different scenarios or improve coverage. Edit the `fuzzer.go` file to introduce new types of inputs or test more complex cases.

3. **Update Workflow**:  
   Modify the `.github/workflows/fuzz-testing.yml` file to change test parameters, adjust the test ranges, or customize how tests are triggered. You can also extend the workflow to run tests on multiple Go versions or environments.

## **How to Run**
1. **Install Go**:  
    Ensure that you have Go installed on your machine. If not, you can download and install it from the [Official Go Website](https://go.dev/dl/).

2. **Install TPM Libraries**:  
    This program relies on Google's TPM2 implementation for Go. You need to install the necessary libraries by running the following commands:
    Install the main TPM2 library:
      go get github.com/google/go-tpm/tpm2

    Install the Windows transport package (for accessing TPM on Windows):
      go get github.com/google/go-tpm/tpm2/transport

3. **Run the Program**:  
    Navigate to the directory that contains your main.go file in the terminal and run the Go program with the following command:
      go run main.go

## **Contributing**
Contributions are welcome! If you find a bug or have an idea for an improvement, feel free to open an issue or submit a pull request.

