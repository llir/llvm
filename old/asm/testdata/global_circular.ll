@x = global i32** @y

@y = global i32* bitcast (i32*** @x to i32*)
