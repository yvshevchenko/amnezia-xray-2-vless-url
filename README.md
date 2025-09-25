# amnezia-xray-2-vless-url CLI tool

## Purpose

The tool converts Amnezia X-Ray JSON configuration file (aka XRay original) to vless://...-like proxy connection URL.

## Usage

1. Download appropriate binary from https://github.com/yvshevchenko/amnezia-xray-2-vless-url/releases/latest assets section
   (Only linux and windows binaries available for now)

2. Save your Amnezia config as a text file

3. Run the tool

   On Linux:
   ```
   ./anezia-xray-to-vless-url /path/to/file.json
   ```

   On windows:

   ```
   anezia-xray-to-vless-url.exe \path\to\file.json
   ```

   _Replace `/path/to/file.json` with actual path to your configuration file._

4. Use generated URL on your own!
