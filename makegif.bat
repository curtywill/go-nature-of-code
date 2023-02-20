@ECHO OFF
REM This script uses ffmpeg to create a gif based on the images in frames/, then deletes the aformentioned frames.
IF [%1] == [] (
  ECHO expected gif name
  GOTO :EOF
)

SET filename=%1
SET gifcommand=ffmpeg -f image2 -framerate 30 -i frames/frame%%3d.png out/%filename%.gif
%gifcommand%
DEL %cd%\frames\*.png