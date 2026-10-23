components {
  id: "muzzle_flash"
  component: "/main/scripts/muzzle_flash.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/sprites1/muzzleflash.tilesource\"\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
  rotation {
    z: 1.0
    w: 6.123234E-17
  }
}
