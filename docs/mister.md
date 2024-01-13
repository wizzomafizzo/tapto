# MiSTer Setup

- [Installation](#installation)
- [Configuration File](#configuration-file)
- [Mappings Database](#mappings-database)

## Installation

Download [TapTo](https://github.com/wizzomafizzo/tapto/releases/latest/) and copy it to the `Scripts` folder on your MiSTer's SD card.

Optionally, add the following to the `downloader.ini` file on your MiSTer's SD card, to receive updates with the `update_all` or `downloader` scripts:
```
[mrext/tapto]
db_url = https://github.com/wizzomafizzo/tapto/releases/latest/tapto.json
```

Once installed, run `tapto` from the MiSTer `Scripts` menu, a prompt will offer to enable TapTo as a startup service, then the service will be started in the background.

After the initial setup is complete, a status display will be shown. It's ok to exit this screen, the service will continue to run in the background.

### Hardware Setup

Your reader may work out of the box with no extra configuration. Run `tapto` from the `Scripts` menu, plug it in, and check if it shows as connected in the log view.

If you are using a PN532 NFC module connected with a USB to TTL cable, then the following config may be needed in `nfc.ini` in the `Scripts` folder:

```ini
[tapto]
probe_device=yes
allow_commands=no
```

Create this file if it doesn't exist.

If nfc.sh is unable to auto-detect your device, it may be necessary to manually configure the connection string:

```ini
[tapto]
connection_string="pn532_uart:/dev/ttyUSB0"
allow_commands=no
```

Be aware the ttyUSB0 part may be different if you have other devices connected such as tty2oled. For a list of possible devices try:

`ls /dev/serial/by-id` or `ls /dev | grep ttyUSB`

## Configuration File

TapTo supports a `tapto.ini` file in the `Scripts` folder. This file can be used to configure the TapTo service.

If one doesn't exist, create a new one. This example has all the default values:

```ini
[tapto]
connection_string=""
allow_commands=no
disable_sounds=no
```

All lines except the `[tapto]` header are optional.

### connection_string

See [Hardware configuration](#hardware-configuration) for details. This option is for configuration of [libnfc](https://github.com/nfc-tools/libnfc) and it currently required for the PN532 module.

### allow_commands

Enables the [command](#run-a-systemlinux-command-command) custom command to be triggered from a tag. By default this is disabled and only works from the `nfc.csv` file described below.

### disable_sounds

Disables the success and fail sounds played when a tag is scanned.

### probe_device

Enables auto detection of a serial based reader device

## Mappings Database

The NFC script supports a `nfc.csv` file in the top of the SD card. This file can be used to override the text read
from a tag and map it to a different text value. This is useful for mapping Amiibos which are read-only, testing text
values before actually writing them, and is necessary for using the `command` custom command.

Create a file called `nfc.csv` in the top of the SD card, with this as the header:
```csv
match_uid,match_text,text
```

You'll then need to either power cycle your MiSTer, or restart the NFC service by running `nfc` from the `Scripts`
menu, selecting the `Stop` button, then the `Start` button.

After the file is created, the service will automatically reload it every time it's updated.

Here's an example `nfc.csv` file that maps several Amiibos to different functions:
```csv
match_uid,match_text,text
04e5c7ca024980,,**command:reboot
04078e6a724c80,,_#Favorites/Final Fantasy VII.mgl
041e6d5a983c80,,_#Favorites/Super Metroid.mgl
041ff6ea973c81,,_#Favorites/Legend of Zelda.mgl
```

Only one `match_` column is required for an entry, and the `match_uid` can include colons and uppercase characters.
You can get the UID of a tag by checking the output in the `nfc` Script display or on your phone.