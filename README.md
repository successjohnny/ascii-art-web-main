# ASCII Art Web

An implementation of a web graphical interface for the ASCII Art algorithm written in standard Go.

## Description
This project scales the functional core of an algorithmic offline terminal layout generator up onto an interactive web environment interface. Users feed string configurations into form entry modules, select standard file configuration profiles (`standard`, `shadow`, or `thinkertoy`), and watch an automated rendering loop yield raw template outputs mapped right onto their UI view layout.

## Authors
- John Ikwuobe <successjohnny@gmail.com>
- Yima Dooga
- Friday igoche

## Usage: How to run
1. Ensure you place the requested banner source file sets (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) straight into the local `banners/` path location directory.
2. Initialize and kick off the server compiler layout execution workflow from your core workflow terminal line prompt environment:
   ```bash
   go run .
   ```
3. Open a browser window grid display canvas frame pointing towards your running address line context target:
   ```text
   http://localhost:8080
   ```

## Implementation Details: Algorithm
- **Web Layer**: Built utilizing vanilla HTTP multiplex handler routines matching incoming queries against strict explicit endpoints (`/` and `/ascii-art`). Forms pass parameter key pairs through body arrays via safe execution sequences.
- **Data Rendering Routing**: Uses standard engine template parsing rules to reflect states back securely using localized contextual escaping structures.
- **Core Processing Engine**: Evaluates character sequences against index coordinates offset by baseline metrics ($Index = Character - 32$). The application opens local flat character sets, reads block lines line by line, and groups them dynamically into consistent vertical text boxes.
