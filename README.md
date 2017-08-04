# dumpy
Run through an array of Linux commands and generate a simple report continaing the results

## Usage
Simply run the binary on the system to gather information from, the results are printed to stdout so you can pipe easily pipe it

```bash
./dumpy > /some/location/report.txt
```

## Customization
Editing the commands currently required recompiling the project after editing `commands.go`, I plan to allow some configuration via either a file or stdin in the future (and PRs or requests for more defaults are welcome)