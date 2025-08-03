# mpy-bundler

Bundle a MicroPython project to a single file.

Ideally, this file can be minified as well.

## Concept


Consider the following MicroPython project:

```
app/
  a/
    a.py
    b.py
  a.py
```

<details>
<summary>The following Python code:</summary>

```python
# app/a/a.py
x = 1
y = 2

def swap():
  global x, y

  tmp = x
  x = y
  y = x

# app/a/b.py
x = 3
y = 4

def swap():
  global x, y

  tmp = x
  x = y
  y = x

# app/a.py
from .app import a
from .app import b

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
```
</details>

<details>
<summary>Can be bundled to:</summary>

```python
# app/a/a.py becomes app_a_a_
app_a_a_x = 1
app_a_a_y = 2

def app_a_a_swap():
  global app_a_a_x, app_a_a_y

  tmp = app_a_a_x
  app_a_a_x = app_a_a_y
  app_a_a_y = app_a_a_x

# app/a/b.py becomes app_a_b_
app_a_b_x = 3
app_a_b_y = 4

def app_a_b_swap():
  global app_a_b_x, app_a_b_y

  tmp = app_a_b_x
  app_a_b_x = app_a_b_y
  app_a_b_y = app_a_b_x

# app/a.py becomes app_a_

def app_a_swap():
  global app_a_a_x, app_a_a_y
  global app_a_b_x, app_a_b_y
  # Cache a values
  tmp_x = app_a_a_x
  tmp_y = app_a_a_y

  # Set a to b
  app_a_a_x = app_a_b_x
  app_a_a_y = app_a_b_y

  # Set b to cached a
  app_a_b_x = tmp_x
  app_a_b_y = tmp_y
```
</details>
