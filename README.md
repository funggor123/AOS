# plan9_doc
Plan 9
A. Boot Plan9 on Android Devices:
     1. Download and Install Limbo Emulator (2_9_1_arm) 
          Play Store : https://play.google.com/store/apps/details?id=fr.energycube.android.app.com.limbo.emu.main.armv7&hl=en_US
          Source code : https://github.com/limboemu/limbo
     2. Loading Plan 9 Iamge to Limbo Emulator
          a. Boot by the image from 9front
             i. Download 9front image
                http://9front.org/iso/
             ii. Set limbo as following
                 CPU/Board
                  Architecture: x86
                  Machine Type: pc
                  CPU Model: athlon
                  CPU Cores: 1
                  Ram Memory(MB): 64
                 Storage
                  HardDisk A: create a hardDisk for more than or equal to 10GB
                  HardDisk B: create a hardDisk for more than or equal to 5GB
                 Removable Storage
                  CDROM: load the 9front image 
                 Misc
                  VGA Display: std
                  Network Card: pcnet
                  Others: default setting
                 BootSetting
                  Boot from Device: CD ROM
                 User Interface: SDL
              iii. Run Limbo emulator
              *** If you want to use the PC style keyboard, you can install Hacker Keyboard from playstore 
              Hacker Keyboard: https://play.google.com/store/apps/details?id=org.pocketworkstation.pckeyboard&hl=en
              Features: Allow you to use Alt, Ctrl, Fn....
     3. Installing Plan9 to Disk (Ignoring this when you choose method 2)
            a. Press Enter until you see the GUI interface
            b. Type inst/start
            c. Follow http://neonpulse.net/2017/09/30/installing-9front-in-vmware/ or http://fqa.9front.org/fqa4.html
            (Most of the case, you can use the default setting to install)
            d. Long time for copydist, probably more than one hour for android device
            e. If you see the disk cache is full, re-do step 3
     4. Using Plan9 on android devices
            a. Shortcut Key 
                1. Volumn Down = RIGHT MOUSE KEY
                2. Alt + Volumn Down = Middle KEY (Haven't try)
                3. Tap = LEFT MOUSE KEY 
            b. Frequenctly Use softwares
                1. sam (text Editor)
                   i. type sam -d to open sam
                   ii. after openning sam, start to type the content after typing a 
                   iii. when finishing typing the content, type .
                   iv. type w filename to save the file
                2. Mothra (Browser)
                   features: 1. you can download the content of website to as a file to location environment
                             2. test the network connection    
                3. Git  
                   Copy this git shell wrapper script to /rc/bin
                   http://9legacy.org/9legacy/tools/git
                4. Set environment Variable
                   var=xxx
                5  Compile C
                   i. C compilation: 2c testing.c
                   ii: C linking: 2l testing.2
                   iii Running C exe: out.2
B. Running Plan9 on Linux and Windows
