# Pietrzak VDF Simulator
This Golang project implements the Pietrzak Verifiable Delay Function([Simple Verifiable Delay Functions - Krzysztof Pietrzak](https://eprint.iacr.org/2018/627.pdf)),
utilizing cryptographic proofs to ensure data integrity and security for applications requiring verifiable sequential work with outputs that are easy to verify but hard to compute. 
We are currently researching the integration of the delta method to accelerate the proof process, and plan to include it in future updates. 
Additionally, we employ the commit-reveal-recover scheme for secure(https://github.com/tokamak-network/Commit-Reveal-Recover-RNG) and reliable random number generation for blockchain applications. Further enhancements are planned to expand the repository with additional features.


## Table of Contents
- [Dependencies](#dependencies)
- [Installation](#installation)
- [Running the Simulator](#running-the-simulator)
- [Performance Goals and Experimental Results](#performance-goals-and-experimental-results)
- [Contributing](#contributing)
- [License](#license)

## Dependencies
This project is implemented in Go and depends on several external libraries.
To handle dependencies, this project uses Go modules. Here is the `go.mod` file configuration required:

Ensure that you have Go 1.22 or later installed before proceeding.
Ensure that your go.mod file is up-to-date with these entries to correctly handle all dependencies. To sync all modules, run:
```
go mod tidy
```

## Installation
Clone the repository and install the necessary dependencies:
```
git clone https://github.com/tokamak-network/Pietrzak-VDF-Prover.git
go get -u github.com/tokamak-network/Pietrzak-VDF-Prover
```

## Running the Simulator

The provided Go code includes a `main` function and helper functions to set up the environment, evaluate the VDF, generate proofs, and verify them. Here’s how each part works:

### Setup Function
Initializes the cryptographic settings.
- **Parameters**:
    - `bits`: Represents the bit length of the prime number `N`. Typically set to 2048 for strong security.
- **Outputs**:
    - `N`: A large prime number.
    - `x`: A random starting value smaller than `N`.
- **How it works**: Generates a prime number `N` and a starting value `x`, which is essential for the VDF calculation.

### Evaluate Function
Executes the VDF to compute a new value `y` based on `x` over a time parameter `T`.
- **Parameters**:
    - `N`, `x`: Outputs from the setup function.
    - `T`: The time parameter defining the difficulty and duration of the VDF operation. Higher values make the function harder and more time-consuming to compute.
- **Outputs**:
    - `proofList`: A list of claims that summarize the evaluation.
    - `claim`: A single claim containing all values `N`, `x`, `y`, `T`, and `v` (partial evaluation).
- **How it works**: Computes `y` as `x` raised to the power of `2^T` mod `N`. It also calculates `v`, a checkpoint value at `T/2`, to assist in the verification.

### Generate Proof Function
Generates a proof to verify the correctness of the computed values.
- **Parameters**:
    - `claim`: The outputs from the evaluate function.
- **Outputs**:
    - `proofList`: An enhanced list of claims with additional proofs.
- **How it works**: Refines the initial claim by generating additional proofs to support the verification of the VDF computation.

### Verify Function
Verifies the proof to ensure the integrity and correctness of the VDF computation.
- **Parameters**:
    - `proofList`: The list of claims including proofs from the generate proof function.
- **Outputs**:
    - Returns a boolean indicating the success or failure of the verification.
- **How it works**: Uses cryptographic techniques to check the validity of the proofs against the original values.

### Example Usage

To run the simulator, follow these steps:
1. Ensure that you have Go installed and that all dependencies are set up as per the Dependencies section.
2. Navigate to the directory containing the code.
3. Run the code using:
   ```bash
   cd vdf-prover/cmd
   go run main.go
   
## Performance Goals and Experimental Results
Our goal is to implement a Verifiable Delay Function (VDF) that achieves similar performance to the Pietrzak VDF C++ implementation described in the paper, "Implementation Study of Two Verifiable Delay Functions". The paper is available at [this link](https://drops.dagstuhl.de/storage/01oasics/oasics-vol082-tokenomics2020/OASIcs.Tokenomics.2020.9/OASIcs.Tokenomics.2020.9.pdf).

### Experimental Setup Comparison
Our experiments were conducted using an AMD Ryzen 5600X processor, whereas the experimental results reported in the tokenomics paper were obtained using an Intel Core i7 8th Generation.
We anticipate that the performance will differ by about 20-25% due to variations in hardware specifications used in different setups. Based on this, we have derived expected results that reflect these differences.

![](/data/Experimental-setup.png)

| Specification | Cores | Threads | CPU-Z Benchmark | Clock Speed |
| --- | --- | --- | --- | --- |
| AMD Ryzen 5600X | 6 | 12 | 643 | 4.00 - 4.50 GHz |
| Intel Core i7-8700K | 6 | 12 | 500 | 3.70 GHz |

### References
- [Tokenomics 2nd Conference - Implementation Study of Two Verifiable Delay Functions](https://www.youtube.com/watch?v=uC3j0pCEP7o)
- [CPU-Z Benchmark for AMD Ryzen 5 5600X (1T) - CPU-Z VALIDATOR](https://valid.x86.fr/bench/rsf5p1/1)
- [CPU-Z Benchmark for Intel Core i7-8700 (8T) - CPU-Z VALIDATOR](https://valid.x86.fr/bench/d9s89x/8)

### Experimental Results
The experimental results were obtained after conducting five separate trials. After each trial, the results were compiled and analyzed to derive the final outcomes.
We plan to increase the number of samples and conduct the tests again to verify the robustness of our findings and further refine our VDF implementation.Experimental Setup Comparision

![](/data/Evaluation-Result.png)
![](/data/Proof-Result.png)
![](/data/Evaluation-ExResult.png)
![](/data/Proof-ExResult.png)