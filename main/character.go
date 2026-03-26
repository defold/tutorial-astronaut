components {
  id: "character"
  component: "/main/character.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"idle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/character/character.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
}
