Yalahweey
=========
Yalahweey is a desktop notification tool which display's the current state of Egyption power grid
based on the public api provided by http://power-grid-status.mos3abof.com/
The tool will notify you regarding the state of the power gird once it start, then it will update you every 15 minutes

Yalahweey: is an Egyption word, we say when we get shocked by hearing bad news;
 and as we are heading to dark ages beacuse of our power grid current state I felt it's the best name for the tool

Screenshots:
=============
* Gnome3

![Gnome3](http://i.imgur.com/uzPJDIC.png)

* KDE

![KDE](http://i.imgur.com/spp6N9k.jpg)

* Unity

![Unity](http://i.imgur.com/irL90lC.png)
 
Installation Steps:
============
Get the ICON

     test -d ~/.icons || mkdir ~/.icons && wget --no-check-certificate -O ~/.icons/yalahweey.png https://raw.githubusercontent.com/asaif/yalahweey/master/icon/yalahweey.png
Get the Binary

    wget --no-check-certificate -O ~/yalahweey https://github.com/asaif/yalahweey/releases/download/0.0.2/yalahweey && chmod +x ~/yalahweey

How to use it ?
==============
Just double click yalahweey at your home folder, and it will update you every 15 minute with the grid status.

Supported Systems:
===========
**Linux :**
* KDE
* Gnome
* Unity

**Windows :**

If you have enough time and can help with windows, fork and send me a PR
