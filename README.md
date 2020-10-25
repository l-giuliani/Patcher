# Patcher  
### What is Patcher  
Patcher is a software tool that generates and applies patch from and to binary files. It is useful in order to update version of the same executable. to ensure the correct result of the patching operation is possible to generate a file that contains the crc32 of the new version of the executable. At the end of the  apply operation crc32 will be checked to ensure that everything is ok.  
It works very well also with Java jar and war files. For example if you need to update remotely a war file that is large(50 MB for instance), with Patcher you will be able to obtain a patch which does not exceed 150 kB, of course it depends on the changes in the new version.
### Project info 
Is possible to use two versions of Patcher.
- Python program that use a shared library written in GO(source in "patcher" folder)
- Single executable all written in GO(source in "go_patcher" folder)
The single executable written in GO has also a graphic interface that helps to generate the patch 
### Dependencies
- [icedream/go-bsdiff](https://github.com/icedream/go-bsdiff)
- [asticode/go-astilectron](https://github.com/asticode/go-astilectron): for the GUI  
in order to compile these dependencies are needed 
### How can I use it ?  
Single executable:  
You can launch the application without any params and the GUI will appear.  
Using the GUI is very straightforward and you can generate a patch and a crc32 file that you will found in the output folder.  
You can also run commands by command line:
```
#generate patch without crc32
patcher generate full_path_to_file/old_version_of_file full_path_to_file/new_version_of_file output_folder/patch_name
#apply patch without crc32
patcher apply full_path_to_file/old_version_of_file output_folder/patch_name output_folder/patched_file
#print crc32 of file, for debug utility
patcher print-crc32 folder/file
#generate patch with crc32 file
patcher generate-crc32 full_path_to_file/old_version_of_file full_path_to_file/new_version_of_file output_folder/patch_name output_folder/crc32_file_name
#apply patch with crc32 file
patcher apply-crc32 full_path_to_file/old_version_of_file output_folder/patch_name output_folder/patched_file outputfolder/crc32_file_name
```
In order to use the Python version commands are the same, you only have to launch the program with command "python main.py" instead of "patcher"

