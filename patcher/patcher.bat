set PYTHONPATH=patcher\services;patcher\libs;patcher\conf;patcher\utils

::call patcher\patcher-env\Scripts\activate.bat

python patcher\main.py generate-crc32 C:/patchTests/fer-0.1.0.war C:/patchTests/fer-0.1.0_new.war C:/patchTests/fer.patch C:/patchTests/fer-crc32.txt