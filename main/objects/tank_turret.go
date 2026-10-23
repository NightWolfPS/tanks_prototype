components {
  id: "turret_controller"
  component: "/main/scripts/turret_controller.script"
  properties {
    id: "fire_cooldown"
    value: "6.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "turret"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tank/weapon.tilesource\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "projectile_factory"
  type: "factory"
  data: "prototype: \"/main/objects/projectile.go\"\n"
  ""
}
embedded_components {
  id: "muzzle_flash_factory"
  type: "factory"
  data: "prototype: \"/main/objects/muzzle_flash.go\"\n"
  ""
}
