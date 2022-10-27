# InstallLib
> Why struggle just to find expensive solutions to creating cross-platform installers, when you can just do it yourself?

InstallLib is a simple and free library for creating beautiful (working on), lightweight, cross-platform installers.

### Getting Started
To get started, create a new GoLang program. You will need to include
```go
import (
    "github.com/DaRubyMiner360/InstallLib"
    "github.com/AllenDang/giu"
)
```
\
Then, wherever you want to create an installer, create an `Installer` object like so:
```go
Installer{}
```
\
Set the following variables inside the installer:
- Title (string)
- Pages ([]Page)

\
The following variables are optional:
- TitleBarMode (int)
- OnCancel (func(page Page))
- OnBack (func(before Page, after Page))
- OnNext (func(before Page, after Page))
- OnFinish (func(page Page))
- Width (int)
- Height (int)
- CancelText (string)
- BackText (string)
- NextText (string)
- FinishText (string)
- GlobalWidgets ([]giu.Widget)
- Copyright (string)
- Flags (giu.MasterWindowFlags)

\
There's some utility constants for `TitleBarMode`:
- TitleBarMode_None (No title bar)
- TitleBarMode_Custom (InstallLib's custom title bar)
- TitleBarMode_Native (Native title bar)

\
`Installer` objects also contain the following variables for use in methods:
- CurrentPage (int)
- Window (*giu.MasterWindow)

\
A `Page` object requires the following variables to be set:
- Title (string)
- Widgets ([]giu.Widget)

\
`Page` objects also contain the following variable for use in methods:
- Installer (*Installer)

\
When you are ready to start the installer, call `Run()` on it.
