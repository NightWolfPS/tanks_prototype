components {
  id: "effect"
  component: "/main/scripts/impact.script"
}
components {
  id: "impact"
  component: "/main/sounds/impact.sound"
}
components {
  id: "impact_proj"
  component: "/main/sounds/impact_proj.sound"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"impact\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "playback_rate: 10.0\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/sprites1/impact.tilesource\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
}
