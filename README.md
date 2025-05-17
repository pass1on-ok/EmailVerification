# Email Verification Tool

## Installation

To get started with the Email Verification Tool, you need to have Go installed on your system. Follow these steps to install and set up the tool:

1. **Clone the Repository**:

    ```bash
    git clone https://github.com/pass1on-ok/EmailVerification.git
    ```

2. **Navigate to the Project Directory**:

    ```bash
    cd EmailVerification
    ```

3. **Install Dependencies**:

    ```bash
    go mod tidy
    ```

4. **Build the Tool**:

    ```bash
    go build -o EmailVerification
    ```

## Usage

You can use the Email Verification Tool directly from the command line. Here’s how to use it:

1. **Run the Tool**:

    ```bash
    ./ev --email=<email-address>
    ```

2. **Example**:

    ```bash
    ./ev --email=test@example.com
    ```

   The tool will output whether the email address is likely valid or not based on the checks performed.



