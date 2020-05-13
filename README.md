# Walking astronaut tutorial

In this beginner's tutorial you will learn how to capture player input and make a character move and animate. Along the way you will get an introduction to the core building blocks in Defold: game objects, components and collections.

This tutorial project comes preset with all assets that you need. Start by [running the game](defold://build) (<kbd>Project ▸ Build</kbd>) to get a feel for what's in here.

## The building blocks of a Defold game

Open ["/main/main.collection"](defold://open?path=/main/main.collection) (locate the file in the *Assets* pane to the left and double click it.)

<img src="doc/main_collection.png" srcset="doc/main_collection@2x.png 2x">

What you see in the editor is a *Collection* file. When the Defold engine starts, it loads a *bootstrap collection* as specified in the "game.project" settings file. This particular collection file is this project's bootstrap collection, which is why you see the content of it when you run the game. This collection contains two *Game Objects*. And these two game objects each contain *Components*:

<img src="doc/building_blocks.png" srcset="doc/building_blocks@2x.png 2x">

*COLLECTION*
: Collection files contain game objects and other collections (sub-collections). You organize your game objects by adding them to collection files. You use collections to build small things like a player character or a boss, and you use them to build big things like whole levels.

*GAME OBJECT*
: Game objects hold sprites, sounds, 3D models, tiles or scripts (programmed behavior). A game object has position, rotation and scale. You can write script code that manipulate these properties while the game is running. A typical game object is a bullet, a pickup object or a level loader.

*COMPONENT*
: Components are the things that are drawn on screen, make sounds or make interactions happen. Components do not exist on their own but are placed inside game objects. Some components have properties that can be manipulated in runtime and most components can be turned on and off (enabled and disabled). Many component types source their content from separate resource files, like atlas image files, sound files, animation files etc.

## What's in the main.collection file?

Looking at the "main.collection" file, the editor shows the visual content of the collection in the center editor view. In the right hand *Outline* pane, the collection content is presented as a tree structure. There are two game objects in the collection:

1. The astronaut game object which has *Id* "astronaut". It contains a *Sprite* component and a *Script* component. The game object has been added to the main collection from a blueprint game object file named "/main/astronaut.go".

2. The background level game object which has *Id* "level". It contains a *Tilemap* component. This game object has been embedded directly in the collection file so there is no separate file. The tilemap component, however, sources its tilemap data from a separate file named "/main/level.tilemap".

There is no difference between game objects embedded directly in a collection or ones that are based on blueprint files. If a game object exists in only one instance, there is no real advantage to using a blueprint file, but if you want to create many copies of the same object, blueprint files are extremely convenient since they allow you to change all instances at once. In the *Outline* pane the name of the file an object or component is created from is written in italics next to the object id.

## The astronaut game object

Open ["/main/astronaut.go"](defold://open?path=/main/astronaut.go) to view the blueprint file that is used to create the astronaut instance in "main.collection". As with a collection, the editor shows the game object's content in the center editor view and the *Outline* view to the right shows the structure. This game object file consists of two components:

1. A *Script* component, based on the script file "/main/astronaut.script".
2. A *Sprite* component, which is embedded in place in the game object file.

<img src="doc/astronaut_go.png" srcset="doc/astronaut_go@2x.png 2x">

Click the sprite component to select it. The *Properties* view in the lower right corner now lists all properties that are associated with the sprite component. The sprite has its own *Id*, *Position* and *Rotation* properties. These are all relative to the game object that harbours the component. There are also properties that dictate what image or animation the sprite should display:

*Image*
: This property points to an image resource that is used as a source for the sprite's graphics. Image resources are *Atlas* or *Tilesource* files. *Atlas* files are collections of separate images that have been baked into a larger image for efficiency. Here, the property is set to the file "/main/astronaut.atlas".

*Default Animation*
: This property indicates which particular image or animation in the image resource should be used. Here, the property is set to the "idle" animation.

[Run the game again](defold://build). Notice that the astronaut sprite is looping through an idle animation. Let's now have a look at how that animation is set up.

## Atlas animations

Open the file ["/main/astronaut.atlas"](defold://open?path=/main/astronaut.atlas). The editor view in the center shows each image that has been added to the atlas. The *Outline* view shows all the images and how animations are organized.

<img src="doc/astronaut_atlas.png" srcset="doc/astronaut_atlas@2x.png 2x">

An *Animation Group* is a list of images that are played in a specified order at a specified playback speed. There is currently only one animation group present. It's called "idle" and consists of five separate images. Select the "idle" animation in the outline and choose <kbd>View ▸ Play</kbd> from the menu to preview the animation. You may have to select <kbd>View ▸ Frame Selection</kbd> to zoom the editor camera to cover the whole atlas.

To start adding the walk animations to the astronaut, <kbd>right click</kbd> the root of the atlas outline and select <kbd>Add Animation Group</kbd>.

<img src="doc/add_animation.png" srcset="doc/add_animation@2x.png 2x">

Click the new animation group (named "New Animation") and give it the *Id* "left". Then <kbd>right click</kbd> the animation group and select <kbd>Add Images...</kbd>

<img src="doc/add_images.png" srcset="doc/add_images@2x.png 2x">

Type "left" in the top filter text field to see only images with the name "left" in them. Select all the images that appear (hold <kbd>Shift</kbd> and click) and confirm with <kbd>OK</kbd>.

<img src="doc/select_images.png" srcset="doc/select_images@2x.png 2x">

The new animation has 6 images in it. Select <kbd>View ▸ Play</kbd> from the menu to preview the animation. The speed of the animation is too high so reduce the *Fps* (frames per second) property of the "left" animation from 60 to 15.

Repeat these last steps and add animations for walking "right", "front" and "back" to the atlas in the same way as you added the "left" animation.

## The astronaut script component

Remember that the astronaut game object has a *Script* component based on the file "/main/astronaut.script"? Open ["/main/astronaut.script"](defold://open?path=/main/astronaut.script) to view the Lua script file. As you can see, the script file contains a set of empty functions. These are the *lifetime functions* of the astronaut:

`init(self)`
: This function is called when the component is initialized, before anything appears on the screen. You will use this function to set a few things up.

`final(self)`
: This function is called when the component is being removed: when the game object is deleted or right before the engine shuts down your game.

`update(self, dt)`
: This function is called once each frame. It is useful for doing manipulations and calculations that need to happen it real-time. You will use this function to move the game object based on input.

`on_message(self, message_id, message, sender)`
: This function is called each time a message is sent to the script component. Message passing is a central feature of Defold but we are not doing any in this tutorial.

`on_input(self, action_id, action)`
: This function is called each time an input action is sent to the script component. Input actions are defined in the file ["/input/game.input_binding"](defold://open?path=/input/game.input_binding). This project has bindings already set up for the arrow buttons: "left", "right", "front" and "back" and you will use this function to react to input.

`on_reload(self)`
: This function is called whenever the current script component is *hot-reloaded* into a running game. This is very useful to inspect or manipulate the state of a game object at reload to test things or do debugging.

## Programming the astronaut movement

You are now ready to write a bit of Lua code to play the animations and to make the astronaut game object move. First, change the content of the `init()` function to the following:

```lua
local speed = 150                                             -- [1]

function init(self)
    msg.post(".", "acquire_input_focus")                      -- [2]
    self.dir = vmath.vector3()                                -- [3]
end
```
1. Define a local variable (constant) that holds the movement speed (in pixels/s).
2. Send a built in engine message to the current game object (".") telling it to listen to input.
3. Define a variable that is part of the current script component instance (`self`). The variable will hold the movement direction, expressed as a vector. It is initially zero.

Second, change the content of the `on_input()` function:

```lua
function on_input(self, action_id, action)
    if action_id == hash("front") then                        -- [1]
        self.dir.y = -1
    elseif action_id == hash("back") then
        self.dir.y = 1
    elseif action_id == hash("left") then                     -- [2]
        self.dir.x = -1
    elseif action_id == hash("right") then
        self.dir.x = 1
    end
end
```
1. Actions defined in the input binding file are sent in the `action_id` parameter. If the user presses the "front" or "back" button, set the Y component of the movement direction vector.
2. If the user presses the "left" or "right" button, set the X component of the movement direction vector.

Note that if the player presses "front" and "left" at the same time, two calls will be done to `on_input()` and both the X and Y components of the direction vector will be altered.

Third, change the content of the `update()` function:

```lua
function update(self, dt)
    if vmath.length_sqr(self.dir) > 1 then                   -- [1]
        self.dir = vmath.normalize(self.dir)
    end
    local p = go.get_position()                              -- [2]
    go.set_position(p + self.dir * speed * dt)               -- [3]
    self.dir = vmath.vector3()                               -- [4]
end
```
1. When the `update()` function is called, the engine has already processed all input, meaning that the direction vector is set. In the case of diagonal movement, the length of the movement vector is greater than 1. Normalizing the direction vector makes it length 1 and diagonal movement will have the same speed as horizontal and vertical movement.
2. Get the position of the current game object. The name `go` *does not* refer to the current game object. It is the name of the Lua module that contains all game object functions.
3. Set the position of the current game object to the old position plus the direction vector scaled with the speed constant and `dt`. Multiplying with `dt` makes the movement velocity independent of the update frequency.
4. Reset the direction vector since it is set each frame in `on_input`.

[Run the game again](defold://build) and verify that movement works as expected. The astronaut should move in all 8 directions.

## Adding animations to the movement

The final piece of the puzzle is to change the animation that is played depending on the movement direction. For that, you have to first add a variable that holds the current animation:

```lua
function init(self)
    msg.post(".", "acquire_input_focus")
    self.dir = vmath.vector3()
    self.current_anim = nil                                  -- [1]
end
```
1. Store the currently running animation.

Then you need to add code in `update()` that changes animation based on direction:

```lua
function update(self, dt)
    if vmath.length_sqr(self.dir) > 1 then
        self.dir = vmath.normalize(self.dir)
    end
    local p = go.get_position()
    go.set_position(p + self.dir * speed * dt)

    -- animate the astronaut

    local anim = hash("idle")                                  -- [1]

    if self.dir.x > 0 then                                     -- [2]
        anim = hash("right")
    elseif self.dir.x < 0 then
        anim = hash("left")
    elseif self.dir.y > 0 then
        anim = hash("back")
    elseif self.dir.y < 0 then
        anim = hash("front")
    end

    if anim ~= self.current_anim then                          -- [3]
        msg.post("#sprite", "play_animation", { id = anim })   -- [4]
        self.current_anim = anim                               -- [5]
    end

    -- done animating

    self.dir = vmath.vector3()
end
```
1. Local variable that starts with the default animation id. With no input, this is the animation you get.
2. Test against the movement direction and set the `anim` variable depending on the value of the X and Y component in the direction vector.
3. If `anim` is different than the current animation:
4. then play the new animation,
5. then set the current animation id to the id in `anim`.

The test against the current animation is required, otherwise the code would restart the same animation over and over again each frame. Also note that there are no separate diagonal walk animations but the code uses "left" and "right" for the diagonals.

[Run the game again](defold://build) and verify that the astronaut moves and animates correctly.

**Congratulations! You have now finished this tutorial. We hope that you found it instructive.**

Check out the [documentation pages](https://defold.com/learn) for more examples, tutorials, manuals and API docs.

If you run into trouble, help is available in [our forum](https://forum.defold.com).

Happy Defolding!


## The complete movement script

```lua
local speed = 150

function init(self)
    msg.post(".", "acquire_input_focus")
    self.dir = vmath.vector3()
    self.current_anim = nil
end

function update(self, dt)
    if vmath.length_sqr(self.dir) > 1 then
        self.dir = vmath.normalize(self.dir)
    end
    local p = go.get_position()
    go.set_position(p + self.dir * speed * dt)

    -- animate the astronaut

    local anim = hash("idle")

    if self.dir.x > 0 then
        anim = hash("right")
    elseif self.dir.x < 0 then
        anim = hash("left")
    elseif self.dir.y > 0 then
        anim = hash("back")
    elseif self.dir.y < 0 then
        anim = hash("front")
    end

    if anim ~= self.current_anim then
        msg.post("#sprite", "play_animation", { id = anim })
        self.current_anim = anim
    end

    -- done animating
    
    self.dir = vmath.vector3()
end

function on_input(self, action_id, action)
    if action_id == hash("front") then
        self.dir.y = -1
    elseif action_id == hash("back") then
        self.dir.y = 1
    elseif action_id == hash("left") then
        self.dir.x = -1
    elseif action_id == hash("right") then
        self.dir.x = 1
    end
end
```

---
