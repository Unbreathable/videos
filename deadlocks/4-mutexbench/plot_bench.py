import matplotlib.pyplot as plt
import re
from collections import defaultdict

def parse_bench(filename):
    data = defaultdict(lambda: defaultdict(list))
    # Regex to match BenchmarkHashing[Type]/N=[N]-16 [count] [value] ns/op
    pattern = re.compile(r'BenchmarkHashing(Mutex|Worker)/N=(\d+)-\d+\s+\d+\s+([\d.]+)\s+ns/op')
    
    with open(filename, 'r') as f:
        for line in f:
            match = pattern.search(line)
            if match:
                bench_type, n, val = match.groups()
                data[bench_type][int(n)].append(float(val))
    
    # Average the values
    results = {}
    for bench_type, n_vals in data.items():
        results[bench_type] = {n: sum(vals)/len(vals) for n, vals in n_vals.items()}
    return results

data = parse_bench('bench_results.txt')

# Plot 1: Total time
plt.figure(figsize=(10, 6))
for label, values in data.items():
    sorted_ns = sorted(values.keys())
    sorted_vals = [values[n] for n in sorted_ns]
    plt.plot(sorted_ns, sorted_vals, marker='o', label=label)

plt.xlabel('Number of Hashes (N)')
plt.ylabel('Total ns/op')
plt.title('Benchmark: Mutex vs Channel Worker (Total Time)')
plt.legend()
plt.grid(True)
plt.savefig('hashing_benchmark.png')
print("Graph saved to hashing_benchmark.png")

# Plot 2: Normalized time
plt.figure(figsize=(10, 6))
for label, values in data.items():
    sorted_ns = sorted(values.keys())
    sorted_vals = [values[n] / n for n in sorted_ns]
    plt.plot(sorted_ns, sorted_vals, marker='o', label=label)

plt.xlabel('Number of Hashes (N)')
plt.ylabel('ns/op per Hash')
plt.title('Benchmark: Mutex vs Channel Worker (Normalized per Hash)')
plt.legend()
plt.grid(True)
plt.savefig('hashing_benchmark_normalized.png')
print("Graph saved to hashing_benchmark_normalized.png")
