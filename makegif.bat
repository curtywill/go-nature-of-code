@ECHO OFF
REM This script uses ffmpeg to create a gif based on the images in out/, then deletes the aformentioned frames.
IF [%1] == [] (
  ECHO expected gif name
  GOTO :EOF
)

SET filename=%1
SET gifcommand=ffmpeg -f image2 -framerate 15 -i out/frame%%3d.png gifs/%filename%.gif
%gifcommand%
DEL %cd%\out\*.png