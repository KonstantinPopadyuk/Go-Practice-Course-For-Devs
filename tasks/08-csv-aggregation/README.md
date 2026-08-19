# 08. Streaming CSV Aggregation

**Time:** 20–25 minutes.

Implement `RevenueByProduct(r io.Reader)`. The input CSV has the header `product,quantity,unit_price`; calculate the revenue for each product in cents.

## Contract

- `quantity` is a positive integer.
- `unit_price` is a non-negative decimal amount with exactly 0–2 digits after the decimal point.
- Do not use `float64` for monetary values: `10.10` must be converted exactly to `1010`.
- CSV fields may be quoted and may contain commas—use `encoding/csv`.
- An empty product name, an invalid header, or a malformed row is an error.
- The error must include the record number (the header is record 1).
- Process the input as a stream; there is no need for `io.ReadAll`.

Run: `go test ./tasks/08-csv-aggregation -v`.

## Going Further

Detect `int64` overflow both when multiplying the price by the quantity and when adding revenue totals.
