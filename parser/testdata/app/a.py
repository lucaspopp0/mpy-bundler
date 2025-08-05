from .a import a
from .a import b

def swap():
  # Cache a values
  tmp_x = a.x
  tmp_y = a.y

  # Set a to b
  a.x = b.x
  a.y = b.y

  # Set b to cached a
  b.x = tmp_x
  b.y = tmp_y
