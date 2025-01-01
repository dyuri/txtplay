/*
TODO
- import bubbletea
- import cobra
- create a new bubbletea program
  - create a new model
    - model: width, height, content (string), folder (string), current_frame (int), running (bool)
  - create a new init function
  - create a new update function
    - handle key events (q - quit, p - pause)
	- handle tick events (every 20ms)
	  - if running is true, increment current_frame
	  - if current_frame is greater than the number of frames, set current_frame to 0
	  - read the content of the `current_frame`-th file from the folder, and set it to the model as `content`
  - create a new view function that displays the `content`
- create a new cobra command
  - create a new root command
  - create a play command
    - with a `folder` argument
	- it should set the folder in the model and start the bubbletea program
- main function
  - run the cobra command
*/
